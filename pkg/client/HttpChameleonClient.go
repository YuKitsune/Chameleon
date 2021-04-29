package client

import (
	"bytes"
	json "encoding/json"
	"errors"
	"fmt"
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/resources"
	"github.com/yukitsune/chameleon/pkg/smtp"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type HttpChameleonClient struct {
	BaseURL    *url.URL
	httpClient http.Client
	logger     log.ChameleonLogger
}

func NewHttpChameleonClient(baseUrl *url.URL, logger log.ChameleonLogger) HttpChameleonClient {
	return HttpChameleonClient{
		BaseURL:    baseUrl,
		httpClient: http.Client{},
		logger:     logger,
	}
}

func (c *HttpChameleonClient) Validate(sender string, recipient string) error {

	req, err := getValidateRequest(c, sender, recipient)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	// If we got a 200 response, then we're good to go!
	if res.StatusCode == 200 {
		return nil
	}

	// Todo: Improve handling for other status codes
	body, err := ioutil.ReadAll(res.Body)
	c.logger.WithFields(log.Fields{
		"code": strconv.Itoa(res.StatusCode),
		"data": body,
	}).Errorf("Expected 200 when validating sender/recipient but found %d", res.StatusCode)
	return errors.New(fmt.Sprintf("Unexpected response code: %d", res.StatusCode))
}

func (c *HttpChameleonClient) Handle(e *smtp.Envelope) error {

	req, err := getHandleRequest(c, e)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	// If we got a 200 response, then we're good to go!
	if res.StatusCode == 200 {
		return nil
	}

	// Todo: Improve handling for other status codes
	body, err := ioutil.ReadAll(res.Body)
	c.logger.WithFields(log.Fields{
		"code": strconv.Itoa(res.StatusCode),
		"data": body,
	}).Errorf("Expected 200 when validating sender/recipient but found %d", res.StatusCode)
	return errors.New(fmt.Sprintf("Unexpected response code: %d", res.StatusCode))
}

func getValidateUrl(sender string, recipient string) *url.URL {
	query := fmt.Sprintf("sender=%s&recipient=%s", sender, recipient)
	url := &url.URL{
		Path:     "/validate",
		RawQuery: query,
	}
	return url
}

func getHandleUrl() *url.URL {
	return &url.URL{Path: "/handle"}
}

func getValidateRequest(client *HttpChameleonClient, sender string, recipient string) (*http.Request, error) {

	// Get the URL
	validateUrl := getValidateUrl(sender, recipient)
	fullUrl := client.BaseURL.ResolveReference(validateUrl)

	// Create the HTTP request
	req, err := http.NewRequest("GET", fullUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}

func getHandleRequest(client *HttpChameleonClient, e *smtp.Envelope) (*http.Request, error) {

	// Read the raw data from the envelope
	var rawData []byte
	reader := e.NewReader()
	_, err := reader.Read(rawData)
	if err != nil {
		return nil, err
	}

	// Create the resource
	resource := &resources.HandleRequestResource{
		Recipient: e.RcptTo[len(e.RcptTo)-1].String(),
		RawData:   rawData,
	}

	// Convert the resource to JSON
	jsonString, err := json.Marshal(resource)
	if err != nil {
		return nil, err
	}

	// Get the URL
	handleUrl := getHandleUrl()
	fullUrl := client.BaseURL.ResolveReference(handleUrl)

	// Create the HTTP request
	req, err := http.NewRequest("POST", fullUrl.String(), bytes.NewBuffer(jsonString))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	return req, nil
}
