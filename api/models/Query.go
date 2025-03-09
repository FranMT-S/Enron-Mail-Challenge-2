package models

import (
	"time"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/constants"
)

var (
	_default_size = constants.PAGINATOR_MINSIZE
)

// default structure of highlight that return the highlight between mark tag
var _default_hightlight = HighlightQuery{
	PreTag:  []string{"<mark class='highlight'>"},
	PostTag: []string{"<mark>"},
	Fields:  make(map[string]struct{}),
}

// Operator AND, OR
type IOperator string

var (
	AND IOperator = "AND"
	OR  IOperator = "OR"
)

/*
Represents a query in ZincSearch.

- Query stores different filters for the query, including Must, Should, Not, and Range filters.
- Fields contains the list of fields used to search for values. If empty, the search is performed across all fields.
- From and size are used together to create pagination.
- Sort contains the list of field names used to sort the results.

Must information in [Go Documentation](https://zincsearch-docs.zinc.dev/api-es-compatible/search/search/))
*/
type Query struct {
	Query     QueryField      `json:"query"`
	Fields    []string        `json:"_source"`
	From      int             `json:"from"`
	Size      int             `json:"size"`
	Sort      []string        `json:"sort"`
	Highlight *HighlightQuery `json:"highlight,omitempty"`
}

// Return a new instance of Query of Zincsearch with fields initialized
func NewQuery() *Query {
	copyHighLight := _default_hightlight
	return &Query{
		Fields: make([]string, 0),
		Query: QueryField{
			Bool: BoolQuery{
				Must:        make([]any, 0),
				Should:      make([]any, 0),
				Not:         make([]any, 0),
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

func (zb *Query) SetSort(fields []string) *Query {
	zb.Sort = fields
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

// Create a new SearchQueryString and set how the unique value y the Must Queries list.
func (zb *Query) SetQueryString(query string) *Query {
	queryString := SearchQueryString{
		QueryString: QueryString{
			Query: query,
		},
	}

	zb.Query.Bool.Must = []any{queryString}
	return zb
}

// Create a SearchQueryString from queryString then is added in the Must Queries List
// if querystring is empty so return the same Query
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

func (zb *Query) AddMatchFieldFilter(key, value string, operator IOperator, isExclusionFilter bool) *Query {
	field := make(map[string]SearchMatchField, 0)
	field[key] = SearchMatchField{
		Operator:    operator,
		QueryString: value,
	}

	matchField := SearchMatch{
		Match: field,
	}

	if isExclusionFilter {
		zb.Query.Bool.Not = append(zb.Query.Bool.Must, matchField)
	} else {
		zb.Query.Bool.Must = append(zb.Query.Bool.Must, matchField)
	}
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

/*
BoolQuery represents a boolean query containing the following conditions:

  - Must: A list of conditions that must be satisfied by the query.
  - Should: A list of conditions that are optional but will improve relevance if matched.
  - MustNot: A list of conditions where the match must be excluded.
  - RangeFilter: An optional filter that applies date range conditions to the query.

Each condition is represented as a list of values of type `any` that must be fill with SearchQueryString, SearchMatch or SearchMatchAll.
*/
type BoolQuery struct {
	Must        []any       `json:"must"`
	Should      []any       `json:"should"`
	Not         []any       `json:"must_not"`
	RangeFilter []DateRange `json:"filter,omitempty"`
}

type SearchQueryString struct {
	QueryString QueryString `json:"query_string"`
}

// Query the analized
type QueryString struct {
	Query string `json:"query"`
}

/*
SearchMatch represents a match filter.
The map must contain the field names as keys and the corresponding query values as the values.
*/
type SearchMatch struct {
	Match map[string]SearchMatchField `json:"match"`
}

// Represent the data of the field in a Search Match
//
// Operator: OR to search any words in the query of AND to search each words in the query
//
// querystring: the value to searchs
type SearchMatchField struct {
	Operator    IOperator `json:"operator"`
	QueryString string    `json:"query"`
}

/*
Represents a query to find all values
*/
type SearchMatchAll struct {
	All struct{} `json:"match_all"`
}

/*
Represents a query to filter other queries in a interval of time
*/
type DateRange struct {
	Range struct {
		Date DateInterval `json:"date"`
	} `json:"range"`
}

type DateInterval struct {
	GreatherThan         *time.Time `json:"gt,omitempty"`
	LessThan             *time.Time `json:"lt,omitempty"`
	GreatherThanOrEquals *time.Time `json:"gte,omitempty"`
	LessThanOrEquals     *time.Time `json:"lte,omitempty"`
	TimeZone             string     `json:"time_zone,omitempty"`
	Format               string     `json:"format,omitempty"`
}

type HighlightQuery struct {
	PreTag  []string            `json:"pre_tags"`
	PostTag []string            `json:"post_tags"`
	Fields  map[string]struct{} `json:"fields"`
}
