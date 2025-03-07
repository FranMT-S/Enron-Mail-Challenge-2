package helpers

import (
	"strings"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
)

func CreateSortFields(fields []string) []string {
	sortFields := []string{}

	for _, f := range fields {
		parts := strings.Split(f, ".")

		if len(parts) != 2 {
			continue
		}

		field := strings.ToLower(parts[0])
		sortType := strings.ToLower(parts[1])

		switch field {
		case models.FromField, models.ToField, models.SubjectField:
			if sortType == "asc" {
				sortFields = append(sortFields, field+"_sort")
			} else if sortType == "desc" {
				sortFields = append(sortFields, "-"+field+"_sort")
			}
		case models.XToField,
			models.BodyField,
			models.XFromField,
			models.DateField,
			models.DateTimeField:
			if sortType == "asc" {
				sortFields = append(sortFields, field)
			} else if sortType == "desc" {
				sortFields = append(sortFields, "-"+field)
			}
		default:
			continue
		}
	}

	return sortFields
}
