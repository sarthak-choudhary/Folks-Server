package elasticsearch

import (
	"fmt"
	"os"

	elastic "github.com/olivere/elastic/v7"
)

//Model - Document to be stored in index
type Model struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Category    string `json:"category,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetESClient - gives the connection to elastic search server
func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(os.Getenv("ELASTICSEARCH_URL")), elastic.SetHealthcheck(false), elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}
	fmt.Println("ES initialized")

	return client, err
}

