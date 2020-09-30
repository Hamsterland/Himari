package commands

import (
	"github.com/Necroforger/dgrouter/exrouter"
)

func pingCommand(ctx *exrouter.Context) {
	_, _ = ctx.Reply("Pong")
}
