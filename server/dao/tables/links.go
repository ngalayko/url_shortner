// Code generated by gen_schema_tables.go DO NOT EDIT.

package tables

import (
	"bytes"
	"fmt"

	"go.uber.org/zap"

	"github.com/ngalayko/url_shortner/server/dao"
	"github.com/ngalayko/url_shortner/server/schema"
)

// GetLinkById returns Link from db or cache
func (t *Tables) GetLinkById(id uint64) (*schema.Link, error) {
	return t.GetLinkByFields(dao.NewParam(1).Add("id", id))
}

// GetLinkByFields returns Links from db or cache
func (t *Tables) GetLinkByFields(field dao.Param) (*schema.Link, error) {
	fields := dao.NewParams(1).Append(field)

	ll, err := t.SelectLinksByFields(fields)
	if err != nil {
		return nil, err
	}

	return ll[0], nil
}

// SelectLinksByFields select many links by fields
func (t *Tables) SelectLinksByFields(fields dao.Params) ([]*schema.Link, error) {

	if fields.Len() == 0 {
		return nil, nil
	}

	result := make([]*schema.Link, 0, fields.Len())
	missedFields := dao.NewParams(fields.Len())
	for _, f := range fields {

		if value, ok := t.cache.Load(t.linksCacheKey(f["id"])); ok {
			result = append(result, value.(*schema.Link))
			continue
		}

		missedFields = append(missedFields, f)
	}

	if missedFields.Len() == 0 {
		return result, nil
	}

	b := bytes.Buffer{}
	b.WriteString("SELECT * " +
		"FROM links " +
		"WHERE ")

	i := 1
	values := make([]interface{}, 0, missedFields.Len())
	for fi, missedF := range missedFields {

		if fi > 0 {
			b.WriteString(" OR ")
		}

		b.WriteRune('(')
		j := 0
		for key, value := range missedF {
			values = append(values, value)

			if j > 0 {
				b.WriteString(" AND ")
			}

			b.WriteString(fmt.Sprintf("%s = $%d", key, i))

			i++
			j++
		}
		b.WriteRune(')')
	}

	ll := make([]*schema.Link, 0, missedFields.Len())
	if err := t.db.Select(&ll, b.String(), values...); err != nil {
		return nil, err
	}

	for _, l := range ll {
		t.cache.Store(t.linksCacheKey(l.ID), l)
		result = append(result, l)
	}

	return result, nil
}

// InsertLink inserts Link in db and cache
func (t *Tables) InsertLink(l *schema.Link) error {
	return t.db.Mutate(func(tx *dao.Tx) error {

		insertSQL := "INSERT INTO links " +
			"(user_id, url, short_url, views_limit, views, expired_at, created_at, deleted_at) " +
			"VALUES " +
			"($1, $2, $3, $4, $5, $6, $7, $8) " +
			"RETURNING id"

		var id uint64
		if err := tx.Get(&id, insertSQL, l.UserID, l.URL, l.ShortURL, l.ViewsLimit, l.Views, l.ExpiredAt, l.CreatedAt, l.DeletedAt); err != nil {
			return err
		}
		l.ID = id

		t.logger.Info("Link created",
			zap.Reflect("$.Name", l),
		)
		t.cache.Store(t.linksCacheKey(l.ID), l)
		return nil
	})
}

// UpdateLink updates Link in db and cache
func (t *Tables) UpdateLink(l *schema.Link) error {
	return t.db.Mutate(func(tx *dao.Tx) error {

		updateSQL := "UPDATE links " +
			"SET " +
			"user_id = $1, " +
			"url = $2, " +
			"short_url = $3, " +
			"views_limit = $4, " +
			"views = $5, " +
			"expired_at = $6, " +
			"created_at = $7, " +
			"deleted_at = $8 " +
			fmt.Sprintf("WHERE id = %d", l.ID)

		_, err := tx.Exec(updateSQL, l.UserID, l.URL, l.ShortURL, l.ViewsLimit, l.Views, l.ExpiredAt, l.CreatedAt, l.DeletedAt)
		if err != nil {
			return err
		}

		t.logger.Info("Link updated",
			zap.Reflect("$.Name", l),
		)
		t.cache.Store(t.linksCacheKey(l.ID), l)
		return nil
	})
}

func (t *Tables) linksCacheKey(id interface{}) string {
	b := bytes.Buffer{}
	b.WriteString("link")

	b.WriteString(fmt.Sprintf("_id=%v", id))

	return b.String()
}
