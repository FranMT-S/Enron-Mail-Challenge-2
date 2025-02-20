package db

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
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

type ZincError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
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
		if errRes := isConnectionError(err); errRes != nil {
			return nil, errRes.WithLogError(err)
		}

		return nil, apierrors.ErrResponseRequestNotProcessed.WithLogError(err)
	}

	if res.StatusCode != http.StatusOK {
		bodyError, _ := io.ReadAll(res.Body)
		if errRes := IsResponseErrorFromZincSearch(bodyError); errRes != nil {
			return nil, errRes
		}

		errRes := apierrors.ErrResponseRequestNotProcessed.WithLogError(apierrors.ErrDoRequestFailed).SetStatus(res.StatusCode)
		return nil, errRes
	}

	return res, nil
}

func isConnectionError(err error) *apierrors.ResponseError {
	var netErr net.Error
	var urlErr *url.Error

	if errors.As(err, &netErr) {
		return apierrors.ErrResponseDataBaseNotConnection
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return apierrors.ErrResponseRequestTimeOut
	}

	if errors.As(err, &urlErr) {
		if errors.As(urlErr.Err, &netErr) {
			return apierrors.ErrResponseDataBaseNotConnection
		}
	}

	return nil
}

func IsResponseErrorFromZincSearch(body []byte) *apierrors.ResponseError {
	var zincErr apierrors.ZincSearchError
	var zincErrType apierrors.ZincSearchErrorWithType

	err := json.Unmarshal(body, &zincErr)
	if err == nil {
		if errRes := apierrors.IsZincError(zincErr); errRes != nil {
			return errRes
		}
	}

	err = json.Unmarshal(body, &zincErrType)
	if err == nil {
		if errRes := apierrors.IsZincErrorWithType(zincErrType); errRes != nil {
			return errRes
		}
	}

	return nil
}
