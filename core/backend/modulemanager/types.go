package modulemanager

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/homescreenrocks/homescreen/core/backend/types"
)

type Module struct {
	PluginUrl string
	Metadata  *Metadata
	Settings  map[string]types.Setting
}

type Metadata struct {
	ID          string
	Name        string
	Version     string
	Description string
	Dir         string
}

func GetMetadata(pluginUrl string) (*Metadata, error) {
	res, err := http.Get(fmt.Sprintf("http://%s/v1/metadata", pluginUrl))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status code %d", res.StatusCode)
	}

	metadata := new(Metadata)
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

type RegisterRequest struct {
	PluginURL string `json:"plugin-url"`
}

type HttpError struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
