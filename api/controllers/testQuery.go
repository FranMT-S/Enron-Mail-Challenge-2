package controllers

import (
	"net/http"
	"strings"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/db"
	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/helpers"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/middlewares"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
	"github.com/go-chi/render"
)

func TestQueryBuilderfunc(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, err := middlewares.GetPageFromContext(r)
	sortParam := r.URL.Query().Get("sort")
	sort := strings.Split(sortParam, ",")
	if err != nil {
		apierrors.RenderJSON(w, err)
		return
	}

	size, err := middlewares.GetSizeFromContext(r)
	if err != nil {
		apierrors.RenderJSON(w, err)
		return
	}
	fields := []string{models.FromField, models.ToField, models.DateTimeField, models.DateField, models.SubjectField}

	sortFields := helpers.CreateSortFields(sort)
	if len(sortFields) == 0 || sortParam == "" {
		sortFields = []string{"-date"}
	}

	Query, errRes := db.QueryBuilder(query, page, size, fields, sortFields)
	if errRes != nil {
		apierrors.RenderJSON(w, errRes)
		return
	}
	render.JSON(w, r, Query)
}
