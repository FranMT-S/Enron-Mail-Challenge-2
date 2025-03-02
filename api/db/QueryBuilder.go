package db

import (
	"regexp"
	"strings"
	"time"

	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/helpers"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
)

func QueryBuilder(queryString string, page, size int, fields []string) (*models.Query, *apierrors.ResponseError) {
	QueryModel := models.NewQuery()
	from := (page - 1) * size
	QueryModel.SetFields(fields).SetSize(size).SetFrom(from)
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
	// replace possible the pattern field without value field:
	query = strings.ReplaceAll(query, ":", "")

	QueryModel.AddQueryString(query)
	return QueryModel, nil
}

func CleanCharacters(s string) string {
	s = strings.TrimSpace(s)
	// s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "*", "")
	return s
}

func GetFieldsExpresionList(queryString string) (fields []string) {
	searchFieldsExpresionRegex := regexp.MustCompile(helpers.FieldRegex)
	matchFieldsExpresionList := searchFieldsExpresionRegex.FindAllString(queryString, -1)

	return matchFieldsExpresionList
}

func CleanFieldsExpresion(queryString string) string {
	searchFieldsExpresionRegex := regexp.MustCompile(helpers.FieldRegex)
	return searchFieldsExpresionRegex.ReplaceAllString(queryString, "")

}

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

	rangeFilter.Range.Date.Format = time.RFC3339
	Query.AddRangeFilter(rangeFilter)

	return nil
}

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
