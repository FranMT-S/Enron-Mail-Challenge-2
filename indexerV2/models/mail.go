package models

type Email struct {
	MessageID               string `json:"message_id"`
	Date                    string `json:"date"`
	DateTime                string `json:"datetime"`
	From                    string `json:"from"`
	FromSort                string `json:"from_sort"`
	To                      string `json:"to"`
	ToSort                  string `json:"to_sort"`
	Bcc                     string `json:"bcc"`
	Cc                      string `json:"cc"`
	Subject                 string `json:"subject"`
	SubjectSort             string `json:"subject_sort"`
	MimeVersion             string `json:"mime_version"`
	ContentType             string `json:"content_type"`
	ContentTransferEncoding string `json:"content_transfer_encoding"`
	XFrom                   string `json:"x_from"`
	XTo                     string `json:"x_to"`
	XCc                     string `json:"x_cc"`
	XBcc                    string `json:"x_bcc"`
	XFolder                 string `json:"x_folder"`
	XOrigin                 string `json:"x_origin"`
	XFileName               string `json:"x_file_name"`
	Body                    string `json:"body"`
}
