package gql

import (
	"errors"
	"time"

	"github.com/anshalshukla/folks/util"

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

func followUser(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var user_id string

	user := p.Context.Value("user").(*models.User)
	id := user.ID

	if p.Args["id"] != nil {
		user_id = p.Args["id"].(string)
	}

	userID, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return nil, err
	}

	var following bool
	following = false

	for _, obj := range user.Following {
		if obj == userID {
			following = true
			break
		}
	}

	if following == true {
		err = errors.New("User is already following this account")
		return nil, err
	}

	result, err := query.FollowUser(id, userID, mongo.Session)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func acceptRequest(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var user_id string

	user := p.Context.Value("user").(*models.User)
	id := user.ID
	if p.Args["id"] != nil {
		user_id = p.Args["id"].(string)
	}

	userID, err := primitive.ObjectIDFromHex(user_id)

	if err != nil {
		return nil, err
	}

	var requested bool
	requested = false

	for _, obj := range user.RequestsReceived {
		if obj == userID {
			requested = true
			break
		}
	}

	if requested != true {
		err = errors.New("User doesn't has such follow request")
		return nil, err
	}

	result, err := query.AcceptRequest(id, userID, mongo.Session)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func myProfile(p graphql.ResolveParams) (interface{}, error) {

	user := p.Context.Value("user").(*models.User)
	return user, nil
}

func getNearByEvents(rp graphql.ResolveParams) (interface{}, error) {
	var longitude float32
	var latitude float32
	var radius float32
	if rp.Args["locationLatitude"] != nil && rp.Args["locationLongitude"] != nil && rp.Args["radius"] != nil {
		longitude = rp.Args["locationLongitude"].(float32)
		latitude = rp.Args["locationLatitude"].(float32)
		radius = rp.Args["locationLatitude"].(float32)
	}
	result, err := util.NearbyEvents(mongo.Session, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getNearByEventsWithImages(rp graphql.ResolveParams) (interface{}, error) {
	var longitude float32
	var latitude float32
	var radius float32
	if rp.Args["locationLatitude"] != nil && rp.Args["locationLongitude"] != nil && rp.Args["radius"] != nil {
		longitude = rp.Args["locationLongitude"].(float32)
		latitude = rp.Args["locationLatitude"].(float32)
		radius = rp.Args["radius"].(float32)
	}
	result, err := util.NearbyEventsWithImages(mongo.Session, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}
	return result, nil
}
