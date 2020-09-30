package commands

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(router *exrouter.Route) {

	ping := Command{
		Name:        "ping",
		Description: "Responds with \"Pong\"",
		Handler:     pingCommand,
	}

	echo := Command{
		Name:        "echo",
		Description: "Echoes a message",
		Handler:     echoCommand,
	}

	avatar := Command{
		Name:        "avatar",
		Description: "Retrieves an avatar",
		Handler:     avatarCommand,
	}

	commands := []Command{ping, echo, avatar}

	for _, command := range commands {
		router.On(command.Name, command.Handler).Desc(command.Description)
	}
}

func RegisterHandler(router *exrouter.Route, session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		_ = router.FindAndExecute(s, "!", s.State.User.ID, m.Message)
	})
}
