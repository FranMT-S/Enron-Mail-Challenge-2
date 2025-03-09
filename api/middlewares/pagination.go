package middlewares

import (
	"context"
	"net/http"
	"strconv"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/constants"
	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
)

type paginatorFunc func(maxSize int) func(next http.Handler) http.Handler

// Add page and size to http context,
var Paginator paginatorFunc = func(maxSize int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			pageParam := query.Get(constants.PARAM_PAGE)
			if pageParam == "" {
				pageParam = r.Form.Get(constants.PARAM_PAGE)
			}

			page, err := strconv.Atoi(pageParam)
			if err != nil || page < 1 {
				page = 1
			}

			sizeParam := query.Get(constants.PARAM_SIZE)
			if sizeParam == "" {
				sizeParam = r.Form.Get(constants.PARAM_SIZE)
			}

			size, err := strconv.Atoi(sizeParam)
			if err != nil || size < 1 {
				size = constants.PAGINATOR_MINSIZE
			}

			if size > maxSize {
				size = maxSize
			}

			ctx := context.WithValue(r.Context(), constants.PARAM_PAGE, page)
			ctx = context.WithValue(ctx, constants.PARAM_SIZE, size)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetPageFromContext(r *http.Request) (int, *apierrors.ResponseError) {
	page, ok := r.Context().Value(constants.PARAM_PAGE).(int)
	if !ok {
		return -1, apierrors.ErrResponsePaginatorNotIsSetup
	}
	return page, nil
}

func GetSizeFromContext(r *http.Request) (int, *apierrors.ResponseError) {
	size, ok := r.Context().Value(constants.PARAM_SIZE).(int)
	if !ok {
		return -1, apierrors.ErrResponsePaginatorNotIsSetup
	}
	return size, nil
}
