//go:generate go run gen_names_array.go

package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"text/template"
)

var (
	tablesArrayTemplate = regexp.MustCompile(`(?s)string\n}{.*?}[^),]`)
	structs             = []strGen{}
)

const (
	generateSchemaTabelsFile = "gen_schema_tables.go"

	schemaPath = "../schema"

	tablesTemplate = `}{
	// this array generated by gen_names_array.go DO NOT EDIT.
	{{- range . }}
		{reflect.TypeOf(schema.{{ .Name }}{}), "{{ .NamePlural }}"},
	{{- end }}
	}
`
)

type tableData struct {
	Name       string
	NamePlural string
	SqlQueries []string
}

func main() {

	if err := getSchemaNamesByAst(); err != nil {
		panic(err)
	}

	tablesTemplate := template.Must(template.New("enums_template").Parse(tablesTemplate))

	tablesData := make([]tableData, 0, len(structs))
	for _, str := range structs {
		tablesData = append(tablesData, tableData{
			Name:       str.Name,
			NamePlural: strings.ToLower(str.Name) + "s",
		})
	}

	tablesBuffer := &bytes.Buffer{}
	if err := tablesTemplate.Execute(tablesBuffer, tablesData); err != nil {
		log.Fatal(err)
	}

	generateSchemaTablesData, err := ioutil.ReadFile(generateSchemaTabelsFile)
	if err != nil {
		log.Fatal(err)
	}

	generateSchemaTablesData = tablesArrayTemplate.ReplaceAll(generateSchemaTablesData, tablesBuffer.Bytes())

	err = ioutil.WriteFile(generateSchemaTabelsFile, generateSchemaTablesData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

type strGen struct {
	Name string
}

type VisitorFunc func(n ast.Node) ast.Visitor

func (f VisitorFunc) Visit(n ast.Node) ast.Visitor { return f(n) }

func getSchemaNamesByAst() error {

	fs := token.NewFileSet()

	pkgs, err := parser.ParseDir(fs, schemaPath, nil, 0)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		ast.Walk(VisitorFunc(findTypes), pkg)
	}

	return nil
}

func findTypes(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.Package:
		return VisitorFunc(findTypes)
	case *ast.File:
		return VisitorFunc(findTypes)
	case *ast.GenDecl:
		if n.Tok == token.TYPE {
			return VisitorFunc(findTypes)
		}
	case *ast.TypeSpec:
		structs = append(structs, strGen{
			Name: n.Name.Name,
		})
	}
	return nil
}
