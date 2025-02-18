package indexer

import (
	"fmt"
	"sync"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/config"
	db "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/database"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/models"
)

type BatchManager struct {
	id         int
	post       db.IPostMail
	mails      []models.Email
	LimitBatch int
	postSended int
}

func NewBatchManager(LimitBatch int) *BatchManager {
	return &BatchManager{
		post:       db.ZincDatabase(),
		mails:      make([]models.Email, 0),
		LimitBatch: LimitBatch,
	}
}

func NewBatchManagerWithID(LimitBatch, id int) *BatchManager {
	return &BatchManager{
		id:         id,
		post:       db.ZincDatabase(),
		mails:      make([]models.Email, 0),
		LimitBatch: LimitBatch,
	}
}

func (bm *BatchManager) SubmitAndSend(indexName string, mail *models.Email) error {
	bm.mails = append(bm.mails, *mail)
	var err error = nil
	if len(bm.mails) >= bm.LimitBatch {
		bm.postSended++
		fmt.Println("Batch Manager ", bm.id, " sended batch: ", bm.postSended)
		// err = bm.post.PostMails(indexName, bm.mails)
		bm.mails = nil
	}

	return err
}

func PoolMailBatch(mailsQueoeCh chan *models.Email, wg *sync.WaitGroup) {
	defer wg.Done()
	batchManager := NewBatchManagerWithID(1000, 1)

	for mail := range mailsQueoeCh {
		batchManager.SubmitAndSend(config.CFG.DatabaseName, mail)
	}
}
