package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-go/graphql"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

var speciesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "pokemon_species",
		// we define the name and the fields of our
		// object. In this case, we have one solitary
		// field that is of type string
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var pokemonType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "pokemon",
		// we define the name and the fields of our
		// object. In this case, we have one solitary
		// field that is of type string
		Fields: graphql.Fields{
			"entry_number": &graphql.Field{
				Type: graphql.Int,
			},
			"pokemon_species": &graphql.Field{
				Type: speciesType,
			},
		},
	},
)
var dataType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "data",
		Fields: graphql.Fields{

			"name": &graphql.Field{
				Type: graphql.String,
			},
			"pokemon": &graphql.Field{
				Type: graphql.NewList(pokemonType),
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"data": &graphql.Field{
				Type: dataType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					resp, err := hitAPI()
					if err != nil {
						return nil, err
					}
					return resp, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
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

//Helper function to import json from file to map
func hitAPI() (resp Response, err error) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		return resp, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		return resp, err
	}
	fmt.Println(string(responseData))
	_ = json.Unmarshal(responseData, &resp)
	return
}
