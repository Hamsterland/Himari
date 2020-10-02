package commands

import (
	"Himari/nhentai"
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"strconv"
	"strings"
	"time"
)

func nhentaiInfoCommand(ctx *exrouter.Context) {
	args := Truncate(ctx.Args)
	author := ctx.Msg.Author

	c, _ := ctx.Channel(ctx.Msg.ChannelID)
	if !c.NSFW {
		_, _ = ctx.Reply(fmt.Sprintf("%s this command can only be used in an NSFW chnanel.", author.Mention()))
	}

	if len(args) == 0 {
		_, _ = ctx.Reply(fmt.Sprintf("%s please specify an Id.", author.Mention()))
		return
	}

	if len(args) > 1 {
		_, _ = ctx.Reply(fmt.Sprintf("%s please only provide one Id.", author.Mention()))
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		_, _ = ctx.Reply(fmt.Sprintf("%s please provide a valid Id.", author.Mention()))
		return
	}

	book, err := nhentai.SearchBook(id)
	if err != nil {
		_, _ = ctx.Reply(fmt.Sprintf("%s could not find a doujin with Id %d", author.Mention(), id))
		return
	}

	var title string

	if book.Titles.Pretty != "" {
		title = book.Titles.Pretty
	} else if book.Titles.English != "" {
		title = book.Titles.English
	} else if book.Titles.Japanese != "" {
		title = book.Titles.Japanese
	}

	cover, _ := book.GetCoverUrl()

	var tags []string
	for _, tag := range book.Tags {
		tags = append(tags, tag.Name)
	}

	embed := NewEmbed().
		SetTitle(title).
		SetURL(fmt.Sprintf("https://nhentai.net/g/%d/", book.Id)).
		SetImage(cover)

	if book.Scanlator != "" {
		embed.AddField("Scanlator", book.Scanlator, false)
	}

	if len(book.Tags) > 0 {
		embed.AddField("Tags", strings.Join(tags, ", "), false)
	}

	embed.
		AddField("Uploaded", book.GetUtcUploadDate().Format(time.RFC850), true).
		AddField("Favourites", fmt.Sprintf("%d", book.TotalFavorites), true).
		AddField("Pages", fmt.Sprintf("%d", book.GetPageCount()), true)

	_, _ = ctx.Ses.ChannelMessageSendEmbed(ctx.Msg.ChannelID, embed.MessageEmbed)
}

func init() {
	NewCommand("nhinfo", "Searches for a doujin by Id.", nhentaiInfoCommand)
}
