package commands

import "github.com/Necroforger/dgrouter/exrouter"

type Command struct {
	Name        string
	Description string
	Handler     exrouter.HandlerFunc
}

func New(name string, description string, handler exrouter.HandlerFunc) *Command {
	command := &Command{
		Name:        name,
		Description: description,
		Handler:     handler,
	}
	add(command)
	return command
}
