package query

import (
	"context"
	"fmt"
	"github.com/anshalshukla/folks/db/models"
	"github.com/anshalshukla/folks/pkg/elasticsearch/query"
	"github.com/olivere/elastic/v7"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// AddUser adds a new user to the database
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
	//	Name:        user.FirstName + " " + user.LastName,
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
	err = query.InsertData(context.TODO(), ec, userFullName, id, "", "", "", 0)

	id = id[10:34]
	return id, nil
}
