package commands

import (
	"Himari/nhentai"
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"strconv"
	"strings"
)

func nhentaiInfoCommand(ctx *exrouter.Context) {
	args := Truncate(ctx.Args)
	mention := ctx.Msg.Author.Mention()

	c, _ := ctx.Channel(ctx.Msg.ChannelID)
	if !c.NSFW {
		_, _ = ctx.Reply(fmt.Sprintf("%s this command can only be used in an NSFW chnanel.", mention))
		return
	}

	if len(args) == 0 {
		_, _ = ctx.Reply(fmt.Sprintf("%s please specify an Id.", mention))
		return
	}

	if len(args) > 1 {
		_, _ = ctx.Reply(fmt.Sprintf("%s please only provide one Id.", mention))
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		_, _ = ctx.Reply(fmt.Sprintf("%s please provide a valid Id.", mention))
		return
	}

	book, err := nhentai.SearchBook(id)
	if err != nil {
		_, _ = ctx.Reply(fmt.Sprintf("%s could not find a doujin with Id %d", mention, id))
		return
	}

	var title string

	// Try set the book title.
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

		link := fmt.Sprintf("[%s](https://nhentai.net%s)", tag.Name, tag.Url)

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
		SetURL(fmt.Sprintf("https://nhentai.net/g/%d/", book.ID)).
		SetImage(cover)

	if len(characters) > 0 {
		embed.AddField("Characters", strings.Join(characters, ", "), true)
	}

	if len(languages) > 0 {
		embed.AddField("Languages", strings.Join(languages, ", "), true)
	}

	if len(categories) > 0 {
		embed.AddField("Categories", strings.Join(categories, ", "), true)
	}

	if len(groups) > 0 {
		embed.AddField("Groups", strings.Join(groups, ", "), true)
	}

	if book.Scanlator != "" {
		embed.AddField("Scanlator", book.Scanlator, true)
	}

	embed.
		AddField("Pages", strconv.Itoa(book.TotalPages), true).
		AddField("Favourites", fmt.Sprintf("%d Favourites", book.TotalFavorites), true)

	if len(tags) > 0 {

		var removed bool
		str := strings.Join(tags, ", ")
		for len(str) > 924 {
			tags = tags[:len(tags)-1]
			str = strings.Join(tags, ", ")
			removed = true
			if len(str) > 924 {
				tags = tags[:len(tags)-1]
			}
		}

		var space string
		if removed {
			str = strings.Join(tags, ", ")
			space = fmt.Sprintf("\nSome tags could not be displayed due to space limitations.\nPlease use the `nhtags %d` command to see more.", book.ID)
		}

		final := fmt.Sprintf("%s %s", str, space)
		embed.AddField("Tags", final, false)
	}

	_, _ = ctx.Ses.ChannelMessageSendEmbed(ctx.Msg.ChannelID, embed.MessageEmbed)
}

func init() {
	NewCommand("nhinfo", "Searches for a doujin by Id.", nhentaiInfoCommand)
}
