package entities_test

import (
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/persistence/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockName(t *testing.T) {
	name := (&entities.Mock{}).Name()
	assert.Equal(t, []byte(commons.MockEntityName), name)
}

func TestMockKey(t *testing.T) {
	key, err := (&entities.Mock{
		Tag:    "1234",
		Method: "GET",
		Path:   "/naruto",
		Status: 200,
		Body:   "naruto-rocks",
	}).Key()
	assert.NoError(t, err)
	assert.Equal(
		t,
		[]byte("{\"tag\":\"1234\",\"method\":\"GET\",\"path\":\"/naruto\"}\n"),
		key,
	)
}

func TestMockEncode(t *testing.T) {
	bytes, err := (&entities.Mock{
		Tag:    "1234",
		Method: "GET",
		Path:   "/naruto",
		Status: 200,
		Body:   "naruto-rocks",
	}).Encode()
	assert.NoError(t, err)
	assert.Equal(
		t,
		[]byte("{\"tag\":\"1234\",\"method\":\"GET\",\"path\":"+
			"\"/naruto\",\"status\":200,\"body\":\"naruto-rocks\"}\n"),
		bytes,
	)
}

func TestMockDecode(t *testing.T) {
	mock := &entities.Mock{}
	err := mock.Decode([]byte("{\"tag\":\"1234\",\"method\":\"GET\",\"path\":" +
		"\"/naruto\",\"status\":200,\"body\":\"naruto-rocks\"}\n"))
	assert.NoError(t, err)
	assert.Equal(t, "1234", mock.Tag)
	assert.Equal(t, "GET", mock.Method)
	assert.Equal(t, "/naruto", mock.Path)
	assert.Equal(t, 200, mock.Status)
	assert.Equal(t, "naruto-rocks", mock.Body)
}
