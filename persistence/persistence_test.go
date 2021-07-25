package persistence_test

import (
	"errors"
	"github.com/sinhashubham95/moxy/persistence"
	"github.com/sinhashubham95/moxy/persistence/entities"
	"github.com/stretchr/testify/assert"
	bolt "go.etcd.io/bbolt"
	"testing"
)

var keyError = errors.New("test key error")
var encodeError = errors.New("test encode error")
var decodeError = errors.New("test decode error")

// Naruto Entity
type Naruto struct {
	Naruto int `json:"naruto"`
}

func (t *Naruto) Name() []byte {
	if t.Naruto == 0 {
		return []byte("")
	}
	return []byte("naruto")
}

func (t *Naruto) Key() ([]byte, error) {
	if t.Naruto == 1 {
		return nil, keyError
	}
	return []byte("naruto"), nil
}

func (t *Naruto) Encode() ([]byte, error) {
	if t.Naruto == 2 {
		return nil, encodeError
	}
	return []byte("naruto"), nil
}

func (t *Naruto) Decode([]byte) error {
	if t.Naruto == 3 {
		return decodeError
	}
	return nil
}

func TestSave(t *testing.T) {
	mock := &entities.Mock{
		Tag:    "1234",
		Method: "GET",
		Path:   "/naruto",
		Status: 200,
		Body:   "naruto-rocks",
	}
	err := persistence.Save(mock)
	assert.NoError(t, err)

	savedMock := &entities.Mock{
		Tag:    "1234",
		Method: "GET",
		Path:   "/naruto",
	}
	err = persistence.View(savedMock)
	assert.NoError(t, err)
	assert.Equal(t, mock, savedMock)
}

func TestSaveCreateBucketError(t *testing.T) {
	naruto := &Naruto{Naruto: 0}
	err := persistence.Save(naruto)
	assert.Error(t, err)
	assert.Equal(t, bolt.ErrBucketNameRequired, err)
}

func TestSaveKeyError(t *testing.T) {
	naruto := &Naruto{Naruto: 1}
	err := persistence.Save(naruto)
	assert.Error(t, err)
	assert.Equal(t, keyError, err)
}

func TestSaveEncodeError(t *testing.T) {
	naruto := &Naruto{Naruto: 2}
	err := persistence.Save(naruto)
	assert.Error(t, err)
	assert.Equal(t, encodeError, err)
}

func TestView(t *testing.T) {
	mock := &entities.Mock{
		Tag:    "1234",
		Method: "GET",
		Path:   "/naruto",
		Status: 200,
		Body:   "naruto-rocks",
	}
	err := persistence.Save(mock)
	assert.NoError(t, err)

	savedMock := &entities.Mock{
		Tag:    "1234",
		Method: "GET",
		Path:   "/naruto",
	}
	err = persistence.View(savedMock)
	assert.NoError(t, err)
	assert.Equal(t, mock, savedMock)
}

func TestViewBucketNotExists(t *testing.T) {
	naruto := &Naruto{Naruto: 0}
	err := persistence.View(naruto)
	assert.Error(t, err)
	assert.Equal(t, persistence.ErrEntityNotFound, err)
}

func TestViewKeyError(t *testing.T) {
	naruto := &Naruto{Naruto: 4}
	err := persistence.Save(naruto)
	assert.NoError(t, err)
	defer func(t *testing.T, entity persistence.Entity) {
		assert.NoError(t, persistence.Delete(entity))
	}(t, naruto)
	naruto = &Naruto{Naruto: 1}
	err = persistence.View(naruto)
	assert.Error(t, err)
	assert.Equal(t, keyError, err)
}

func TestViewRecordNotFoundError(t *testing.T) {
	naruto := &Naruto{Naruto: 4}
	err := persistence.View(naruto)
	assert.Error(t, err)
	assert.Equal(t, persistence.ErrRecordNotFound, err)
}

func TestViewDecodeJSONError(t *testing.T) {
	naruto := &Naruto{Naruto: 3}
	err := persistence.Save(naruto)
	assert.NoError(t, err)
	defer func(t *testing.T, entity persistence.Entity) {
		assert.NoError(t, persistence.Delete(entity))
	}(t, naruto)
	err = persistence.View(naruto)
	assert.Error(t, err)
	assert.Equal(t, decodeError, err)
}

func TestDelete(t *testing.T) {
	mock := &entities.Mock{
		Tag:    "1234",
		Method: "GET",
		Path:   "/naruto",
		Status: 200,
		Body:   "naruto-rocks",
	}
	err := persistence.Save(mock)
	assert.NoError(t, err)
	err = persistence.Delete(mock)
	assert.NoError(t, err)
	savedMock := &entities.Mock{
		Tag:    "1234",
		Method: "GET",
		Path:   "/naruto",
	}
	err = persistence.View(savedMock)
	assert.Error(t, err)
	assert.Equal(t, persistence.ErrRecordNotFound, err)
}

func TestDeleteBucketNotExists(t *testing.T) {
	naruto := &Naruto{Naruto: 0}
	err := persistence.Delete(naruto)
	assert.Error(t, err)
	assert.Equal(t, persistence.ErrEntityNotFound, err)
}

func TestDeleteKeyError(t *testing.T) {
	naruto := &Naruto{Naruto: 1}
	err := persistence.Delete(naruto)
	assert.Error(t, err)
	assert.Equal(t, keyError, err)
}

func TestClose(t *testing.T) {
	persistence.Close()
	persistence.Close()
}
