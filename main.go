package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func hitAPI(region string, param string) (resp Response, err error) {

	url := fmt.Sprintf("http://pokeapi.co/api/v2/pokedex/%s/%s", region, param)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return resp, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		return resp, err
	}
	_ = json.Unmarshal(responseData, &resp)
	resp.URLHit = url
	return
}
