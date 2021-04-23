package user_queries

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/wefolks/backend/db/models"
	query2 "github.com/wefolks/backend/elasticsearch/query"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// AddUser adds a new user-queries to the database
func AddUser(user *models.User, client *mongo.Client, ec *elastic.Client) (string, error) {
	collection := client.Database("folks").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
		return "", err
	}

	id := fmt.Sprintf("%s", insertResult.InsertedID)

	//var conn *g.ClientConn
	//conn, err = g.Dial("3.142.74.30:9000", g.WithInsecure())
	//
	//if err != nil {
	//	log.Fatalf("Object Could not be added to elasticsearch database")
	//	return id, err
	//}
	//
	//defer conn.Close()
	//
	//c := grpc.NewSearchServiceClient(conn)
	//
	//item := grpc.Item{
	//	Id:          id,
	//	Name:        user-queries.FirstName + " " + user-queries.LastName,
	//	Owner:       "",
	//	Category:    "",
	//	Description: "",
	//	Type:        0,
	//}
	//
	//_, err = c.AddItem(context.Background(), &item)
	//
	//if err != nil {
	//	log.Fatalf("Object Could not be added to elasticsearch database")
	//	return id, err
	//}
	userFullName := user.FirstName+" "+user.LastName
	err = query2.InsertData(context.TODO(), ec, userFullName, id, "", "", "", 0)

	id = id[10:34]
	return id, nil
}
