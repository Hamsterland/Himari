package commands

import (
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"strings"
)

func echoCommand(ctx *exrouter.Context) {
	args := Truncate(ctx.Args)
	echo := strings.Join(args, " ")
	final := fmt.Sprintf("**%s:** %s", ctx.Msg.Author, echo)
	_, _ = ctx.Reply(final)
}
