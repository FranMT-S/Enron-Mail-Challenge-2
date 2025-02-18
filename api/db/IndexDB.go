package db

import (
	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
)

type IndexDB interface {
	SearchMails(query models.Query) (*models.Hits[models.Email], *apierrors.ResponseError)
	SearchMail(id string) (*models.Hit[models.Email], *apierrors.ResponseError)
}
