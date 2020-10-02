package commands

import "github.com/Necroforger/dgrouter/exrouter"

type Command struct {
	Name        string
	Description string
	Handler     exrouter.HandlerFunc
}

func NewCommand(name string, description string, handler exrouter.HandlerFunc) *Command {
	command := &Command{
		Name:        name,
		Description: description,
		Handler:     handler,
	}
	addCommand(command)
	return command
}
