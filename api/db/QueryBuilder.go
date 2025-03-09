package db

import (
	"regexp"
	"strings"
	"time"

	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/helpers"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
)

/*
QueryBuilder constructs a query with pagination, field selection, and sorting.

parameters:

  - querystring the string with the query to analize
  - page to return
  - size number of element by page
  - sort list of the name of fields to order, if the name start with "-" is order descending
*/
func QueryBuilder(queryString string, page, size int, fields []string, sort []string) (*models.Query, *apierrors.ResponseError) {
	QueryModel := models.NewQuery()
	from := (page - 1) * size
	QueryModel.SetFields(fields).SetSize(size).SetFrom(from).SetSort(sort)
	if queryString == "" {
		QueryModel.SetGetAllQuery()
		return QueryModel, nil
	}

	matchFieldsExpresionList := GetFieldsExpresionList(queryString)
	query := CleanFieldsExpresion(queryString)

	if errRes := processAndAddFilteringQueries(QueryModel, matchFieldsExpresionList); errRes != nil {
		return nil, errRes
	}

	query = CleanCharacters(query)
	// Replace possible patterns where a field is present without a value, such as "to:"
	query = strings.ReplaceAll(query, ":", "")

	QueryModel.AddQueryString(query)
	return QueryModel, nil
}

// clean special characters
func CleanCharacters(s string) string {
	s = strings.TrimSpace(s)
	// s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "*", "")
	s = strings.ReplaceAll(s, ",", "")
	return s
}

// Return a list of match of the fields in the query
//
// example: To:sara, To:(sara Mary)
func GetFieldsExpresionList(queryString string) (fields []string) {
	searchFieldsExpresionRegex := regexp.MustCompile(helpers.FieldRegex)
	matchFieldsExpresionList := searchFieldsExpresionRegex.FindAllString(queryString, -1)

	return matchFieldsExpresionList
}

// Remove all match of the fields in the query
func CleanFieldsExpresion(queryString string) string {
	searchFieldsExpresionRegex := regexp.MustCompile(helpers.FieldRegex)
	return searchFieldsExpresionRegex.ReplaceAllString(queryString, "")
}

// Process all fields expresion list and create the type of filter the add the filters to the Query Model
func processAndAddFilteringQueries(Query *models.Query, fieldsExpresion []string) *apierrors.ResponseError {
	layouts := []string{"2006-01-02", "2006/01/02"}
	rangeFilter := models.DateRange{}

	for _, field := range fieldsExpresion {
		parts := strings.SplitN(field, ":", 2)
		if len(parts) < 2 {
			continue
		}

		key, value, isExclusionFilter, operator := GetFilterInformation(parts)

		switch key {
		case "since", "until", "before", "after", "date":
			date, err := helpers.DateParsed(value, "UTC", layouts)
			if err != nil {
				return err
			}

			switch key {
			case "since":
				rangeFilter.Range.Date.GreatherThanOrEquals = &date
			case "until":
				rangeFilter.Range.Date.LessThanOrEquals = &date
			case "after":
				rangeFilter.Range.Date.GreatherThan = &date
			case "before":
				rangeFilter.Range.Date.LessThan = &date
			case "date":
				rangeFilter.Range.Date.GreatherThanOrEquals = &date
				rangeFilter.Range.Date.LessThanOrEquals = &date
			}

		default:
			Query.AddMatchFieldFilter(key, value, operator, isExclusionFilter)
		}
	}

	if !IsEmptyDateRange(rangeFilter) {
		rangeFilter.Range.Date.Format = time.RFC3339
		Query.AddRangeFilter(rangeFilter)
	}

	return nil
}

// Get the information of the type of filter
// parts is a array with two elements, the key or name of field and the value of the field
func GetFilterInformation(parts []string) (key, value string, isExclusionFilter bool, operator models.IOperator) {
	operator = models.AND
	key = strings.ToLower(parts[0])
	optionsFlag := key[0:1]
	value = CleanCharacters(parts[1])

	isExclusionFilter = optionsFlag == "-"
	isOperatorOR := optionsFlag == "*"
	key = CleanCharacters(key)

	if isOperatorOR {
		operator = models.OR
	}

	return key, value, isExclusionFilter, operator
}

func IsEmptyDateRange(dr models.DateRange) bool {
	return dr == models.DateRange{}
}
