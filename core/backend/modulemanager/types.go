package modulemanager

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/homescreenrocks/homescreen/shared"
)

func GetMetadata(moduleUrl string) (*shared.ModuleMetadata, error) {
	res, err := http.Get(fmt.Sprintf("%s/v1/metadata", moduleUrl))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status code %d", res.StatusCode)
	}

	metadata := new(shared.ModuleMetadata)
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

type HttpError struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
