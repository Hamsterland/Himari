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

	router.On(ping.Name, ping.Handler).Desc(ping.Description)

	avatar := Command{
		Name:        "avatar",
		Description: "Retrieves your avatar",
		Handler:     avatarCommand,
	}

	router.On(avatar.Name, avatar.Handler).Desc(avatar.Description)
}

func RegisterHandler(router *exrouter.Route, session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		_ = router.FindAndExecute(s, "!", s.State.User.ID, m.Message)
	})
}
