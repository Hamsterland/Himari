package commands

import (
	"Himari/config"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
)

var Commands []*Command

func Initialize(router *exrouter.Route, session *discordgo.Session) {
	enableCommands(router)
	registerHandler(router, session)
}

func add(command *Command) {
	Commands = append(Commands, command)
}

func enableCommands(router *exrouter.Route) {
	for _, command := range Commands {
		router.On(command.Name, command.Handler).Desc(command.Description)
	}
}

func registerHandler(router *exrouter.Route, session *discordgo.Session) {
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		_ = router.FindAndExecute(s, config.Prefix, s.State.User.ID, m.Message)
	})
}
