package bot

import (
	"Himari/commands"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func Start() {
	configuration := &configuration{}
	configuration.parse()

	bot, err := discordgo.New("Bot " + configuration.Token)

	if err != nil {
		fmt.Println("error creating session", err)
		return
	}

	commands.RegisterRouter(bot)

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
