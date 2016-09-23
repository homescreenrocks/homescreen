package storage

import "fmt"

type KeyNotFound struct {
	Key string `json:"key"`
}

func (err KeyNotFound) Error() string {
	return fmt.Sprintf("Key '%s' not found in datastore", err.Key)
}

func (err KeyNotFound) String() string {
	return err.Error()
}
