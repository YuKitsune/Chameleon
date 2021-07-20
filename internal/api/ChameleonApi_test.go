package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yukitsune/chameleon/internal/api/model"
	apitesting "github.com/yukitsune/chameleon/internal/api/testing"
	"testing"
)

var apiUrl = "http://localhost:8000/alias"

// Todo: Should come through and make these a bit more generic. E.g. have a test method for each request type (POST,
// 	GET, PUT, DELETE) then sub in a URL, object, assertion callback, etc.

func TestCreateAlias(t *testing.T) {
	apitesting.TestCreate(t, apiUrl, makeAlias, func(t *testing.T, expectedBytes []byte, actualBytes []byte) error {
		var expectedAlias model.Alias
		err := json.NewDecoder(bytes.NewReader(expectedBytes)).Decode(&expectedAlias)
		if err != nil {
			return err
		}

		var actualAlias model.Alias
		err = json.NewDecoder(bytes.NewReader(actualBytes)).Decode(&actualAlias)
		if err != nil {
			return err
		}

		assert.Equal(t, expectedAlias.Username, actualAlias.Username)
		assert.Equal(t, expectedAlias.SenderWhitelistPattern, actualAlias.SenderWhitelistPattern)

		return nil
	})
}

func TestGetAlias(t *testing.T) {

	params := map[string]string{
		"sender": fmt.Sprintf("test@%s.app", t.Name()),
		"recipient": "YuKitsune",
	}

	apitesting.TestGet(t, apiUrl, makeAlias, params, func(t *testing.T, expectedBytes []byte, actualBytes []byte) error {
		var expectedAlias model.Alias
		err := json.NewDecoder(bytes.NewReader(expectedBytes)).Decode(&expectedAlias)
		if err != nil {
			return err
		}

		var actualAlias model.Alias
		err = json.NewDecoder(bytes.NewReader(actualBytes)).Decode(&actualAlias)
		if err != nil {
			return err
		}

		assert.Equal(t, expectedAlias.Username, actualAlias.Username)
		assert.Equal(t, expectedAlias.SenderWhitelistPattern, actualAlias.SenderWhitelistPattern)
		return nil
	})
}

func TestUpdateAlias(t *testing.T) {

	var oldSenderWhitelistPattern string
	newSenderWhitelistPattern := fmt.Sprintf(".+@%s\\.com", t.Name())

	apitesting.TestUpdate(
		t,
		apiUrl,
		makeAlias,
		func (createdBytes []byte) ([]byte, error) {
			var createdAlias model.Alias
			err := json.NewDecoder(bytes.NewReader(createdBytes)).Decode(&createdAlias)
			if err != nil {
				return nil, err
			}

			oldSenderWhitelistPattern = createdAlias.SenderWhitelistPattern
			createdAlias.SenderWhitelistPattern = newSenderWhitelistPattern

			return json.Marshal(createdAlias)
		},
		func(t *testing.T, expectedBytes []byte, actualBytes []byte) error {
			var expectedAlias model.Alias
			err := json.NewDecoder(bytes.NewReader(expectedBytes)).Decode(&expectedAlias)
			if err != nil {
				return err
			}

			var actualAlias model.Alias
			err = json.NewDecoder(bytes.NewReader(actualBytes)).Decode(&actualAlias)
			if err != nil {
				return err
			}

			assert.Equal(t, expectedAlias.Username, actualAlias.Username)
			assert.NotEqual(t, oldSenderWhitelistPattern, actualAlias.SenderWhitelistPattern)
			assert.Equal(t, newSenderWhitelistPattern, actualAlias.SenderWhitelistPattern)
			return nil
		})
}

func TestDeleteAlias(t *testing.T) {
	apitesting.TestDelete(t, apiUrl, makeAlias)
}

func makeAlias(t *testing.T) (interface{}, []byte, error){
	alias := &model.Alias{
		Username:               "YuKitsune",
		SenderWhitelistPattern: fmt.Sprintf(".+@%s\\.app", t.Name()),
	}

	jsonBytes, err := json.Marshal(alias)
	if err != nil {
		return nil, nil, err
	}

	return alias, jsonBytes, nil
}
