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
	var locationLatitude, locationLongitude float64
	var datetime string
	var inviteList []primitive.ObjectID
	var picturesUrls []string
	var admins []primitive.ObjectID

	user := p.Context.Value("user").(*models.User)
	hostedBy := user.ID
	owner := user.FirstName + " " + user.LastName

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
		locationLatitude = p.Args["locationLatitude"].(float64)
	}

	if p.Args["locationLongitude"] != nil {
		locationLongitude = p.Args["locationLongitude"].(float64)
	}

	if p.Args["datetime"] != nil {
		datetime = p.Args["datetime"].(string)
	}

	if p.Args["inviteList"] != nil {
		inviteList = p.Args["inviteList"].([]primitive.ObjectID)
	}

	if p.Args["picturesUrls"] != nil {
		picturesUrls = p.Args["picturesUrls"].([]string)
	}

	if p.Args["admins"] != nil {
		admins = p.Args["admins"].([]primitive.ObjectID)
	}

	t, err := time.Parse("2006-01-02T15:04:05.000Z", datetime)

	if err != nil {
		return nil, err
	}

	admins = append(admins, hostedBy)
	result, err = query.AddEvent(name, description, destination, locationLatitude, locationLongitude, t, hostedBy, inviteList, picturesUrls, admins, owner, mongo.Session)
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
	var locationLatitude, locationLongitude float64
	var datetime string
	var t time.Time
	var picturesUrls []string

	user := p.Context.Value("user").(*models.User)
	id := p.Args["id"].(string)

	event, err := query.GetEvent(id, mongo.Session)
	if event.HostedBy != user.ID {
		err = errors.New("event can be updated by its owner only")
		return nil, err
	}

	if p.Args["name"] != nil {
		name = p.Args["name"].(string)
	} else {
		name = event.Name
	}

	if p.Args["description"] != nil {
		description = p.Args["description"].(string)
	} else {
		description = event.Description
	}

	if p.Args["destination"] != nil {
		destination = p.Args["destination"].(string)
	} else {
		destination = event.Destination
	}

	if p.Args["locationLatitude"] != nil {
		locationLatitude = p.Args["locationLatitude"].(float64)
	} else {
		locationLatitude = event.LocationLatitude
	}

	if p.Args["locationLongitude"] != nil {
		locationLongitude = p.Args["locationLongitude"].(float64)
	} else {
		locationLongitude = event.LocationLongitude
	}

	if p.Args["datetime"] != nil {
		datetime = p.Args["datetime"].(string)
	} else {
		t = event.Datetime
	}

	if p.Args["picturesUrls"] != nil {
		picturesUrls = p.Args["picturesUrls"].([]string)
	}

	if datetime != "" {
		t, err = time.Parse("2006-01-02T15:04:05.000Z", datetime)

		if err != nil {
			return nil, err
		}
	}

	result, err = query.UpdateEvent(event.ID, name, description, destination, locationLatitude, locationLongitude, t, picturesUrls, mongo.Session)
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

