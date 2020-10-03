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
		return
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
	} else {
		title = "Himari Error: Title Not Found (Chinese and Korean Titles Unsupported)"
	}

	cover, _ := book.GetCoverUrl()

	var tags []string
	var artists []string
	var characters []string
	var languages []string
	var groups []string
	var categories []string

	for _, tag := range book.Tags {

		// Get the markdown link
		link := fmt.Sprintf("[%s](https://nhentai.net%s)", tag.Name, tag.Url)

		// Find specific information.
		switch tag.Type {
		case "artist":
			artists = append(artists, link)
		case "character":
			characters = append(characters, link)
		case "language":
			languages = append(languages, link)
		case "group":
			groups = append(groups, link)
		case "category":
			categories = append(categories, link)
		default:
			tags = append(tags, link)
		}
	}

	embed := NewEmbed().
		SetTitle(title).
		SetURL(fmt.Sprintf("https://nhentai.net/g/%d/", book.Id)).
		SetImage(cover)

	if book.Scanlator != "" {
		embed.AddField("Scanlator", book.Scanlator, true)
	}

	if len(tags) > 0 {
		embed.AddField("Tags", strings.Join(tags, ", "), true)
	}

	if len(characters) > 0 {
		embed.AddField("Characters", strings.Join(characters, ", "), true)
	}

	if len(languages) > 0 {
		embed.AddField("Languages", strings.Join(languages, ", "), true)
	}

	if len(groups) > 0 {
		embed.AddField("Groups", strings.Join(groups, ", "), true)
	}

	if len(categories) > 0 {
		embed.AddField("Categories", strings.Join(groups, ", "), true)
	}

	embed.
		AddField("Uploaded", book.GetUtcUploadDate().Format(time.RFC850), true).
		AddField("Favourites", fmt.Sprintf("%d", book.TotalFavorites), true).
		AddField("Pages", fmt.Sprintf("%d", book.GetPageCount()), false)

	_, _ = ctx.Ses.ChannelMessageSendEmbed(ctx.Msg.ChannelID, embed.MessageEmbed)
}

func init() {
	NewCommand("nhinfo", "Searches for a doujin by Id.", nhentaiInfoCommand)
}
