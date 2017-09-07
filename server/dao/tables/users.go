// Code generated by gen_schema_tables.go DO NOT EDIT.

package tables

import (
	"bytes"
	"fmt"

	"go.uber.org/zap"

	"github.com/ngalayko/url_shortner/server/dao"
	"github.com/ngalayko/url_shortner/server/schema"
)

// SelectUserById returns User from db or cache
func (t *Tables) SelectUserById(id uint64) (*schema.User, error) {
	return t.SelectUserByFields(map[string]interface{}{"id": id})
}

// SelectUserByIds returns Users from db or cache
func (t *Tables) SelectUserByFields(fields map[string]interface{}) (*schema.User, error) {

	if len(fields) == 0 {
		return nil, nil
	}

	if value, ok := t.cache.Load(t.usersCacheKey(fields)); ok {
		return value.(*schema.User), nil
	}

	b := bytes.Buffer{}
	b.WriteString("SELECT * " +
		"FROM users " +
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

	u := &schema.User{}
	if err := t.db.Get(u, b.String(), values...); err != nil {
		return nil, err
	}

	t.cache.Store(t.usersCacheKey(fields), u)
	return u, nil
}

// InsertUser inserts User in db and cache
func (t *Tables) InsertUser(u *schema.User) error {
	return t.db.Mutate(func(tx *dao.Tx) error {

		insertSQL := "INSERT INTO users " +
			"(first_name, last_name, created_at, deleted_at) " +
			"VALUES " +
			"($1, $2, $3, $4) " +
			"RETURNING id"

		var id uint64
		if err := tx.Get(&id, insertSQL, u.FirstName, u.LastName, u.CreatedAt, u.DeletedAt); err != nil {
			return err
		}
		u.ID = id

		t.logger.Info("User created",
			zap.Reflect("$.Name", u),
		)
		t.cache.Store(t.usersCacheKey(map[string]interface{}{"id": u.ID}), u)
		return nil
	})
}

// UpdateUser updates User in db and cache
func (t *Tables) UpdateUser(u *schema.User) error {
	return t.db.Mutate(func(tx *dao.Tx) error {

		updateSQL := "UPDATE users " +
			"SET " +
			"first_name = $1, " +
			"last_name = $2, " +
			"created_at = $3, " +
			"deleted_at = $4 " +
			fmt.Sprintf("WHERE id = %d", u.ID)

		_, err := tx.Exec(updateSQL, u.FirstName, u.LastName, u.CreatedAt, u.DeletedAt)
		if err != nil {
			return err
		}

		t.logger.Info("User updated",
			zap.Reflect("$.Name", u),
		)
		t.cache.Store(t.usersCacheKey(map[string]interface{}{"id": u.ID}), u)
		return nil
	})
}

func (t *Tables) usersCacheKey(fields map[string]interface{}) string {
	b := bytes.Buffer{}
	b.WriteString("user")

	for k, v := range fields {
		b.WriteString(fmt.Sprintf("_%s=%v", k, v))
	}

	return b.String()
}
