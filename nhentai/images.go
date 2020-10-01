package nhentai

import (
	"errors"
)

type Images struct {
	Pages     []Page `json:"pages"`
	Cover     Page   `json:"cover"`
	Thumbnail Page   `json:"thumbnail"`
}

type Page struct {
	FileType string `json:"t"`
	Width    int    `json:"w"`
	Height   int    `json:"h"`
}

type FileType string

const (
	Jpg FileType = "jpg"
	Png FileType = "png"
	Gif FileType = "gif"
)

func (page *Page) GetPageFileType() (FileType, error) {
	switch page.FileType {
	case "j":
		return Jpg, nil
	case "p":
		return Png, nil
	case "g":
		return Gif, nil
	default:
		return "", errors.New("unknown file type: " + page.FileType)
	}
}
