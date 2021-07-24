package persistence

import (
	"errors"
	"github.com/sinhashubham95/moxy/flags"
	bolt "go.etcd.io/bbolt"
	"io/fs"
	"log"
)

// Entity is the set of methods which have to be implemented by any entity
type Entity interface {
	Name() ([]byte, error)
	Key() ([]byte, error)
	Encode() ([]byte, error)
	Decode(bytes []byte) error
}

var db *bolt.DB

func init() {
	var err error
	db, err = bolt.Open(flags.PersistencePath(), fs.ModePerm, nil)
	if err != nil {
		// opening the database is a mandatory thing
		// if unsuccessful, then panic the application to an error state
		panic(err)
	}
}

// Save is used to save the entry in the proper bucket with the proper key
func Save(entity Entity) error {
	return db.Update(func(tx *bolt.Tx) error {
		name, err := entity.Name()
		if err != nil {
			// error getting entity name
			return err
		}
		b, err := tx.CreateBucketIfNotExists(name)
		if err != nil {
			// error creating or fetching bucket
			return err
		}
		key, err := entity.Key()
		if err != nil {
			// error getting the key for this entity entry
			return err
		}
		value, err := entity.Encode()
		if err != nil {
			// error encoding the value to bytes
			return err
		}
		// finally save this in the bucket
		return b.Put(key, value)
	})
}

// View is used to fetch the entry in the proper bucket with the proper key
func View(entity Entity) error {
	return db.View(func(tx *bolt.Tx) error {
		name, err := entity.Name()
		if err != nil {
			// error getting entity name
			return err
		}
		b := tx.Bucket(name)
		if b == nil {
			return errors.New("entity not found")
		}
		key, err := entity.Key()
		if err != nil {
			// error getting the key for this entity entry
			return err
		}
		value := b.Get(key)
		if value == nil {
			return errors.New("record not found")
		}
		return entity.Decode(value)
	})
}

// Delete is used to delete the entry in the proper bucket with the proper key
func Delete(entity Entity) error {
	return db.Update(func(tx *bolt.Tx) error {
		name, err := entity.Name()
		if err != nil {
			// error getting entity name
			return err
		}
		b, err := tx.CreateBucketIfNotExists(name)
		if err != nil {
			// error creating or fetching bucket
			return err
		}
		key, err := entity.Key()
		if err != nil {
			// error getting the key for this entity entry
			return err
		}
		return b.Delete(key)
	})
}

// Close is used to close and clean the db interactions
func Close() {
	err := db.Close()
	if err != nil {
		log.Printf("Error occurred while closing the database %+v.", err)
	}
}
