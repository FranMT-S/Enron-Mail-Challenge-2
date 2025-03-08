package indexer

import (
	"fmt"

	db "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/database"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/models"
)

type MailsUploader struct {
	ID          int
	postHandler db.IPostMail
	indexName   string
}

func NewMailsUploader(indexName string, PostMailHandler db.IPostMail) *MailsUploader {
	return &MailsUploader{
		postHandler: PostMailHandler,
		indexName:   indexName,
	}
}

func (um *MailsUploader) Upload(mails []*models.Email) (int, error) {
	var err error = nil
	err = um.postHandler.PostMails(um.indexName, mails)
	if err != nil {
		return 0, err
	}
	total := len(mails)
	fmt.Println("upload mails:", total)
	return total, nil
}
