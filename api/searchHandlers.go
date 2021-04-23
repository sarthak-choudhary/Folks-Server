package api

import (
	"context"
	"encoding/json"
	"fmt"
	q "github.com/wefolks/backend/elasticsearch/query"
	"net/http"
	"time"

	"github.com/olivere/elastic/v7"
)

//Body - struct for request body
type Body struct {
	Name     string `json:"name"`
	Category string `json:"category,omitempty"`
}

//Result - struct for response
type Result struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
	Type        int32  `json:"type"`
}

//GetData - http endpoint for obtaining search results
func GetData(esclient *elastic.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body Body
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			println("i was here")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			payload := struct {
				Error string `json:"error"`
			}{Error: "Invalid or Incomplete fields"}

			json.NewEncoder(w).Encode(payload)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		result, err := q.Data(ctx, esclient, body.Name, body.Category)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var data []Result

		for _, element := range result {
			item := Result{
				ID:          element.ID,
				Name:        element.Name,
				Category:    element.Category,
				Owner:       element.Owner,
				Description: element.Description,
				Type:        int32(element.Type),
			}

			data = append(data, item)
		}

		payload := struct {
			Data []Result `json:"data"`
		}{
			Data: data,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(payload)

	})
}