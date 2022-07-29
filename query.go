package main

import (
	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"kanto": &graphql.Field{
				Type: pokemonDataType,
				Args: graphql.FieldConfigArgument{
					"param1": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					param, ok := p.Args["param1"].(string)
					if !ok {
						param = ""
					}

					resp, err := hitAPI("kanto", param)
					if err != nil {
						return nil, err
					}
					return resp, nil
				},
			},
			"hoenn": &graphql.Field{
				Type: pokemonDataType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					param, ok := p.Args["param1"].(string)
					if !ok {
						param = ""
					}
					resp, err := hitAPI("hoenn", param)
					if err != nil {
						return nil, err
					}
					return resp, nil
				},
			},
			"product": &graphql.Field{
				Type:        productType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// Find product
						for _, product := range products {
							if int(product.ID) == id {
								return product, nil
							}
						}
					}
					return nil, nil
				},
			},
			/* Get (read) product list
			   http://localhost:8080/product?query={list{id,name,info,price}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get product list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return products, nil
				},
			},
		},
	})
