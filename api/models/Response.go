package models

type EmailHitsResponse[T any] struct {
	Total  int `json:"total"`
	Emails []Hit[T]
}

type EmailSummaryResponse struct {
	Total  int `json:"total"`
	Emails []EmailSummary
}
