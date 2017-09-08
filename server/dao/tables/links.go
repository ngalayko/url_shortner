// Code generated by gen_schema_tables.go DO NOT EDIT.

package tables

import (
	"bytes"
	"fmt"

	"go.uber.org/zap"

	"github.com/ngalayko/url_shortner/server/dao"
	"github.com/ngalayko/url_shortner/server/schema"
)

// SelectLinkById returns Link from db or cache
func (t *Tables) SelectLinkById(id uint64) (*schema.Link, error) {
	return t.SelectLinkByFields(map[string]interface{}{"id": id})
}

// SelectLinkByIds returns Links from db or cache
func (t *Tables) SelectLinkByFields(fields map[string]interface{}) (*schema.Link, error) {

	if len(fields) == 0 {
		return nil, nil
	}

	if value, ok := t.cache.Load(t.linksCacheKey(fields)); ok {
		return value.(*schema.Link), nil
	}

	b := bytes.Buffer{}
	b.WriteString("SELECT * "+
		"FROM links "+
		"WHERE ")

	i := 1
	values := []interface{}{}
	for k, v := range fields {
		values = append(values, v)

		if i > 1 {
			b.WriteString(" AND \n")
		}

		b.WriteString(fmt.Sprintf("%s = $%d", k, i))
	}

	l := &schema.Link{}
	if err := t.db.Get(l, b.String(), values...); err != nil {
		return nil, err
	}

	t.cache.Store(t.linksCacheKey(fields), l)
	t.cache.Store(t.linksCacheKey(map[string]interface{}{"id": l.ID}), l)
	return l, nil
}

// InsertLink inserts Link in db and cache
func (t *Tables) InsertLink(l *schema.Link) error {
	return t.db.Mutate(func(tx *dao.Tx) error {

		insertSQL := "INSERT INTO links " +
			"(user_id, url, short_url, views, expired_at, created_at, deleted_at) " +
			"VALUES " +
			"($1, $2, $3, $4, $5, $6, $7) " +
			"RETURNING id"

		var id uint64
		if err := tx.Get(&id, insertSQL, l.UserID, l.URL, l.ShortURL, l.Views, l.ExpiredAt, l.CreatedAt, l.DeletedAt); err != nil {
			return err
		}
		l.ID = id

		t.logger.Info("Link created",
			zap.Reflect("$.Name", l),
		)
		t.cache.Store(t.linksCacheKey(map[string]interface{}{"id": l.ID}), l)
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
			"views = $4, " +
			"expired_at = $5, " +
			"created_at = $6, " +
			"deleted_at = $7 " +
			fmt.Sprintf("WHERE id = %d", l.ID)

		_, err := tx.Exec(updateSQL, l.UserID, l.URL, l.ShortURL, l.Views, l.ExpiredAt, l.CreatedAt, l.DeletedAt)
		if err != nil {
			return err
		}

		t.logger.Info("Link updated",
			zap.Reflect("$.Name", l),
		)
		t.cache.Store(t.linksCacheKey(map[string]interface{}{"id": l.ID}), l)
		return nil
	})
}

func (t *Tables) linksCacheKey(fields map[string]interface{}) string {
	b := bytes.Buffer{}
	b.WriteString("link")

	for k, v := range fields {
		b.WriteString(fmt.Sprintf("_%s=%v", k, v))
	}

	return b.String()
}
