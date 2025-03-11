package db

import (
	"context"

	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
)

type IndexDB interface {
	SearchMails(ctx context.Context, query models.Query) (*models.Hits[models.Email], *apierrors.ResponseError)
	SearchMail(ctx context.Context, id string) (*models.Hit[models.Email], *apierrors.ResponseError)
}
