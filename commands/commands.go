package commands

import (
	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

func RegisterRouter(session *discordgo.Session) {

	router := dgc.Create(&dgc.Router{
		Prefixes: []string{"!"},
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "Ping",
		Description: "A simple call-and-response",
		Usage:       "ping",
		Example:     "ping",
		IgnoreCase:  true,
		Handler:     pingCommand,
	})

	//router.RegisterCmd(&dgc.Command{
	//	Name: "Echo",
	//	Description: "Echoes a message",
	//	Usage: "echo",
	//	Example: "echo Hello World",
	//	IgnoreCase: true,
	//	Handler: echoCommand,
	//})

	router.Initialize(session)
}
