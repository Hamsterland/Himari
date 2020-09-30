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

	bot, err := discordgo.New("Bot " + configuration.Token)

	if err != nil {
		fmt.Println("error creating session", err)
		return
	}

	router := exrouter.New()
	commands.RegisterCommands(router)
	commands.RegisterHandler(router, bot)

	err = bot.Open()

	if err != nil {
		fmt.Println("error starting bot", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = bot.Close()
}
