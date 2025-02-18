package models

type EmailSummaryResponse struct {
	Total  int            `json:"total"`
	Emails []EmailSummary `json:"emails"`
}
