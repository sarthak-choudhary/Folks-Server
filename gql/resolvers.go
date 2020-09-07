package gql

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/anshalshukla/folks/db/models"
	"github.com/anshalshukla/folks/db/query"

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
	var name, description, destination string
	var locationLatitude, locationLongitude float32
	var datetime time.Time
	var participants []primitive.ObjectID
	var picturesUrls []string

	user := p.Context.Value("user").(*models.User)
	hostedBy := user.ID

	if p.Args["name"] != nil {
		name = p.Args["name"].(string)
	}

	if p.Args["description"] != nil {
		description = p.Args["description"].(string)
	}

	if p.Args["destination"] != nil {
		destination = p.Args["destination"].(string)
	}

	if p.Args["locationLatitude"] != nil {
		locationLatitude = p.Args["locationLatitude"].(float32)
	}

	if p.Args["locationLongitude"] != nil {
		locationLongitude = p.Args["locationLongitude"].(float32)
	}

	if p.Args["datetime"] != nil {
		datetime = p.Args["datetime"].(time.Time)
	}

	if p.Args["participant"] != nil {
		participants = p.Args["participant"].([]primitive.ObjectID)
	}

	if p.Args["picturesUrls"] != nil {
		picturesUrls = p.Args["picturesUrls"].([]string)
	}

	result, err = query.AddEvent(name, description, destination, locationLatitude, locationLongitude, datetime, hostedBy, participants, picturesUrls, mongo)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func updateEvents(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}
	var name, description, destination string
	var locationLatitude, locationLongitude float32
	var datetime time.Time
	var participants []primitive.ObjectID
	var picturesUrls []string

	user := p.Context.Value("user").(*models.User)
	loggedInUser := user.ID

	if p.Args["name"] != nil {
		name = p.Args["name"].(string)
	}

	if p.Args["description"] != nil {
		description = p.Args["description"].(string)
	}

	if p.Args["destination"] != nil {
		destination = p.Args["destination"].(string)
	}

	if p.Args["locationLatitude"] != nil {
		locationLatitude = p.Args["locationLatitude"].(float32)
	}

	if p.Args["locationLongitude"] != nil {
		locationLongitude = p.Args["locationLongitude"].(float32)
	}

	if p.Args["datetime"] != nil {
		datetime = p.Args["datetime"].(time.Time)
	}

	if p.Args["participant"] != nil {
		participants = p.Args["participant"].([]primitive.ObjectID)
	}

	if p.Args["picturesUrls"] != nil {
		picturesUrls = p.Args["picturesUrls"].([]string)
	}

}
