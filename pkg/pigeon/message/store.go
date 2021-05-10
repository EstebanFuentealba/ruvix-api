package message

import (
	"encoding/json"
	"os"

	"github.com/boltdb/bolt"
	"github.com/oklog/ulid"
)

// BoltDatastore store data in db using bolt as a db backend
type BoltDatastore struct {
	DB *bolt.DB
}

var (
	// MsgBucket ...
	MsgBucket = []byte("messages")
)

// NewBoltDatastore returns a new datastore instance or an error if
// a datasore cannot be returned
func NewBoltDatastore(path string) (*BoltDatastore, error) {
	db, err := bolt.Open(path, os.ModePerm, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, berr := tx.CreateBucketIfNotExists(MsgBucket)
		return berr
	})
	if err != nil {
		return nil, err
	}

	return &BoltDatastore{
		DB: db,
	}, nil
}

// MessageStore ...
type MessageStore struct {
	Dst *BoltDatastore
}

// NewMessageStore ...
func NewMessageStore(dst *BoltDatastore) (*MessageStore, error) {
	return &MessageStore{
		Dst: dst,
	}, nil
}

// AddMessage ...
func (ss *MessageStore) AddMessage(m Message) error {
	err := ss.Dst.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(MsgBucket)

		k, merr := m.ID.MarshalBinary()
		if merr != nil {
			return merr
		}

		v, err := json.Marshal(m)
		if err != nil {
			return err
		}

		return b.Put(k, v)
	})
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (ss *MessageStore) Get(id ulid.ULID) (*Message, error) {
	msg := &Message{}

	err := ss.Dst.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(MsgBucket)

		k, err := id.MarshalBinary()
		if err != nil {
			return err
		}
		v := b.Get(k)

		err = json.Unmarshal(v, msg)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// UpdateContent ...
func (ss *MessageStore) UpdateContent(id ulid.ULID, content string) error {
	msg := &Message{}

	return ss.Dst.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(MsgBucket)

		k, err := id.MarshalBinary()
		if err != nil {
			return err
		}
		v := b.Get(k)

		err = json.Unmarshal(v, msg)
		if err != nil {
			return err
		}

		msg.Content = content

		v, err = json.Marshal(msg)
		if err != nil {
			return err
		}

		return b.Put(k, v)
	})
}

// UpdateStatus ...
func (ss *MessageStore) UpdateStatus(id ulid.ULID, status string) error {
	msg := &Message{}

	return ss.Dst.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(MsgBucket)

		k, err := id.MarshalBinary()
		if err != nil {
			return err
		}
		v := b.Get(k)

		err = json.Unmarshal(v, msg)
		if err != nil {
			return err
		}

		msg.Status = string(status)

		v, err = json.Marshal(msg)
		if err != nil {
			return err
		}

		return b.Put(k, v)
	})
}
