package gql

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

func _validate(value string) error {
	if len(value) < 3 {
		return errors.New("The minimum length required is 3")
	}
	return nil
}

var ID = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "ID",
	Description: "The `id` scalar type represents a ID Object.",
	Serialize: func(value interface{}) interface{} {
		return value
	},
	ParseValue: func(value interface{}) interface{} {
		var err error
		switch value.(type) {
		case string:
			err = _validate(value.(string))
		default:
			err = errors.New("Must be of type string")
		}
		if err != nil {
			log.Fatal(err)
		}
		return value
	},
	ParseLiteral: func(valueAst ast.Value) interface{} {
		if valueAst.GetKind() == kinds.StringValue {
			err := _validate(valueAst.GetValue().(string))
			if err != nil {
				log.Fatal(err)
			}
			return valueAst
		} else {
			log.Fatal("Must be of type string.")
		}
	},
})

var EventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"destination": &graphql.Field{
			Type: graphql.String,
		},
		"datetime": &graphql.Field{
			Type: graphql.String,
		},
		"hostedBy": &graphql.Field{
			Type: ID,
		},
		"participants": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"picturesUrls": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})