func declineRequest(p graphql.ResolveParams) (interface{}, error) {
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

	result, err := query.DeclineRequest(id, userID, mongo.Session)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func requestEvent(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var event_id string

	user := p.Context.Value("user").(*models.User)
	id := user.ID

	if p.Args["id"] != nil {
		event_id = p.Args["id"].(string)
	}

	event, err := query.GetEvent(event_id, mongo.Session)
	if err != nil {
		return nil, err
	}

	result, err := query.RequestEvent(id, event.ID, mongo.Session)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func acceptParticipants(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var event_id string
	var users []interface{}
	var users_id []primitive.ObjectID

	user := p.Context.Value("user").(*models.User)

	if p.Args["userID"] != nil {
		users = p.Args["userID"].([]interface{})
	}

	if p.Args["eventID"] != nil {
		event_id = p.Args["eventID"].(string)
	}

	event, err := query.GetEvent(event_id, mongo.Session)

	isAdmin := false
	for _, obj := range event.Admins {
		if obj == user.ID {
			isAdmin = true
			break
		}
	}

	if isAdmin == false {
		err = errors.New("Only admins can accept new participants")
		return nil, err
	}

	for _, user := range users {
		userID, err := primitive.ObjectIDFromHex(user.(string))

		if err == nil {
			users_id = append(users_id, userID)
		}
	}

	result, err := query.AcceptParticipants(event.ID, users_id, mongo.Session)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func declineParticipants(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var event_id string
	var users []interface{}
	var users_id []primitive.ObjectID

	user := p.Context.Value("user").(*models.User)

	if p.Args["userID"] != nil {
		users = p.Args["userID"].([]interface{})
	}

	if p.Args["eventID"] != nil {
		event_id = p.Args["eventID"].(string)
	}

	event, err := query.GetEvent(event_id, mongo.Session)

	isAdmin := false
	for _, obj := range event.Admins {
		if obj == user.ID {
			isAdmin = true
			break
		}
	}

	if isAdmin == false {
		err = errors.New("Only admins can accept new participants")
		return nil, err
	}

	for _, user := range users {
		userID, err := primitive.ObjectIDFromHex(user.(string))

		if err == nil {
			users_id = append(users_id, userID)
		}
	}

	result, err := query.DeclineParticipants(event.ID, users_id, mongo.Session)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func inviteParticipants(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var event_id string
	var users []interface{}
	var users_id []primitive.ObjectID

	user := p.Context.Value("user").(*models.User)

	if p.Args["userID"] != nil {
		users = p.Args["userID"].([]interface{})
	}

	if p.Args["eventID"] != nil {
		event_id = p.Args["eventID"].(string)
	}

	event, err := query.GetEvent(event_id, mongo.Session)

	isAdmin := false
	for _, obj := range event.Admins {
		if obj == user.ID {
			isAdmin = true
			break
		}
	}

	if isAdmin == false {
		err = errors.New("Only admins can invite new participants")
		return nil, err
	}

	for _, user := range users {
		userID, err := primitive.ObjectIDFromHex(user.(string))

		if err == nil {
			users_id = append(users_id, userID)
		}
	}

	result, err := query.InviteParticipants(event.ID, users_id, mongo.Session)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func assignAdmin(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var event_id string
	var admins []interface{}
	var admins_id []primitive.ObjectID

	user := p.Context.Value("user").(*models.User)

	if p.Args["admins"] != nil {
		admins = p.Args["admins"].([]interface{})
	}

	if p.Args["eventID"] != nil {
		event_id = p.Args["eventID"].(string)
	}

	event, err := query.GetEvent(event_id, mongo.Session)

	if err != nil {
		return nil, err
	}

	isAdmin := false
	for _, obj := range event.Admins {
		if obj == user.ID {
			isAdmin = true
			break
		}
	}

	if isAdmin == false {
		err = errors.New("Only admins can assign Admin status")
		return nil, err
	}

	for _, obj := range admins {
		adminID, err := primitive.ObjectIDFromHex(obj.(string))

		if err == nil {
			admins_id = append(admins_id, adminID)
		}
	}

	result, err := query.AssignAdmin(event.ID, admins_id, mongo.Session)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func acceptInvite(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var event_id string

	user := p.Context.Value("user").(*models.User)
	id := user.ID

	if p.Args["eventID"] != nil {
		event_id = p.Args["eventID"].(string)
	}

	eventID, err := primitive.ObjectIDFromHex(event_id)

	if err != nil {
		return nil, err
	}

	invited := false

	for _, obj := range user.InvitesReceived {
		if obj == eventID {
			invited = true
			break
		}
	}

	if invited != true {
		err = errors.New("User doesn't has such invite")
		return nil, err
	}

	result, err := query.AcceptInvite(id, eventID, mongo.Session)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func declineInvite(p graphql.ResolveParams) (interface{}, error) {
	var err error
	var event_id string

	user := p.Context.Value("user").(*models.User)
	id := user.ID

	if p.Args["eventID"] != nil {
		event_id = p.Args["eventID"].(string)
	}

	eventID, err := primitive.ObjectIDFromHex(event_id)

	if err != nil {
		return nil, err
	}

	invited := false

	for _, obj := range user.InvitesReceived {
		if obj == eventID {
			invited = true
			break
		}
	}

	if invited != true {
		err = errors.New("User doesn't has such invite")
		return nil, err
	}

	result, err := query.DeclineInvite(id, eventID, mongo.Session)

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
	var longitude float64
	var latitude float64
	var radius float64
	if rp.Args["locationLatitude"] != nil && rp.Args["locationLongitude"] != nil && rp.Args["radius"] != nil {
		longitude = rp.Args["locationLongitude"].(float64)
		latitude = rp.Args["locationLatitude"].(float64)
		radius = rp.Args["radius"].(float64)
	}
	result, err := util.NearbyEvents(mongo.Session, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getNearByEventsWithImages(rp graphql.ResolveParams) (interface{}, error) {
	var longitude float64
	var latitude float64
	var radius float64
	if rp.Args["locationLatitude"] != nil && rp.Args["locationLongitude"] != nil && rp.Args["radius"] != nil {
		longitude = rp.Args["locationLongitude"].(float64)
		latitude = rp.Args["locationLatitude"].(float64)
		radius = rp.Args["radius"].(float64)
	}
	result, err := util.NearbyEventsWithImages(mongo.Session, latitude, longitude, radius)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func changePassword(rp graphql.ResolveParams) (interface{}, error)	{
	var newPassword string
	var oldPassword string
	user := rp.Context.Value("user").(*models.User)
	id := user.ID

	if rp.Args["newPassword"] != nil && rp.Args["oldPassword"] != nil {
		newPassword = rp.Args["newPassword"].(string)
		oldPassword = rp.Args["oldPassword"].(string)
	}
	result, err := query.ChangePassword(oldPassword, newPassword, id, mongo.Session)
	if err != nil {
		return nil, err
	}
	return result, nil
}