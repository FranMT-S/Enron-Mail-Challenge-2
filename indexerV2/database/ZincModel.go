package db

import "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/models"

type TypeProperty string

const (
	Text    TypeProperty = "text"
	Keyword TypeProperty = "keyword"
	Numeric TypeProperty = "numeric"
	Date    TypeProperty = "date"
	Bool    TypeProperty = "bool"
)

type ZincSearchIndex struct {
	Name        string   `json:"name"`
	StorageType string   `json:"storage_type"`
	Mappings    Mappings `json:"mappings"`
}

type Mappings struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	MessageID               Property `json:"message_id"`
	Date                    Property `json:"date"`
	DateTime                Property `json:"datetime"`
	From                    Property `json:"from"`
	From_Sort               Property `json:"from_sort"`
	To                      Property `json:"to"`
	To_Sort                 Property `json:"to_sort"`
	Bcc                     Property `json:"bcc"`
	Cc                      Property `json:"cc"`
	Subject                 Property `json:"subject"`
	Subject_Sort            Property `json:"subject_sort"`
	MimeVersion             Property `json:"mime_version"`
	ContentType             Property `json:"content_type"`
	ContentTransferEncoding Property `json:"content_transfer_encoding"`
	XFrom                   Property `json:"x_from"`
	XTo                     Property `json:"x_to"`
	XCc                     Property `json:"x_cc"`
	XBcc                    Property `json:"x_bcc"`
	XFolder                 Property `json:"x_folder"`
	XOrigin                 Property `json:"x_origin"`
	XFileName               Property `json:"x_file_name"`
	Body                    Property `json:"body"`
}

type Property struct {
	Type          TypeProperty `json:"type"`
	Format        string       `json:"format,omitempty"`
	Index         bool         `json:"index"`
	Store         bool         `json:"store"`
	Sortable      bool         `json:"sortable"`
	Aggregratable bool         `json:"aggregatable"`
	Highlightable bool         `json:"highlightable"`
}

type IPostMail interface {
	PostMails(indexName string, emails []*models.Email) error
}
