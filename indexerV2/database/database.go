package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/config"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/indexer/models"
)

var z_database *zincDatabase

type zincDatabase struct {
	client *http.Client
}

// ZincDatabase returns a single instance of the database
func ZincDatabase() *zincDatabase {
	if z_database == nil {
		transport := &http.Transport{
			TLSHandshakeTimeout:   30 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
			MaxIdleConns:          20, // only one host isn't necessary many connections
			MaxConnsPerHost:       20, // only one host isn't necessary many connections
			MaxIdleConnsPerHost:   10,
			IdleConnTimeout:       60 * time.Second,
		}

		z_database = &zincDatabase{client: &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}}
	}

	return z_database
}

// Create index in zincsearch if not exist
func (db zincDatabase) CreateIndex(indexName string) error {
	index := GetIndexString(indexName)
	data := strings.NewReader(index)
	req, err := http.NewRequest(http.MethodPost, config.ZINC_CREATE_INDEX_URL, data)
	if err != nil {
		return err
	}

	ZincHeader(req)

	resp, err := db.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		return err
	}

	if string(body) != fmt.Sprintf(`{"error":"index [%v] already exists"}`, indexName) {
		fmt.Println("index created")
	}

	return nil
}

// Postmail send many mails to database.
func (db zincDatabase) PostMails(indexName string, emails []*models.Email) error {
	reader, err := bulkMailReader(indexName, emails)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, config.ZINC_BULK_API_URL, reader)
	if err != nil {
		return err
	}

	ZincHeader(req)

	resp, err := db.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Create a reader to upload mails in zincsearch
func bulkMailReader(indexname string, mails []*models.Email) (io.Reader, error) {
	bulkResponse := models.ZincBulkResponse{
		Index:   indexname,
		Records: mails,
	}

	jsonByte, err := json.Marshal(bulkResponse)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonByte)
	return reader, nil
}
