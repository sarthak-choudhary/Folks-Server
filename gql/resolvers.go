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

	result, err = query.GetAllEvents(mongo.Session)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getAllUsers(_ graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}

	result, err = query.GetAllUsers(mongo.Session)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getAllSquads(_ graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}

	result, err = query.GetAllSquads(mongo.Session)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getEvent(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}
	var id string

	if p.Args["_id"] != nil {
		id = p.Args["_id"].(string)
	}

	result, err = query.GetEvent(id, mongo.Session)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getUser(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}
	var id string

	if p.Args["_id"] != nil {
		id = p.Args["_id"].(string)
	}

	result, err = query.GetUserByID(id, mongo.Session)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getSquad(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}
	var id string

	if p.Args["_id"] != nil {
		id = p.Args["_id"].(string)
	}

	result, err = query.GetSquad(id, mongo.Session)
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

	result, err = query.AddEvent(name, description, destination, locationLatitude, locationLongitude, datetime, hostedBy, participants, picturesUrls, mongo.Session)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func addSquad(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}
	var name, description string
	var groupImages []string
	var admins []primitive.ObjectID
	var members []primitive.ObjectID
	var invitesSent []primitive.ObjectID

	user := p.Context.Value("user").(*models.User)

	admins = append(admins, user.ID)

	if p.Args["name"] != nil {
		name = p.Args["name"].(string)
	}

	if p.Args["description"] != nil {
		description = p.Args["description"].(string)
	}

	if p.Args["groupImages"] != nil {
		groupImages = p.Args["groupImages"].([]string)
	}

	result, err = query.AddSquad(name, description, groupImages, admins, members, invitesSent, mongo.Session)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func updateEvent(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var result interface{}
	var name, description, destination string
	var locationLatitude, locationLongitude float32
	var datetime time.Time
	var participants []primitive.ObjectID
	var picturesUrls []string

	user := p.Context.Value("user").(*models.User)
	userID := user.ID
	id := p.Args["_id"].(string)

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

	result, err = query.UpdateEvent(id, name, description, destination, locationLatitude, locationLongitude, datetime, userID, participants, picturesUrls, mongo.Session)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func addFollower(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var username string

	user := p.Context.Value("user").(*models.User)
	userID := user.ID

	if p.Args["username"] != nil {
		username = p.Args["username"].(string)
	}

	result, err := query.FollowUser(userID, username, mongo.Session)
	if err != nil {
		return nil, err
	}
	return result, nil
}
