package main

import (
	"Himari/commands"
	"Himari/config"
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"regexp"
	"syscall"
)

func main() {

	// Create the application configuration.
	config.Create()

	// Start a new session.
	session, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("error creating session", err)
		return
	}

	// Initialize the command router.
	router := exrouter.New()
	commands.Initialize(router, session)

	// Open the session.
	err = session.Open()

	if err != nil {
		fmt.Println("error starting session", err)
		return
	}

	// Temporary MessageCreate handler.
	session.AddHandler(imagesOnly)

	// Wait for Ctrl + C to close.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close
	_ = session.Close()
}

func imagesOnly(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Check if the message came from the Nino channel.
	if m.ChannelID != "760935127303192616" {
		return
	}

	// If it is a webhook, it's cool.
	if m.WebhookID != "" {
		return
	}

	// Check if the message has an image url. If so, it's fine.
	match, _ := regexp.MatchString("(http(s?):)([/|.|\\w|\\s|-])*\\.(?:jpg|gif|png)", m.Content)
	if match {
		return
	}

	// Otherwise, check if there are attachments. If so, it's also fine.
	if len(m.Attachments) > 0 {
		return
	}

	// But otherwise, delete the message because we don't want it.
	_ = s.ChannelMessageDelete(m.ChannelID, m.ID)
}
