package query

import (
	"context"
	"encoding/json"

	"github.com/WeFolks/search_service/elasticsearch"

	"github.com/olivere/elastic/v7"
)

//Data - used to query data from elastic seacrh database
func Data(ctx context.Context, esclient *elastic.Client, name, category string) ([]elasticsearch.Model, error) {
	if category != "" {
		return queryEvent(ctx, esclient, name, category)
	}

	return queryUserAndSquad(ctx, esclient, name)
}

func queryEvent(ctx context.Context, esclient *elastic.Client, name, category string) ([]elasticsearch.Model, error) {
	multiQuery := elastic.NewMultiMatchQuery(name, "name", "description").Type("phrase_prefix")
	matchQuery := elastic.NewMatchQuery("category", category)
	query := elastic.NewBoolQuery().Must(multiQuery, matchQuery)
	searchResult, err := esclient.Search().Index("search_data").Query(query).Do(ctx)
	result := []elasticsearch.Model{}
	if err != nil {
		return result, err
	}
	for _, hit := range searchResult.Hits.Hits {
		var response elasticsearch.Model
		json.Unmarshal(hit.Source, &response)
		result = append(result, response)
	}

	return result, nil
}

func queryUserAndSquad(ctx context.Context, esclient *elastic.Client, name string) ([]elasticsearch.Model, error) {
	multiQuery := elastic.NewMultiMatchQuery(name, "name", "owner").Type("phrase_prefix")

	searchResult, err := esclient.Search().Index("search_data").Query(multiQuery).Do(ctx)

	result := []elasticsearch.Model{}

	if err != nil {
		return result, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var response elasticsearch.Model
		json.Unmarshal(hit.Source, &response)

		result = append(result, response)
	}

	return result, nil
}
