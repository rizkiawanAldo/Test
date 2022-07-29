package main

import (
	"github.com/graphql-go/graphql"
)

// A Response struct to map the Entire Response
type Response struct {
	URLHit       string         `json:"url_hitted"`
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	Pokemon      []Pokemon      `json:"pokemon_entries"`
	IsMainSeries bool           `json:"is_main_series"`
	VersionGroup []VersionGroup `json:"version_groups"`
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
type VersionGroup struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var versionGroupType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "version_groups",
		// we define the name and the fields of our
		// object. In this case, we have one solitary
		// field that is of type string
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"url": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var speciesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "pokemon_species",
		// we define the name and the fields of our
		// object. In this case, we have one solitary
		// field that is of type string

		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				// Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// 	for _, val := range pokemonDataType.Fields() {
				// 		if val.Type == params.Args["name"] {
				// 			return val.Type, nil
				// 		}
				// 		return val.Type, nil
				// 	}
				// 	return nil, nil
				// },
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
var pokemonDataType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "region",
		Fields: graphql.Fields{
			"url_hitted": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"pokemon": &graphql.Field{
				Type: graphql.NewList(pokemonType),
			},
			"is_main_series": &graphql.Field{
				Type: graphql.Boolean,
			},
			"version_groups": &graphql.Field{
				Type: graphql.NewList(versionGroupType),
			},
		},
	},
)

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"info": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Info  string  `json:"info,omitempty"`
	Price float64 `json:"price"`
}
