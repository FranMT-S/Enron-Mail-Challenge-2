package services

import (
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/db"
	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
)

type IMailService interface {
	GetMailsHitsAndTotal(query string, page int, size int) (*models.EmailSummaryResponse, *apierrors.ResponseError)
	GetMailByID(id string) (*models.Email, *apierrors.ResponseError)
}

type MailService struct {
	indexDB db.IndexDB
}

func NewMailService(indexDb db.IndexDB) *MailService {
	return &MailService{
		indexDB: indexDb,
	}
}

func (ms MailService) GetMailsHitsAndTotal(queryString string, page int, size int) (*models.EmailSummaryResponse, *apierrors.ResponseError) {
	fields := []string{
		models.FromField, models.ToField, models.DateField,
		models.SubjectField, models.XToField, models.XFromField,
	}

	Query, err := db.QueryBuilder(queryString, page, size, fields)
	if err != nil {
		return nil, err
	}

	hits, err := ms.indexDB.SearchMails(*Query)
	if err != nil {
		return nil, err
	}

	emails := []models.EmailSummary{}

	total := hits.Hits.Total.Value
	for _, hit := range hits.Hits.Hits {
		email := models.EmailSummary{}
		email.ID = hit.ID
		email.From = hit.Source.From
		email.To = hit.Source.To
		email.Date = hit.Source.Date
		email.XFrom = hit.Source.XFrom
		email.XTo = hit.Source.XTo
		email.Subject = hit.Source.Subject
		emails = append(emails, email)
	}

	emailsResponse := &models.EmailSummaryResponse{
		Total:  total,
		Emails: emails,
	}

	return emailsResponse, nil
}

func (ms MailService) GetMailByID(id string) (*models.Email, *apierrors.ResponseError) {
	hit, err := ms.indexDB.SearchMail(id)
	if err != nil {
		return nil, err
	}

	email := hit.Source
	email.ID = hit.ID

	return &email, nil
}
