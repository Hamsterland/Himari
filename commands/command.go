package commands

import "github.com/Necroforger/dgrouter/exrouter"

type Command struct {
	Name        string
	Description string
	Alias       []string
	Handler     exrouter.HandlerFunc
}
