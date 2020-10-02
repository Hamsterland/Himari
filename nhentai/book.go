package nhentai

import (
	"errors"
	"fmt"
	"time"
)

type Search struct {
	Books             []Book `json:"result"`
	TotalResultPages  int    `json:"num_pages"`
	TotalBooksPerPage int    `json:"per_page"`
}

type Book struct {
	Id             int32  `json:"id"`
	MediaId        string `json:"media_id"`
	Titles         Titles `json:"title"`
	Images         Images `json:"images"`
	Scanlator      string `json:"scanlator"`
	UnixUploadDate int64  `json:"upload_date"`
	Tags           []Tag  `json:"tags"`
	TotalPages     int32  `json:"num_pages"`
	TotalFavorites int32  `json:"num_favorites"`
}

func (book *Book) GetAllUrls() ([]string, error) {
	var pages []string
	for i := range book.Images.Pages {
		page, err := book.GetPageUrl(i + 1)
		if err != nil {
			return nil, err
		}

		pages = append(pages, page)
	}
	return pages, nil
}

func (book *Book) GetPageUrl(page int) (string, error) {
	if page < 0 || page > book.GetPageCount() {
		return "", errors.New("index out of range")
	}

	fileType, err := book.Images.Pages[page-1].GetPageFileType()
	if err != nil {
		return "", err
	}

	return ImageEndpoint + fmt.Sprintf(PageSearchEndpoint, book.MediaId, page, fileType), nil
}

func (book *Book) GetCoverUrl() (string, error) {
	fileType, err := book.Images.Cover.GetPageFileType()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(CoverImageEndpoint, book.MediaId, fileType), nil
}

func (book *Book) GetPageCount() int {
	return len(book.Images.Pages) + 2
}

func (book *Book) GetUtcUploadDate() time.Time {
	return time.Unix(book.UnixUploadDate, 0)
}
