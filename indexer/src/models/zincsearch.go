package models

type ZincBulkResponse struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}
