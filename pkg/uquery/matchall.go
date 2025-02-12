package uquery

import (
	"github.com/blugelabs/bluge"
	v1 "github.com/zinclabs/zinc/pkg/meta/v1"
)

func MatchAllQuery(iQuery *v1.ZincQuery) (bluge.SearchRequest, error) {
	dateQuery := bluge.NewDateRangeQuery(iQuery.Query.StartTime, iQuery.Query.EndTime).SetField("@timestamp")
	allquery := bluge.NewMatchAllQuery()
	query := bluge.NewBooleanQuery().AddMust(dateQuery).AddMust(allquery)

	searchRequest := buildRequest(iQuery, query)

	return searchRequest, nil
}
