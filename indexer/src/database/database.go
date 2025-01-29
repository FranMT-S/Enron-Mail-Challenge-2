package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/src/models"
)

var z_database *zincDatabase

type zincDatabase struct {
	client *http.Client
}

// ZincDatabase returns a single instance of the database
func ZincDatabase() *zincDatabase {
	if z_database == nil {
		z_database = &zincDatabase{client: &http.Client{
			Timeout: time.Second * 30,
		}}
	}

	return z_database
}

func (db zincDatabase) CreateIndex() error {
	indexName := os.Getenv("INDEX")

	url := os.Getenv("URL") + "/index"

	data := strings.NewReader(GetIndexString(indexName))
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		// _logs.Println(constants_log.ERROR_CREATE_BASE + ":" + err.Error())
		// _logs.LogSVG(
		// 	constants_log.FILE_NAME_ERROR_DATABASE,
		// 	constants_log.OPERATION_DATABASE,
		// 	constants_log.ERROR_CREATE_BASE,
		// 	err,
		// )

		return err
	}

	ZincHeader(req)

	resp, err := db.client.Do(req)
	if err != nil {

		// _logs.Println(constants_log.ERROR_DATA_BASE + ":" + err.Error())
		// _logs.LogSVG(
		// 	constants_log.FILE_NAME_ERROR_DATABASE,
		// 	constants_log.OPERATION_DATABASE,
		// 	constants_log.ERROR_DATA_BASE,
		// 	err,
		// )

		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// _logs.Println(constants_log.ERROR_DATA_BASE + ":" + err.Error())
		// _logs.LogSVG(
		// 	constants_log.FILE_NAME_ERROR_DATABASE,
		// 	constants_log.OPERATION_DATABASE,
		// 	constants_log.ERROR_DATA_BASE,
		// 	err,
		// )

		return err
	}

	if string(body) != fmt.Sprintf(`{"error":"index [%v] already exists"}`, indexName) {
		// _logs.ColorGreen()
		fmt.Println("index created")
		fmt.Println(string(body))
		// _logs.ColorWhite()
	}

	return nil
}

// BulkRequest makes a request to the database to store the files.
func (db zincDatabase) PostMails(emails []models.Email) error {

	url := os.Getenv("URL") + "/_bulkv2"

	fmt.Println(url)
	reader, err := getMailsReader(emails)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return err
	}

	ZincHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("\n" + string(body))
	return nil
}

func getMailsReader(mails []models.Email) (io.Reader, error) {
	bulkResponse := models.ZincBulkResponse{
		Index:   os.Getenv("INDEX"),
		Records: mails,
	}
	println(bulkResponse.Index)
	jsonByte, err := json.Marshal(bulkResponse)
	if err != nil {
		return nil, err
	}
	println(string(jsonByte))

	reader := bytes.NewReader(jsonByte)
	return reader, nil
}

// BulkRequest makes a request to the database to store the files.
func BulkRequest(command, mailsData string) error {

	url := os.Getenv("URL") + command

	data := strings.NewReader(mailsData)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return err
	}

	ZincHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("\n" + string(body))
	return nil
}
