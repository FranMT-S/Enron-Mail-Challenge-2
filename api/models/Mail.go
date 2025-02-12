package models

const (
	MessageIDField               string = "message_id"
	DateField                    string = "date"
	FromField                    string = "from"
	ToField                      string = "to"
	BccField                     string = "bcc"
	CcField                      string = "cc"
	SubjectField                 string = "subject"
	MimeField                    string = "mime_version"
	ContentTypeField             string = "content_type"
	ContentTransferEncodingField string = "content_transfer_encoding"
	XFromField                   string = "x_from"
	XToField                     string = "x_to"
	XCcField                     string = "x_cc"
	XBccField                    string = "x_bcc"
	XFolderField                 string = "x_folder"
	XOriginField                 string = "x_origin"
	XFileName                    string = "x_file_name"
	BodyField                    string = "body"
)

type Email struct {
	MessageID               string `json:"message_id"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	To                      string `json:"to"`
	Bcc                     string `json:"bcc"`
	Cc                      string `json:"cc"`
	Subject                 string `json:"subject"`
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

type EmailSummary struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
}
