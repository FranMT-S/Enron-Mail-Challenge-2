package indexer

import (
	"sync"

	db "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/database"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/models"
)

type BatchMails struct {
	id         int
	post       db.IPostMail
	mails      []*models.Email
	LimitBatch int
	postSended int
}

func NewBatchMails(LimitBatch int) *BatchMails {
	if LimitBatch <= 0 {
		LimitBatch = 1000
	}

	return &BatchMails{
		post:       db.ZincDatabase(),
		mails:      make([]*models.Email, 0),
		LimitBatch: LimitBatch,
	}
}

func NewBatchMailsWithID(LimitBatch, id int) *BatchMails {
	if LimitBatch <= 0 {
		LimitBatch = 1000
	}

	return &BatchMails{
		id:         id,
		post:       db.ZincDatabase(),
		mails:      make([]*models.Email, 0),
		LimitBatch: LimitBatch,
	}
}

func (bm *BatchMails) Add(mail *models.Email) {
	if mail != nil {
		bm.mails = append(bm.mails, mail)
	}
}

func (bm *BatchMails) IsFill() bool {
	return len(bm.mails) >= bm.LimitBatch
}

func (bm *BatchMails) HasElements() bool {
	return len(bm.mails) > 0
}

func (bm *BatchMails) SendMails(receiver chan []*models.Email) {
	bm.postSended++
	receiver <- bm.mails
}

func (bm *BatchMails) CleanBatch() {
	bm.mails = nil
}

func BatchQueoue(bathSize int, mailsQueoeCh chan *models.Email, receiver chan []*models.Email, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	batch := NewBatchMails(1000)

	for mail := range mailsQueoeCh {
		batch.Add(mail)
		if batch.IsFill() {
			batch.SendMails(receiver)
			batch.CleanBatch()
		}
	}

	if batch.HasElements() {
		batch.SendMails(receiver)
		batch.CleanBatch()
	}

	close(receiver)
}
