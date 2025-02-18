package indexer

import (
	"fmt"
	"log"
	"sync"

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

func (um *MailsUploader) Upload(mails []*models.Email) error {
	var err error = nil
	if len(mails) > 0 {
		err = um.postHandler.PostMails(um.indexName, mails)
	}

	return err
}

func UploaderPool(
	poolSize int,
	indexName string,
	mailHandler db.IPostMail,
	mailsBatchQueoeCh chan []*models.Email,
	errhQueoeCh chan error,
	wg *sync.WaitGroup,
) {
	wg.Add(poolSize)
	for i := 0; i < poolSize; i++ {
		go func(id int) {
			defer wg.Done()
			uploader := NewMailsUploader(indexName, mailHandler)
			uploader.ID = id
			count := 0
			for batch := range mailsBatchQueoeCh {
				err := uploader.Upload(batch)
				if err != nil && errhQueoeCh != nil {
					errhQueoeCh <- err
					log.Println("Failed upload batch mails")
					continue
				}
				count++
				fmt.Printf("batch %v upload a batch:%v\n", uploader.ID, count)
			}

		}(i)
	}
}
