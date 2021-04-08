package query

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/WeFolks/search_service/elasticsearch"

	"github.com/olivere/elastic/v7"
)

//InsertData - function to insert data into elastic database
func InsertData(ctx context.Context, esclient *elastic.Client, name, id, category, owner, description string, Type int) error {
	newEntry := elasticsearch.Model{
		ID:          id,
		Name:        name,
		Category:    category,
		Owner:       owner,
		Type:        Type,
		Description: description,
	}

	dataJSON, err := json.Marshal(newEntry)
	js := string(dataJSON)
	_, err = esclient.Index().Index("search_data").BodyJson(js).Do(ctx)

	if err != nil {
		return err
	}

	fmt.Println("[Elastic][InsertProduct]Insertion Successful")

	return nil
}
