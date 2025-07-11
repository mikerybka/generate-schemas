package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/mikerybka/util"
)

func main() {
	schemasURL := "http://localhost:3000"
	schemasDir := filepath.Join(util.HomeDir(), "schemas")
	schemas, err := listSchemas(schemasURL)
	if err != nil {
		panic(err)
	}
	for _, e := range schemas.Value {
		s, err := getSchema(schemasURL + "/" + e.Name)
		if err != nil {
			panic(err)
		}
		err = writeSchema(schemasDir, e.Name, s)
		if err != nil {
			panic(err)
		}
	}
}

func getSchema(addr string) (*Schema, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	res := &GetSchemaResponse{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, err
	}
	return res.Value, nil
}

func writeSchema(dir string, name string, schema *Schema) error {
	// TODO: Write Go file
	// TODO: Write TypeScript file
	b, _ := json.Marshal(schema)
	fmt.Println(string(b))
	return nil
}

type ListSchemasResponse struct {
	Value []DirEntry `json:"value"`
}

type GetSchemaResponse struct {
	Value *Schema `json:"value"`
}

type DirEntry struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Schema struct {
	Fields []Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func listSchemas(schemasURL string) (*ListSchemasResponse, error) {
	resp, err := http.Get(schemasURL)
	if err != nil {
		return nil, err
	}
	res := &ListSchemasResponse{}
	err = json.NewDecoder(resp.Body).Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
