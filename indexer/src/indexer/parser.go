package indexer

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/src/models"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/src/shared"
)

type IParserMail interface {
	ConvertFileToMail(file *os.File) (*models.Email, error)
}

type MimeParse struct {
}

var poolMails = &sync.Pool{
	New: func() interface{} {
		// Función que se ejecuta cuando no hay objetos disponibles en el pool
		return &models.Email{}
	},
}

func (parser *MimeParse) ConvertFileToMail(file *os.File) (*models.Email, error) {
	var partialLine string
	var lastField string
	mail := poolMails.Get().(*models.Email)
	isHeader := true
	KB := 4

	buffer := make([]byte, KB*1024)
	// fmt.Println("-----------Initial CHUNK-----------")
	for {
		numBytes, err := file.Read(buffer)
		if err == io.EOF {
			if len(partialLine) > 0 {
				fmt.Println(partialLine)
			}
			break
		}
		if err != nil {
			return nil, err
		}

		chunk := string(buffer[:numBytes])

		if isHeader {
			processChunk(chunk, mail, &lastField, &isHeader)
		} else {
			mail.Body += chunk
		}
	}

	mail.Body = strings.TrimSpace(mail.Body)
	return mail, nil
}

func processChunk(chunk string, mail *models.Email, lastField *string, isHeader *bool) (err error) {
	// fmt.Println("-----Chunk-----------")
	// fmt.Println(chunk)
	// fmt.Println("-----Chunk-----------")
	// fmt.Println()
	lines := strings.Split(chunk, "\n")
	var field, value string
	for i, line := range lines {

		// fmt.Println("Last Field:", *lastField)
		if strings.TrimSpace(line) == "" {
			*isHeader = false
		}

		if !(*isHeader) {
			mail.Body += strings.Join(lines[i:], "")
			return nil
		}

		fields := strings.SplitN(line, ":", 2)

		if len(fields) > 1 {
			field, value = strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1])
			field = strings.ToLower(field)
			*lastField = field
		} else {
			field = strings.ToLower(*lastField)
			value = strings.TrimSpace(fields[0])
		}
		// fmt.Println("Field:", field, ",Value:", value)
		// fmt.Println()
		switch field {
		case "message-id":
			mail.MessageID += value
		case "date":
			date, err := shared.ParseDate(value)
			if err != nil {
				fmt.Println("Date Error")
				return err
			}
			mail.Date = date
		case "from":
			mail.From += value
		case "to":
			mail.To += value
		case "bcc":
			mail.Bcc += value
		case "cc":
			mail.Cc += value
		case "subject":
			mail.Subject += value
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
