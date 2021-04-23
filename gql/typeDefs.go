package gql

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
)

//ID validation
func _validate(value string) error {
	if len(value) == 24 {
		return errors.New("The minimum length required is 3")
	}
	return nil
}

//ID scalar added
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
		}
		log.Fatal("Must be of type string.")
		return nil
	},
})

//UserType scalar added
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "user-queries",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: ID,
		},
		"firstName": &graphql.Field{
			Type: graphql.String,
		},
		"lastName": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		"password": &graphql.Field{
			Type: graphql.String,
		},
		"phoneNo": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"bio": &graphql.Field{
			Type: graphql.String,
		},
		"age": &graphql.Field{
			Type: graphql.Int,
		},
		"gender": &graphql.Field{
			Type: graphql.Int,
		},
		"isPublic": &graphql.Field{
			Type: graphql.Boolean,
		},
		"interests": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"isComplete": &graphql.Field{
			Type: graphql.Boolean,
		},
		"following": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"followedByCount": &graphql.Field{
			Type: graphql.Int,
		},
		"events": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"picturesUrls": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"requestsSent": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"requestsReceived": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"invitesSent": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"invitesReceived": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"fcmToken":	&graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

//EventType scalar added
var EventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "event-queries",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			Type: ID,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"category": &graphql.Field{
			Type: graphql.Int,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"destination": &graphql.Field{
			Type: graphql.String,
		},
		"locationLatitude": &graphql.Field{
			Type: graphql.Float,
		},
		"locationLongitude": &graphql.Field{
			Type: graphql.Float,
		},
		"datetime": &graphql.Field{
			Type: graphql.String,
		},
		"createdOn": &graphql.Field{
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
		"waitlist": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"admins": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"invitelist": &graphql.Field{
			Type: graphql.NewList(ID),
		},
	},
})

//SquadType scalar added
var SquadType = graphql.NewObject(graphql.ObjectConfig{
	Name: "squad-queries",
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
		"groupImages": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"admins": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"members": &graphql.Field{
			Type: graphql.NewList(ID),
		},
		"invitesSent": &graphql.Field{
			Type: graphql.NewList(ID),
		},
	},
})
