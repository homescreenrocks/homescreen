package storage

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/gin-gonic/gin"
	"github.com/homescreenrocks/homescreen/shared"
)

var defaultBucket = []byte("homescreen")

type Storage struct {
	db *bolt.DB
}

func New(path string) (*Storage, error) {
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(defaultBucket)
		return err
	})
	if err != nil {
		return nil, err
	}

	return &Storage{db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) SetRaw(key string, value []byte) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		b.Put([]byte(key), value)
		return nil
	})
	return err
}

func (s *Storage) Set(key string, value interface{}) error {
	raw, err := json.Marshal(&value)
	if err != nil {
		return err
	}
	return s.SetRaw(key, raw)
}

func (s *Storage) GetRaw(key string) ([]byte, error) {
	var value []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(defaultBucket)
		value = b.Get([]byte(key))
		return nil
	})

	if value == nil {
		return nil, KeyNotFound{key}
	}

	return value, err
}

func (s *Storage) Get(key string, value interface{}) error {
	raw, err := s.GetRaw(key)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, value)
}

func (s *Storage) RegisterRouterGroup(group *gin.RouterGroup) {
	group.GET("/*key", func(c *gin.Context) {
		var data interface{}

		key := strings.Trim(c.Param("key"), "/")
		err := s.Get(key, &data)
		if err != nil {
			if _, ok := err.(KeyNotFound); ok {
				c.JSON(404, shared.HttpError{"Key not found.", nil})
				return
			}

			c.JSON(500, shared.HttpError{"Failed to read value from datastore.", err})
			return
		}

		c.JSON(200, data)
	})

	group.PUT("/*key", func(c *gin.Context) {
		var data interface{}

		err := c.BindJSON(&data)
		if err != nil {
			c.JSON(400, shared.HttpError{"Failed to decode request body.", err})
			return
		}

		key := strings.Trim(c.Param("key"), "/")
		log.Printf(key)
		err = s.Set(key, data)
		if err != nil {
			c.JSON(500, shared.HttpError{"Failed to write value to datastore.", err})
			return
		}
	})
}
