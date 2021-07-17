package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/pkg/testUtils"
	"io"
	"net/http"
	"testing"
)

var apiUrl = "http://localhost:8000"

// Todo: Should come through and make these a bit more generic. E.g. have a test method for each request type (POST,
// 	GET, PUT, DELETE) then sub in a URL, object, assertion callback, etc.

func TestCreateAlias(t *testing.T) {

	// Create an alias
	reqAlias, _, reader, err := makeAlias(t)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail making alias.")
		return
	}

	// Send the create request
	res, err := http.Post(fmt.Sprintf("%s/alias", apiUrl), "application/json", reader)
	if err != nil {
		testUtils.FailOnError(t, err, "API should not return an error")
		return
	}
	defer func (){ _ = res.Body.Close() }()

	// Read the response body
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to read body")
		return
	}

	// Make sure we have the right status code
	if res.StatusCode != http.StatusCreated {
		t.Logf("API should return 201, found %d. Response %s", res.StatusCode, string(resBytes))
		t.Fail()
		return
	}

	// Unmarshal the response
	var resAlias model.Alias
	err = json.Unmarshal(resBytes, &resAlias)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to unmarshal alias")
		return
	}

	// Assert all the things
	assert.Equal(t, resAlias.Username, reqAlias.Username)
	assert.Equal(t, resAlias.SenderWhitelistPattern, reqAlias.SenderWhitelistPattern)
}

func TestGetAlias(t *testing.T) {

	// Create an alias
	reqAlias, _, reader, err := makeAlias(t)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail making alias.")
		return
	}

	// Send the create request
	createRes, err := http.Post(fmt.Sprintf("%s/alias", apiUrl), "application/json", reader)
	if err != nil {
		testUtils.FailOnError(t, err, "API should not return an error")
		return
	}
	defer func (){ _ = createRes.Body.Close() }()

	// Read the response body
	createResBytes, err := io.ReadAll(createRes.Body)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to read body")
		return
	}

	// Make sure we have the right status code
	if createRes.StatusCode != http.StatusCreated {
		t.Logf("API should return 201, found %d. Response %s", createRes.StatusCode, string(createResBytes))
		t.Fail()
		return
	}

	// Send the get request
	senderAddr := fmt.Sprintf("test@%s.app", t.Name())
	res, err := http.Get(fmt.Sprintf("%s/alias?sender=%s&recipient=%s", apiUrl, senderAddr, reqAlias.Username))
	if err != nil {
		testUtils.FailOnError(t, err, "API should not return an error")
		return
	}
	defer func (){ _ = res.Body.Close() }()

	// Read the response body
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to read body")
		return
	}

	// Make sure we have the right status code
	if res.StatusCode != http.StatusOK {
		t.Logf("API should return 200, found %d. Response %s", res.StatusCode, string(resBytes))
		t.Fail()
		return
	}

	// Unmarshal the response
	var resAlias model.Alias
	err = json.Unmarshal(resBytes, &resAlias)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to unmarshal alias")
		return
	}

	// Assert all the things
	assert.Equal(t, resAlias.Username, reqAlias.Username)
	assert.Equal(t, resAlias.SenderWhitelistPattern, reqAlias.SenderWhitelistPattern)
}

