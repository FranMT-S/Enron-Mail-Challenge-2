package db

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/config"
	apierrors "github.com/FranMT-S/Enron-Mail-Challenge-2/backend/errors"
	"github.com/FranMT-S/Enron-Mail-Challenge-2/backend/models"
)

type zinchsearchDB struct {
	client *http.Client
}

func NewZinchSearchDB() *zinchsearchDB {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		MaxIdleConns:          20, // only one host isn't necessary many connections
		MaxConnsPerHost:       20, // only one host isn't necessary many connections
		MaxIdleConnsPerHost:   20,
		IdleConnTimeout:       30 * time.Second,
	}

	return &zinchsearchDB{
		client: &http.Client{
			Transport: transport,
			Timeout:   30 * time.Second,
		},
	}
}

func (zin *zinchsearchDB) SearchMails(query models.Query) (*models.Hits[models.Email], *apierrors.ResponseError) {
	var err error
	var bytesJson []byte
	url := config.ZINC_SEARCH_URL
	bytesJson, err = json.Marshal(query)

	if err != nil {
		errRes := apierrors.ErrResponseFailedProcessingQuery.WithLogError(err)
		return nil, errRes
	}

	reader := bytes.NewReader(bytesJson)

	res, errRes := zin.doRequest("POST", url, reader)
	if errRes != nil {
		return nil, errRes
	}

	var HitsResponse *models.Hits[models.Email]
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&HitsResponse)
	if err != nil {
		errRes = apierrors.ErrResponseFailedProcessingQuery.WithLogError(err)
		return nil, errRes
	}

	return HitsResponse, nil
}

func (zin *zinchsearchDB) SearchMail(id string) (*models.Hit[models.Email], *apierrors.ResponseError) {
	url := config.ZINC_GET_DOCUMENT_URL + "/" + id
	res, errRes := zin.doRequest("GET", url, nil)
	if errRes != nil {
		return nil, errRes
	}

	var HitResponse *models.Hit[models.Email]
	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(&HitResponse)
	if err != nil {
		errRes = apierrors.ErrResponseFailedProcessingQuery.WithLogError(err)
		return nil, errRes
	}

	return HitResponse, nil
}

func (zinDB zinchsearchDB) doRequest(method string, url string, body io.Reader) (*http.Response, *apierrors.ResponseError) {

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		errRes := apierrors.ErrResponseRequestNotProcessed.WithLogError(err)
		return nil, errRes
	}

	ZincHeader(req)

	res, err := zinDB.client.Do(req)
	if err != nil {
		var errRes *apierrors.ResponseError
		log.Println(err.Error())
		if errors.Is(err, context.DeadlineExceeded) {
			errRes = apierrors.ErrResponseRequestTimeOut
		} else if errors.Is(err, net.ErrClosed) {
			errRes = apierrors.ErrResponseDataBaseNotConnection
		} else {
			errRes = apierrors.ErrResponseRequestNotProcessed.WithLogError(err)
		}

		return nil, errRes
	}

	if res.StatusCode != http.StatusOK {

		errRes := apierrors.ErrResponseRequestNotProcessed.WithLogError(apierrors.ErrDoRequestFailed).SetStatus(res.StatusCode)
		return nil, errRes
	}

	return res, nil
}

func getAllMailsQuery(page, size int, fields []string) (string, *apierrors.ResponseError) {

	jsonByte, err := json.Marshal(fields)

	if err != nil {
		errRes := apierrors.ErrResponseFailedProcessingQuery.WithLogError(err)
		return "", errRes
	}

	from := (page - 1) * size
	query := fmt.Sprintf(`
	{
		"query": {
				"match_all": {}
		},
		"sort": [
				"-date"
		],
		"from": %v,
		"size": %v,
		"_source": %v
	}
`, from, size, string(jsonByte))

	return query, nil
}
