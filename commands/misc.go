package commands

import (
	"github.com/Lukaesebrot/dgc"
)

func pingCommand(ctx *dgc.Ctx) {
	_ = ctx.RespondText("Pong!")
}

//func echoCommand(ctx *dgc.Ctx) {
//	arguments := ctx.Arguments
//	arguments.AsSingle()
//	_ = ctx.RespondText(arguments.Raw())
//}
