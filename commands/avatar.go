package commands

import (
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func avatarCommand(ctx *exrouter.Context) {
	mentions := ctx.Msg.Mentions
	args := Truncate(ctx.Args)

	// Map to hold found users and their avatars.
	// Maximum capacity of 8 for practical purposes.
	users := make(map[*discordgo.User]string)

	// If there are more than 8 provided arguments.
	if len(args) > 8 {
		_, _ = ctx.Reply(fmt.Sprintf("%s I cannot find the avatars for more than eight users!", ctx.Msg.Author.Mention()))
		return
	}

	// If there are no provided arguments.
	if len(args) == 0 {
		users[ctx.Msg.Author] = ctx.Msg.Author.AvatarURL("")
	}

	// If there are multiple provided arguments.
	if len(args) > 0 {

		// If there any mentions, add them to the map.
		if len(mentions) > 0 {
			for _, mention := range mentions {
				users[mention] = mention.AvatarURL("")
			}
		}

		// Otherwise, let's try parse the remaining arguments as users.
		for _, arg := range args {
			lookup, err := ctx.Ses.User(arg)
			if err != nil {
				continue
			}
			users[lookup] = lookup.AvatarURL("")
		}
	}

	// Finally, build the embeds.
	embed := NewEmbed()

	// If there is only one found user, make a single embed with an image.
	if len(users) == 1 {
		for user, url := range users {
			embed.SetTitle(fmt.Sprintf("%s's Avatar", user)).
				SetDescription(url).
				SetImage(url)
		}
	}

	// If there are multiple found users, put all of them into the embed without images.
	if len(users) > 1 {
		builder := strings.Builder{}
		for user, url := range users {
			builder.WriteString("\n\n")
			builder.WriteString(user.Mention() + " " + url)
		}

		title := fmt.Sprintf("%s I found these avatars!", ctx.Msg.Author)
		embed.SetTitle(title).SetDescription(builder.String())
	}

	// Send the embed to the channel.
	_, _ = ctx.Ses.ChannelMessageSendEmbed(ctx.Msg.ChannelID, embed.MessageEmbed)
}

func init() {
	NewCommand("avatar", "Finds user's avatars", avatarCommand)
}
