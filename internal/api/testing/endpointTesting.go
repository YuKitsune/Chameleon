package testing

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yukitsune/chameleon/pkg/testUtils"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type CreateResourceFunc func(t *testing.T) (interface{}, []byte, error)
type UpdateResourceFunc func (original []byte) ([]byte, error)
type AssertFunc func(t *testing.T, expected []byte, actual []byte) error

func TestCreate(t *testing.T, endpoint string, makeResource CreateResourceFunc, assertion AssertFunc) {

	// Create the resource
	_, originalBytes, err := makeResource(t)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to make resource")
		return
	}

	// Send the create request
	resBytes, err := createViaApi(t, endpoint, originalBytes)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to create resource via API")
		return
	}

	// Assert
	err = assertion(t, originalBytes, resBytes)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to assert")
		return
	}
}

func TestGet(t *testing.T, endpoint string, makeResource CreateResourceFunc, params url.Values, assertion AssertFunc) {

	// Create the model
	_, originalBytes, err := makeResource(t)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail making alias.")
		return
	}

	// Send the create request
	_, err = createViaApi(t, endpoint, originalBytes)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to create resource via API")
		return
	}

	// Send the get request
	reqUrl := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	res, err := http.Get(reqUrl)
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
	if ok := assertStatusCode(t, http.StatusOK, res.StatusCode, resBytes); !ok {
		return
	}

	// Assert
	err = assertion(t, originalBytes, resBytes)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to assert")
		return
	}
}

func TestUpdate(t *testing.T, endpoint string, makeResource CreateResourceFunc, updateResource UpdateResourceFunc, assertion AssertFunc) {

	// Create the model
	_, originalBytes, err := makeResource(t)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail making alias.")
		return
	}

	// Send the create request
	createdBytes, err := createViaApi(t, endpoint, originalBytes)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to create resource via API")
		return
	}

	// Update the resource
	updatedBytes, err := updateResource(createdBytes)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to update resource")
		return
	}

	// Send the update request
	putReq, err := http.NewRequest("PUT", endpoint, bytes.NewReader(updatedBytes))
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
	if ok := assertStatusCode(t, http.StatusOK, res.StatusCode, resBytes); !ok {
		return
	}

	// Assert
	err = assertion(t, originalBytes, resBytes)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to assert")
		return
	}
}

func TestDelete(t *testing.T, endpoint string, makeResource CreateResourceFunc) {

	// Create the model
	_, originalBytes, err := makeResource(t)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail making alias.")
		return
	}

	// Send the create request
	createdBytes, err := createViaApi(t, endpoint, originalBytes)
	if err != nil {
		testUtils.FailOnError(t, err, "Should not fail to create resource via API")
		return
	}

	fmt.Println(string(createdBytes))

	// Send the delete request
	putReq, err := http.NewRequest("DELETE", endpoint, bytes.NewReader(createdBytes))
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
	if ok := assertStatusCode(t, http.StatusOK, res.StatusCode, resBytes); !ok {
		return
	}

	// Todo: Assert no content in response body and 404 when sending a GET request
}

func createViaApi(t *testing.T, endpoint string, modelBytes []byte) ([]byte, error) {

	// Send the create request
	res, err := http.Post(endpoint, "application/json", bytes.NewReader(modelBytes))
	if err != nil {
		return nil, err
	}

	defer func (){ _ = res.Body.Close() }()

	// Read the response body
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Make sure we have the right status code
	if ok := assertStatusCode(t, http.StatusCreated, res.StatusCode, resBytes); !ok {
		return nil, fmt.Errorf("assertion failure")
	}

	return resBytes, nil
}

func assertStatusCode(t *testing.T, expected int, actual int, resBytes []byte) bool {
	return assert.Equal(t, expected, actual, "API should return %d but found %d. Response body: %s", expected, actual, string(resBytes))
}
