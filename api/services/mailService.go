package services

import (
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/db"
	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
)

// const (
// 	_message_id                = "message_id"
// 	_date                      = "date"
// 	_from                      = "from"
// 	_to                        = "to"
// 	_bcc                       = "bcc"
// 	_cc                        = "cc"
// 	_subject                   = "subject"
// 	_mime_version              = "mime_version"
// 	_content_type              = "content_type"
// 	_content_transfer_encoding = "content_transfer_encoding"
// 	_x_from                    = "x_from"
// 	_x_to                      = "x_to"
// 	_x_cc                      = "x_cc"
// 	_x_bcc                     = "x_bcc"
// 	_x_folder                  = "x_folder"
// 	_x_origin                  = "x_origin"
// 	_x_file_name               = "x_file_name"
// 	_body                      = "body"
// )

type IMailService interface {
	GetMailsHitsAndTotal(query string, page int, size int, timeZone string) (*models.EmailSummaryResponse, *apierrors.ResponseError)
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

func (ms MailService) GetMailsHitsAndTotal(queryString string, page int, size int, timeZone string) (*models.EmailSummaryResponse, *apierrors.ResponseError) {
	fields := []string{models.FromField, models.ToField, models.DateField, models.SubjectField}
	Query, err := db.QueryBuilder(queryString, page, size, fields, timeZone)
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

	return &email, nil
}
