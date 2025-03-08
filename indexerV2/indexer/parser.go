package indexer

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/models"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/shared"
)

type IParserMail interface {
	ConvertFileToMail(file *os.File) (*models.Email, error)
}

// Read a file and parse its content to a Email model
type MimeParse struct {
}

func (parser *MimeParse) ConvertFileToMail(file *os.File) (*models.Email, error) {
	var lastFieldUpdated string
	mail := &models.Email{}
	isHeader := true
	SizeKB := 4
	buffer := make([]byte, SizeKB*1024)

	r := bufio.NewReader(file)
	for {
		numBytes, err := r.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		chunk := string(buffer[:numBytes])

		if isHeader {
			err := processChunk(chunk, mail, &lastFieldUpdated, &isHeader)
			if err != nil {
				return nil, err
			}
		} else {
			mail.Body += chunk
		}
	}

	mail.Body = strings.TrimSpace(mail.Body)

	return mail, nil
}

func processChunk(chunk string, mail *models.Email, lastFieldUpdated *string, isHeader *bool) (err error) {

	lines := strings.Split(chunk, "\n")
	var field, value string
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			*isHeader = false
		}

		if !(*isHeader) {
			mail.Body += strings.Join(lines[i:], "\n")
			return nil
		}

		fields := strings.SplitN(line, ":", 2)

		if len(fields) > 1 {
			field, value = strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1])
			field = strings.ToLower(field)
			*lastFieldUpdated = field
		} else {
			field = strings.ToLower(*lastFieldUpdated)
			value = strings.TrimSpace(fields[0])
		}

		switch field {
		case "message-id":
			mail.MessageID += value
		case "date":
			datetime, err := shared.ParseDate(value)
			if err != nil {
				fmt.Println("Date Error")
				return err
			}
			dateWithoutOffset := datetime.UTC()

			mail.Date = dateWithoutOffset.Format(time.DateOnly)
			mail.DateTime = datetime.Format(time.RFC3339)
		case "from":
			mail.From += value
			mail.FromSort += value
		case "to":
			mail.To += value
			mail.ToSort += value
		case "bcc":
			mail.Bcc += value
		case "cc":
			mail.Cc += value
		case "subject":
			mail.Subject += value
			mail.SubjectSort += value
		case "mime-version":
			mail.MimeVersion += value
		case "content-type":
			mail.ContentType += value
		case "content-transfer-encoding":
			mail.ContentTransferEncoding += value
		case "x-from":
			mail.XFrom += value
		case "x-to":
			mail.XTo += value
		case "x-cc":
			mail.XCc += value
		case "x-bcc":
			mail.XBcc += value
		case "x-folder":
			mail.XFolder += value
		case "x-origin":
			mail.XOrigin += value
		case "x-filename":
			mail.XFileName += value
		}
	}

	return nil
}
