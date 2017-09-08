//go:generate go run gen_schema_tables.go common.go

package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/ngalayko/url_shortner/server/schema"
)

type templateData struct {
	Name         string
	TableName    string
	Fields       []string
	DbFields     []string
	UniqueFields []string
}

func main() {

	tables := []struct {
		typ       reflect.Type
		tableName string
	}{
		// this array generated by gen_names_array.go DO NOT EDIT.
		{reflect.TypeOf(schema.Link{}), "links"},
		{reflect.TypeOf(schema.User{}), "users"},
	}

	for _, table := range tables {
		if err := generate(table.typ, table.tableName); err != nil {
			panic(err)
		}
	}
}

func generate(typ reflect.Type, tableName string) error {

	file, err := os.Create(filePath + tableName + ".go")
	if err != nil {
		return fmt.Errorf("error opening file %s: %s", typ.Name(), err)
	}
	defer file.Close()

	data := templateData{
		Name:      typ.Name(),
		TableName: tableName,
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		dbTag := getTag(field, "db")

		if getTag(field, "unique") == "true" {
			data.UniqueFields = append(data.UniqueFields, field.Name)
		}

		if dbTag == idTag {
			continue
		}

		data.DbFields = append(data.DbFields, dbTag)
		data.Fields = append(data.Fields, field.Name)
	}

	t := template.Must(template.New(typ.Name()).Funcs(getTemplateFuncs()).Parse(fileTemplate))
	if err := t.Execute(file, data); err != nil {
		return fmt.Errorf("error executing template %s: %s", t.Name(), err)
	}

	return nil
}

func getTag(f reflect.StructField, tagName string) string {
	tag := f.Tag.Get(tagName)

	return strings.Split(tag, ",")[0]
}

const (
	idTag        = "id"
	filePath     = "../dao/tables/"
	fileTemplate = `// Code generated by gen_schema_tables.go DO NOT EDIT.

package tables

import (
	"bytes"
	"fmt"

	"go.uber.org/zap"

	"github.com/ngalayko/url_shortner/server/dao"
	"github.com/ngalayko/url_shortner/server/schema"
)

// Get{{ $.Name }}ById returns {{ $.Name }} from db or cache
func (t *Tables) Get{{ $.Name }}ById(id uint64) (*schema.{{ $.Name }}, error) {
	return t.Get{{ $.Name }}ByFields(map[string]interface{}{"id": id})
}

// Get{{ $.Name }}ByFields returns {{ $.Name }}s from db or cache
func (t *Tables) Get{{ $.Name }}ByFields(fields map[string]interface{}) (*schema.{{ $.Name }}, error) {

	if len(fields) == 0 {
		return nil, nil
	}

	if value, ok := t.cache.Load(t.{{ $.TableName }}CacheKey(fields)); ok {
		return value.(*schema.{{ $.Name }}), nil
	}

	b := bytes.Buffer{}
	b.WriteString("SELECT * " +
		"FROM {{ $.TableName }} " +
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

	{{ alias $.Name }} := &schema.{{ $.Name }}{}
	if err := t.db.Get({{ alias $.Name }}, b.String(), values...); err != nil {
		return nil, err
	}

	t.cache.Store(t.{{ $.TableName }}CacheKey(fields), {{ alias $.Name }})
	t.cache.Store(t.{{ $.TableName }}CacheKey(map[string]interface{}{"id": {{ alias $.Name }}.ID}), {{ alias $.Name }})
	return {{ alias $.Name }}, nil
}

// Insert{{ $.Name }} inserts {{ $.Name }} in db and cache
func (t *Tables) Insert{{ $.Name }}({{ alias $.Name }} *schema.{{ $.Name }}) error {
	return t.db.Mutate(func(tx *dao.Tx) error {

		insertSQL := "INSERT INTO {{ $.TableName }} " +
			"({{ head .DbFields}}{{ range tail .DbFields }}, {{ . }}{{ end }}) " +
			"VALUES " +
			"($1{{ range $index, $element := tail .Fields }}, ${{ sum $index 2 }}{{ end }}) " +
			"RETURNING id"

		var id uint64
		if err := tx.Get(&id, insertSQL, {{ alias $.Name }}.{{ head .Fields }}{{ range tail .Fields }}, {{ alias $.Name }}.{{ . }}{{ end }}); err != nil {
			return err
		}
		{{ alias $.Name }}.ID = id

		t.logger.Info("{{ $.Name }} created",
			zap.Reflect("$.Name", {{ alias $.Name }}),
		)
		t.cache.Store(t.{{ $.TableName }}CacheKey(map[string]interface{}{"id": {{ alias $.Name }}.ID}), {{ alias $.Name }})
		return nil
	})
}

// Update{{ $.Name }} updates {{ $.Name }} in db and cache
func (t *Tables) Update{{ $.Name }}({{ alias $.Name }} *schema.{{ $.Name }}) error {
	return t.db.Mutate(func(tx *dao.Tx) error {

		updateSQL := "UPDATE {{ $.TableName }} " +
			"SET " +
			{{ range $index, $element := body $.DbFields }}"{{ $element }} = ${{ sum $index 1 }}, " +
			{{ end }}"{{ last .DbFields }} = ${{ len .Fields }} " +
			fmt.Sprintf("WHERE id = %d", {{ alias $.Name }}.ID)

		_, err := tx.Exec(updateSQL, {{ alias $.Name }}.{{ head .Fields }}{{ range tail .Fields }}, {{ alias $.Name }}.{{ . }}{{ end }})
		if err != nil {
			return err
		}

		t.logger.Info("{{ $.Name }} updated",
			zap.Reflect("$.Name", {{ alias $.Name }}),
		)
		t.cache.Store(t.{{ $.TableName }}CacheKey(map[string]interface{}{"id": {{ alias $.Name }}.ID}), {{ alias $.Name }})
		return nil
	})
}

func (t *Tables) {{ $.TableName }}CacheKey(fields map[string]interface{}) string {
	b := bytes.Buffer{}
	b.WriteString("{{ underscore $.Name }}")

	for k, v := range fields {
		b.WriteString(fmt.Sprintf("_%s=%v", k, v))
	}

	return b.String()
}
`
)
