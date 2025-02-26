package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/constants"
	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
)

type paginatorFunc func(maxSize int) func(next http.Handler) http.Handler

// add page and size to http context
var Paginator paginatorFunc = func(maxSize int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			pageParam := query.Get(constants.PAGE)
			if pageParam == "" {
				pageParam = r.Form.Get(constants.PAGE)
			}

			page, err := strconv.Atoi(pageParam)
			if err != nil || page < 1 {
				page = 1 // Página por defecto
			}

			sizeParam := query.Get(constants.SIZE)
			if sizeParam == "" {
				sizeParam = r.Form.Get(constants.SIZE)
			}

			size, err := strconv.Atoi(sizeParam)
			if err != nil || size < 1 {
				size = 100
			}

			if size > maxSize {
				size = maxSize
			}

			ctx := context.WithValue(r.Context(), constants.PAGE, page)
			ctx = context.WithValue(ctx, constants.SIZE, size)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (p paginatorFunc) GetPageFromContext(r *http.Request) (int, *apierrors.ResponseError) {
	page, ok := r.Context().Value(constants.PAGE).(int)
	if !ok {
		return -1, apierrors.ErrResponsePaginatorNotIsSetup
	}
	return page, nil
}

func (p paginatorFunc) GetSizeFromContext(r *http.Request) (int, *apierrors.ResponseError) {
	size, ok := r.Context().Value(constants.SIZE).(int)
	if !ok {
		return -1, apierrors.ErrResponsePaginatorNotIsSetup
	}
	return size, nil
}
