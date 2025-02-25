package db

import (
	"fmt"
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

	valueInnerParentheses := `\([^)]+\)`
	NameFieldWithSpecialCharacters := `[-*]?\w+`

	fieldRegex := fmt.Sprintf(`%v:(%v|\S+)`, NameFieldWithSpecialCharacters, valueInnerParentheses)

	searchFieldsExpresionRegex := regexp.MustCompile(fieldRegex)
	fmt.Println("\\")
	matchFieldsExpresionList := searchFieldsExpresionRegex.FindAllString(queryString, -1)

	if errRes := processAndAddFilteringQueries(QueryModel, matchFieldsExpresionList); errRes != nil {
		return nil, errRes
	}

	query := searchFieldsExpresionRegex.ReplaceAllString(queryString, "")

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

func processAndAddFilteringQueries(Query *models.Query, fieldsExpresion []string) *apierrors.ResponseError {

	layouts := []string{"2006-01-02", "2006/01/02"}

	for _, field := range fieldsExpresion {
		isExclusionFilter := false
		isOperatorOR := false
		operator := models.AND
		parts := strings.SplitN(field, ":", 2)
		if len(parts) < 2 {
			continue
		}

		key := strings.ToLower(parts[0])
		optionsFlag := key[0:1]
		value := CleanCharacters(parts[1])

		isExclusionFilter = optionsFlag == "-"
		isOperatorOR = optionsFlag == "*"
		key = CleanCharacters(key)

		if isOperatorOR {
			operator = models.OR
		}

		switch key {
		case "since", "until", "before", "after", "date":
			date, err := helpers.DateParsed(value, "UTC", layouts)
			if err != nil {
				return err
			}

			rangeFilter := models.DateRange{}

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

			rangeFilter.Range.Date.Format = time.RFC3339
			Query.AddRangeFilter(rangeFilter)

		default:
			Query.AddMatchFieldFilter(key, value, operator, isExclusionFilter)
		}
	}

	return nil
}

func isEmptyDateRange(dr models.DateRange) bool {
	return dr == models.DateRange{}
}
