package models

import (
	"time"
)

var (
	_default_size = 30
)

var _default_hightlight = HighlightQuery{
	PreTag:  []string{"<mark class='highlight'>"},
	PostTag: []string{"<mark>"},
	Fields:  make(map[string]struct{}),
}

type Query struct {
	Query     QueryField      `json:"query"`
	Fields    []string        `json:"_source"`
	From      int             `json:"from"`
	Size      int             `json:"size"`
	Sort      []string        `json:"sort"`
	Highlight *HighlightQuery `json:"highlight,omitempty"`
}

func NewQuery() *Query {
	copyHighLight := _default_hightlight
	return &Query{
		Fields: make([]string, 0),
		Query: QueryField{
			Bool: BoolQuery{
				Must:        make([]any, 0),
				RangeFilter: make([]DateRange, 0),
			},
		},
		From:      0,
		Size:      _default_size,
		Sort:      []string{"-date"},
		Highlight: &copyHighLight,
	}
}

func (zb *Query) SetFrom(from int) *Query {
	zb.From = from
	return zb
}

func (zb *Query) SetSize(size int) *Query {
	zb.Size = size
	return zb
}

func (zb *Query) SetFields(fields []string) *Query {
	zb.Fields = fields
	return zb
}

func (zb *Query) SetHightlightFields(fields []string) *Query {
	highlightFields := make(map[string]struct{})

	for _, field := range fields {
		highlightFields[field] = struct{}{}
	}
	zb.Highlight.Fields = highlightFields
	return zb
}

func (zb *Query) SetQueryString(query string) *Query {
	queryString := SearchQueryString{
		QueryString: QueryString{
			Query: query,
		},
	}

	zb.Query.Bool.Must = []any{queryString}
	return zb
}

// if query is empty return same Query
func (zb *Query) AddQueryString(query string) *Query {
	if query == "" {
		return zb
	}

	queryString := SearchQueryString{
		QueryString: QueryString{
			Query: query,
		},
	}

	zb.Query.Bool.Must = append(zb.Query.Bool.Must, queryString)
	return zb
}

func (zb *Query) AddRangeFilter(dateRange DateRange) *Query {

	zb.Query.Bool.RangeFilter = append(zb.Query.Bool.RangeFilter, dateRange)
	return zb
}

func (zb *Query) AddMatchFieldFilter(key, value string) *Query {
	field := make(map[string]string, 0)
	field[key] = value
	matchField := SearchMatch{
		Match: field,
	}

	zb.Query.Bool.Must = append(zb.Query.Bool.Must, matchField)
	return zb
}

func (zb *Query) SetGetAllQuery() *Query {
	queryGetAll := SearchMatchAll{}

	zb.Query.Bool.Must = []any{queryGetAll}

	return zb
}

type QueryField struct {
	Bool BoolQuery `json:"bool"`
}

type BoolQuery struct {
	Must        []any       `json:"must"`
	RangeFilter []DateRange `json:"filter,omitempty"`
}

type SearchQueryString struct {
	QueryString QueryString `json:"query_string"`
}

type QueryString struct {
	Query string `json:"query"`
}

type SearchMatch struct {
	Match map[string]string `json:"match"`
}

type SearchMatchAll struct {
	All struct{} `json:"match_all"`
}

type DateRange struct {
	Range struct {
		Date DateInterval `json:"date"`
	} `json:"range"`
}

// type DateRange struct {
// 	Range struct {
// 		Fields map[string]DateInterval
// 	} `json:"range"`
// }

type DateInterval struct {
	GreatherThanOrEquals *time.Time `json:"gte,omitempty"`
	LessThanOrEquals     *time.Time `json:"lte,omitempty"`
	Format               string     `json:"format,omitempty"`
}

type HighlightQuery struct {
	PreTag  []string            `json:"pre_tags"`
	PostTag []string            `json:"post_tags"`
	Fields  map[string]struct{} `json:"fields"`
}