func TestUpdateAlias(t *testing.T) {

	// Create an alias
	reqAlias, _, reader, err := makeAlias(t)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail making alias.")
		return
	}

	// Send the create request
	createRes, err := http.Post(fmt.Sprintf("%s/alias", apiUrl), "application/json", reader)
	if err != nil {
		testUtils.FailOnError(t, err, "API should not return an error")
		return
	}
	defer func (){ _ = createRes.Body.Close() }()

	// Read the response body
	createResBytes, err := io.ReadAll(createRes.Body)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to read body")
		return
	}

	// Make sure we have the right status code
	if createRes.StatusCode != http.StatusCreated {
		t.Logf("API should return 201, found %d. Response %s", createRes.StatusCode, string(createResBytes))
		t.Fail()
		return
	}

	// Unmarshal the response
	var createResAlias model.Alias
	err = json.Unmarshal(createResBytes, &createResAlias)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to unmarshal alias")
		return
	}

	createResAlias.SenderWhitelistPattern = fmt.Sprintf(".+@%s\\.com", t.Name())
	_, updateReqReader, err := prepareForRequest(createResAlias)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to prepare modify request")
		return
	}

	// Send the update request
	putReq, err := http.NewRequest("PUT", fmt.Sprintf("%s/alias", apiUrl), updateReqReader)
	res, err := http.DefaultClient.Do(putReq)
	if err != nil {
		testUtils.FailOnError(t, err, "API should not return an error")
		return
	}
	defer func (){ _ = res.Body.Close() }()

	// Read the response body
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to read body")
		return
	}

	// Make sure we have the right status code
	if res.StatusCode != http.StatusOK {
		t.Logf("API should return 200, found %d. Response %s", res.StatusCode, string(resBytes))
		t.Fail()
		return
	}

	// Unmarshal the response
	var resAlias model.Alias
	err = json.Unmarshal(resBytes, &resAlias)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to unmarshal alias")
		return
	}

	// Assert all the things
	assert.Equal(t, resAlias.Username, createResAlias.Username)
	assert.NotEqual(t, resAlias.SenderWhitelistPattern, reqAlias.SenderWhitelistPattern)
	assert.Equal(t, resAlias.SenderWhitelistPattern, createResAlias.SenderWhitelistPattern)
}

func TestDeleteAlias(t *testing.T) {

	// Create an alias
	_, _, reader, err := makeAlias(t)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail making alias.")
		return
	}

	// Send the create request
	createRes, err := http.Post(fmt.Sprintf("%s/alias", apiUrl), "application/json", reader)
	if err != nil {
		testUtils.FailOnError(t, err, "API should not return an error")
		return
	}
	defer func (){ _ = createRes.Body.Close() }()

	// Read the response body
	createResBytes, err := io.ReadAll(createRes.Body)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to read body")
		return
	}

	// Make sure we have the right status code
	if createRes.StatusCode != http.StatusCreated {
		t.Logf("API should return 201, found %d. Response %s", createRes.StatusCode, string(createResBytes))
		t.Fail()
		return
	}

	// Unmarshal the response
	var createResAlias model.Alias
	err = json.Unmarshal(createResBytes, &createResAlias)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to unmarshal alias")
		return
	}

	_, createResReader, err := prepareForRequest(createResAlias)

	// Send the delete request
	putReq, err := http.NewRequest("DELETE", fmt.Sprintf("%s/alias", apiUrl), createResReader)
	res, err := http.DefaultClient.Do(putReq)
	if err != nil {
		testUtils.FailOnError(t, err, "API should not return an error")
		return
	}
	defer func (){ _ = res.Body.Close() }()

	// Read the response body
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to read body")
		return
	}

	// Make sure we have the right status code
	if res.StatusCode != http.StatusOK {
		t.Logf("API should return 200, found %d. Response %s", res.StatusCode, string(resBytes))
		t.Fail()
		return
	}

	// Todo: Assert no content in response body and 404 when sending a GET request
}

func makeAlias(t *testing.T) (*model.Alias, []byte, io.Reader, error){
	alias := &model.Alias{
		Username:               "YuKitsune",
		SenderWhitelistPattern: fmt.Sprintf(".+@%s\\.app", t.Name()),
	}

	jsonBytes, reader, err := prepareForRequest(alias)
	if err != nil {
		return nil, nil, nil, err
	}

	return alias, jsonBytes, reader, nil
}

func prepareForRequest(v interface{}) ([]byte, io.Reader, error){

	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return nil, nil, err
	}

	buffer := bytes.NewBuffer(jsonBytes)
	return jsonBytes, buffer, nil
}