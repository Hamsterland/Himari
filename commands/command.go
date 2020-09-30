package commands

import "github.com/Necroforger/dgrouter/exrouter"

// Command represents a generic Himari command.
type Command struct {
	Name        string
	Description string
	Alias       []string
	Handler     exrouter.HandlerFunc
}
