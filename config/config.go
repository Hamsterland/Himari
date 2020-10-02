package config

import (
	"encoding/json"
	"os"
)

var (
	Token  string
	Prefix string
)

func Create() {
	c := get()
	Token = c.Token
	Prefix = c.Prefix
}

type configuration struct {
	Token  string `json:"Token"`
	Prefix string `json:"Prefix"`
}

func get() *configuration {

	// Declare a configuration.
	c := &configuration{}

	// Open the file.
	file, _ := os.Open("config.json")

	// Declare a decoder and deserialize.
	decoder := json.NewDecoder(file)
	_ = decoder.Decode(c)

	// Close the file to prevent access conflicts and return.
	defer file.Close()
	return c
}
