package bot

import (
	"Himari/commands"
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	configuration := &Configuration{}
	configuration.parse()

	session, err := discordgo.New("Bot " + configuration.Token)

	if err != nil {
		fmt.Println("error creating session", err)
		return
	}

	router := exrouter.New()
	commands.Initialize(router, session)

	err = session.Open()

	if err != nil {
		fmt.Println("error starting session", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = session.Close()
}
