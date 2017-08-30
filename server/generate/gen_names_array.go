//go:generate go run gen_names_array.go

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

var (
	schemaDefinitionRegex = regexp.MustCompile(`type (?P<name>\w+) struct {`)
	tablesArrayTemplate   = regexp.MustCompile(`(?s)string\n}{.*?}[^),]`)
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
}

func main() {
	names, err := getSchemaNames()
	if err != nil {
		log.Fatal(err)
	}

	tablesTemplate := template.Must(template.New("enums_template").Parse(tablesTemplate))

	tablesData := make([]tableData, 0, len(names))
	for _, name := range names {
		tablesData = append(tablesData, tableData{
			Name:       name,
			NamePlural: strings.ToLower(name) + "s",
		})
	}

	tablesBuffer := &bytes.Buffer{}
	if err := tablesTemplate.Execute(tablesBuffer, tablesData); err != nil {
		log.Fatal(err)
	}

	generateSchemaTabelsData, err := ioutil.ReadFile(generateSchemaTabelsFile)
	if err != nil {
		log.Fatal(err)
	}

	generateSchemaTabelsData = tablesArrayTemplate.ReplaceAll(generateSchemaTabelsData, tablesBuffer.Bytes())

	err = ioutil.WriteFile(generateSchemaTabelsFile, generateSchemaTabelsData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func getSchemaNames() ([]string, error) {
	d, err := os.Open(schemaPath)
	if err != nil {
		return nil, err
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, file := range files {
		name, err := getSchemaName(schemaPath + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}

	return names, nil
}

func getSchemaName(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	match := schemaDefinitionRegex.FindSubmatch(data)

	if len(match) < 2 {
		return "", fmt.Errorf("probles with file: %s", filePath)
	}

	return string(match[1]), nil
}
