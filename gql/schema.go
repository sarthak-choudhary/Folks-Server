package gql

import (
	"log"

	"github.com/anshalshukla/folks/db"
	"github.com/graphql-go/graphql"
)

var mongo *db.MongoDB

// InitSchema - defines complete graphql schema
func InitSchema(d *db.MongoDB) graphql.Schema {
	mongo = d
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"getAllEvents": &graphql.Field{
					Type:    graphql.NewList(EventType),
					Args:    graphql.FieldConfigArgument{},
					Resolve: getAllEvents,
				},
				"getAllUsers": &graphql.Field{
					Type:    graphql.NewList(UserType),
					Args:    graphql.FieldConfigArgument{},
					Resolve: getAllUsers,
				},
				"getEventByID": &graphql.Field{
					Type: graphql.NewList(EventType),
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: getEvent,
				},
				"getUserByID": &graphql.Field{
					Type: graphql.NewList(UserType),
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
					Resolve: getUser,
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
						"locationLatitude": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"locationLongitude": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"datetime": &graphql.ArgumentConfig{
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
				"updateEvent": &graphql.Field{
					Type: UserType,
					Args: graphql.FieldConfigArgument{
						"_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"description": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"destination": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"locationLatitude": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"locationLongitude": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"datetime": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"participants": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
						"picturesUrls": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
						},
					},
					Resolve: updateEvent,
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
