package gql

import (
	"log"

	"github.com/anshalshukla/events_mongodb/db"
	"github.com/graphql-go/graphql"
)

var s *db.MongoDB

// InitSchema - defines complete graphql schema
func InitSchema(d *db.MongoDB) graphql.Schema {
	s = d
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getAllEvents": &graphql.Field{
					Type:    graphql.NewList(EventType),
					Args:    graphql.FieldConfigArgument{},
					Resolve: getAllEvents,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"addEvent": &graphql.Field{
					Type: EventType,
					Args: graphql.FieldConfigArgument{
						"name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"destination": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"datetime": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"hostedBy": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"participants": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
						"picturesUrls": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
					},
					Resolve: addEvent,
				},
			},
		}),
		Types: []graphql.Type{ID},
	})
	if err != nil {
		log.Fatal(err)
	}
	return graphqlSchema
}
