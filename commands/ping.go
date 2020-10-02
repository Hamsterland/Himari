package commands

import (
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"time"
)

func pingCommand(ctx *exrouter.Context) {
	start := time.Now()
	msg, _ := ctx.Reply("Pinging...")
	end := time.Since(start)
	content := fmt.Sprintf("Pong! (%dms)", end.Milliseconds())
	_, _ = ctx.Ses.ChannelMessageEdit(msg.ChannelID, msg.ID, content)
}

func init() {
	NewCommand("ping", "Sends a ping to Himari", pingCommand)
}
