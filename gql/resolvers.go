package gql

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/anshalshukla/events_mongodb/db/query"
	"github.com/graphql-go/graphql"
)

//Query
func getAllEvents(_ graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}

	result, err = query.GetAllEvents(mongo)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//Mutation
func addEvent(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}

	name := p.Args["name"].(string)
	description := p.Args["description"].(string)
	destination := p.Args["destination"].(string)
	locationLatitude := p.Args["locationLatitude"].(float32)
	locationLongitude := p.Args["locationLongitude"].(float32)
	datetime := p.Args["datetime"].(time.Time)
	hostedBy := p.Args["hostedBy"].(primitive.ObjectID)
	participants := p.Args["participant"].([]primitive.ObjectID)
	picturesUrls := p.Args["picturesUrls"].([]string)

	result, err = query.AddEvent(name, description, destination, locationLatitude, locationLongitude, datetime, hostedBy, participants, picturesUrls, mongo)
	if err != nil {
		return nil, err
	}
	return result, nil
}
