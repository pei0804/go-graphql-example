package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

// https://qiita.com/takochuu/items/96af2ff573ca243b0174

var q graphql.ObjectConfig = graphql.ObjectConfig{
	Name: "query",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: resolveID,
		},
		"name": &graphql.Field{
			Type:    graphql.String,
			Resolve: resolveName,
		},
	},
}

var schemaConfig graphql.SchemaConfig = graphql.SchemaConfig{
	Query: graphql.NewObject(q),
}

var schema, _ = graphql.NewSchema(schemaConfig)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(r.Errors) > 0 {
		fmt.Printf("エラーがあるよ: %v", r.Errors)
	}

	j, _ := json.Marshal(r)
	fmt.Printf("%s \n", j)
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	bufBody := new(bytes.Buffer)
	bufBody.ReadFrom(r.Body)
	query := bufBody.String()

	result := executeQuery(query, schema)
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func resolveID(p graphql.ResolveParams) (interface{}, error) {
	return p.Args[curl -X POST -d '{ id(id: 100), name }' http://localhost:8080/"id"], nil
}
func resolveName(p graphql.ResolveParams) (interface{}, error) {
	return "hoge", nil
}
