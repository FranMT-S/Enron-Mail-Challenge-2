package indexer

import (
	"os"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/src/models"
)

type IMailIndexer interface {
	IndexMails(FilesPathQueoeCh chan string, MailsResultQueoe chan *models.Email)
	IndexMail(FilePath string) *models.Email
}

type Indexer struct {
	MailParser IParserMail
}

func (indexer *Indexer) IndexMails(FilesPathQueoeCh chan string, MailsResultQueoe chan *models.Email) {
	for filePath := range FilesPathQueoeCh {
		file, err := os.Open(filePath)

		if err != nil {
			continue
		}

		mail, err := indexer.MailParser.ConvertFileToMail(file)
		if err != nil {
			continue
		}

		MailsResultQueoe <- mail

		file.Close()
	}

	close(MailsResultQueoe)
}

func (indexer *Indexer) IndexMail(FilePath string) (*models.Email, error) {
	file, err := os.Open(FilePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	mail, _ := indexer.MailParser.ConvertFileToMail(file)

	return mail, nil
}
