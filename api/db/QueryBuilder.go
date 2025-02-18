package db

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
)

func QueryBuilder(queryString string, page, size int, fields []string, timeZone string) (*models.Query, *apierrors.ResponseError) {
	QueryModel := models.NewQuery()
	from := (page - 1) * size
	QueryModel.SetFields(fields).SetSize(size).SetFrom(from)
	if queryString == "" {
		QueryModel.SetGetAllQuery()
		return QueryModel, nil
	}

	fieldInnerParenthesesRegex := `\([^)]+\)`
	fieldRegex := fmt.Sprintf(`\w+:(%v|\S+)`, fieldInnerParenthesesRegex)
	searchFieldsExpresionRegex := regexp.MustCompile(fieldRegex)

	fieldsExpresionList := searchFieldsExpresionRegex.FindAllString(queryString, -1)
	if errRes := processAndAddFilteringQueries(QueryModel, fieldsExpresionList, timeZone); errRes != nil {
		return nil, errRes
	}

	query := searchFieldsExpresionRegex.ReplaceAllString(queryString, "")
	query = cleanValue(query)

	QueryModel.AddQueryString(query)
	return QueryModel, nil
}

func cleanValue(s string) string {
	s = strings.TrimSpace(s)
	// s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	return s
}

func processAndAddFilteringQueries(Query *models.Query, fieldsExpresion []string, timeZone string) *apierrors.ResponseError {
	rangeFilter := models.DateRange{}
	layout := "2006-01-02"

	for _, field := range fieldsExpresion {
		parts := strings.SplitN(field, ":", 2)
		if len(parts) < 2 {
			continue
		}

		key := strings.ToLower(parts[0])
		value := cleanValue(parts[1])

		switch key {
		case "before":
			date, err := dateParsed(layout, value, timeZone)
			if err != nil {
				return err
			}
			rangeFilter.Range.Date.GreatherThanOrEquals = &date
		case "after":
			date, err := dateParsed(layout, value, timeZone)
			if err != nil {
				return err
			}

			rangeFilter.Range.Date.LessThanOrEquals = &date
		default:
			Query.AddMatchFieldFilter(key, value)
		}
	}

	if !isEmptyDateRange(rangeFilter) {
		rangeFilter.Range.Date.Format = time.RFC3339
		Query.AddRangeFilter(rangeFilter)
	}

	return nil
}

func isEmptyDateRange(dr models.DateRange) bool {
	return dr == models.DateRange{}
}

func dateParsed(layout, value, timeZone string) (time.Time, *apierrors.ResponseError) {
	var loc *time.Location
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		loc = time.UTC
	}

	date, err := time.ParseInLocation(layout, value, loc)
	if err != nil {
		resErr := apierrors.ErrResponseDateFormatNotValid.WithLogError(err)
		resErr.Message = resErr.Message + ".Value:" + value
		return time.Time{}, resErr
	}

	return date, nil

}
