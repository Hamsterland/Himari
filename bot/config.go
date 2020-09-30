package bot

import (
	"encoding/json"
	"errors"
	"os"
)

type Configuration struct {
	Token  string `json:"Token"`
	Prefix string `json:"Prefix"`
}

func (c *Configuration) parse() error {

	file, err := os.Open("config.json")

	if err != nil {
		return errors.New("could not open config.json")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)

	defer file.Close()

	if err != nil {
		return errors.New("could not decode config.json")
	}

	return nil
}
