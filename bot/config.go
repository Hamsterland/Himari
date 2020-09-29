package bot

import (
	"encoding/json"
	"errors"
	"os"
)

type configuration struct {
	Token string `json:"token"`
}

func (c *configuration) parse() error {

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(c)

	defer file.Close()

	if err != nil {
		return errors.New("could not decode configuration.json")
	}

	return nil
}
