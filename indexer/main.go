package main

import (
	"fmt"
	"sync"
	"time"

	db "github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/src/database"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/src/indexer"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/src/models"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/src/shared"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/src/workers"
)

func main() {
	shared.InitializeVarEnviroment()
	zincDb := db.ZincDatabase()
	zincDb.CreateIndex()
	filesQueoCh := make(chan string)
	MailResultQueoCh := make(chan *models.Email)
	errorCh := make(chan error)
	var wg sync.WaitGroup
	var wgWorkerPool sync.WaitGroup

	// register time
	startTime := time.Now()
	// filesQueoCh := shared.FindFiles("./mails/mails-1")
	// mailsQueoeCh := make(chan *models.Email)

	indexer := indexer.Indexer{
		MailParser: &indexer.MimeParse{},
	}

	wg.Add(1)
	go shared.FindFilesAsync("./mails/mails-3", filesQueoCh, &wg)

	go func() {
		wg.Wait()
		close(filesQueoCh)
	}()

	workerpool := workers.NewWorkerPool(
		indexer.IndexMail,
		filesQueoCh,
		workers.WorkerPoolConfig[*models.Email]{
			Workers:     10000,
			ResultQueue: MailResultQueoCh,
			ErrorCh:     errorCh,
		},
		&wgWorkerPool,
	)

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		i := 0
		var mails []models.Email

		for mail := range MailResultQueoCh {
			i++
			fmt.Println("body:", mail.Body)
			mails = append(mails, *mail)
		}
		fmt.Println("index")
		zincDb.PostMails(mails)
		wg2.Done()
	}()

	workerpool.Start()
	wg2.Wait()
	// for file := range filesQueoCh {
	// 	fmt.Println(file)

	// 	// if err != nil {
	// 	// 	continue
	// 	// }

	// 	// // Convertir a JSON
	// 	// jsonData, err := json.Marshal(mail)
	// 	// if err != nil {
	// 	// 	fmt.Println("Error al convertir a JSON:", err)
	// 	// 	return
	// 	// }

	// 	// // Convertir a una cadena legible (opcional, sin compactar)
	// 	// jsonPretty, _ := json.MarshalIndent(mail, "", "  ")

	// 	// // Imprimir el JSON como string
	// 	// fmt.Println("JSON Compacto:", string(jsonData))
	// 	// fmt.Println("JSON Legible:")
	// 	// fmt.Println(string(jsonPretty))

	// }

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	seconds := duration.Seconds()
	minutes := duration.Minutes()
	fmt.Printf("The code ran in %.2f seconds\n", seconds)
	fmt.Printf("The code ran in %.2f minutes\n", minutes)
	fmt.Printf("The code ran in %.2f minutes\n", minutes)
	// num := []int{1, 4, 6, 8, 9}
	// fmt.Println(num[len(num)-1])
}
