package commands

import (
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
)

func avatarCommand(ctx *exrouter.Context) {

	mentions := ctx.Msg.Mentions
	args := ctx.Args

	var name string
	var url string

	if len(mentions) == 0 {
		if len(args) == 2 {
			user, err := ctx.Ses.User(ctx.Args[1])
			if err != nil || user == nil {
				replyNotFound(ctx.Msg.Author, ctx.Msg.ChannelID, ctx.Ses)
				return
			}
			name = user.Username + "'s Avatar"
			url = user.AvatarURL("")
		} else if len(args) > 2 {
			replyMoreThanOne(ctx.Msg.Author, ctx.Msg.ChannelID, ctx.Ses)
			return
		} else {
			name = ctx.Msg.Author.Username + "'s Avatar"
			url = ctx.Msg.Author.AvatarURL("")
		}
	} else if len(mentions) == 1 {
		name = mentions[0].Username + "'s Avatar"
		url = mentions[0].AvatarURL("")
	} else {
		replyMoreThanOne(ctx.Msg.Author, ctx.Msg.ChannelID, ctx.Ses)
		return
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    name,
			IconURL: url,
		},
		Image: &discordgo.MessageEmbedImage{
			URL: url,
		},
		Description: url,
	}

	_, _ = ctx.Ses.ChannelMessageSendEmbed(ctx.Msg.ChannelID, embed)
}

func replyNotFound(user *discordgo.User, chnId string, ses *discordgo.Session) {
	_, _ = ses.ChannelMessageSend(chnId, user.Mention()+" I cannot find the user you queried.")
}

func replyMoreThanOne(user *discordgo.User, chnId string, ses *discordgo.Session) {
	_, _ = ses.ChannelMessageSend(chnId, user.Mention()+" I cannot find the Avatar for more than one user!")
}
