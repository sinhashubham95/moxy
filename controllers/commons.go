package controllers

import (
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/persistence"
)

// EncodeJSON is used to encode the given interface into json bytes
var EncodeJSON = commons.EncodeJSON

// DecodeJSON is used to decode the given bytes into interface
var DecodeJSON = commons.DecodeJSON

// PersistenceSave is used to save the entity
var PersistenceSave = persistence.Save

// PersistenceView is used to view the entity
var PersistenceView = persistence.View

// PersistenceDelete is used to delete the entity
var PersistenceDelete = persistence.Delete
