package commands

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"strings"
)

func echoCommand(ctx *exrouter.Context) {
	var sb strings.Builder
	var args []string = ctx.Args
	for _, arg := range args {
		if arg == ctx.Route.Name {
			continue
		} else {
			sb.WriteString(" " + arg)
		}
	}
	echo := sb.String()
	_, _ = ctx.Reply(echo)
}
