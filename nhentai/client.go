package nhentai

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BaseAddress        = "https://nhentai.net/"
	ImageEndpoint      = "https://i.nhentai.net"
	BookEndpoint       = "api/gallery/%d"
	BookSearchEndpoint = "api/galleries/search?query=%s&page=%d&sort=%s"
	PageSearchEndpoint = "/galleries/%s/%d.%s"
	CoverImageEndpoint = "https://t.nhentai.net/galleries/%s/cover.%s"
)

type Sort string

const (
	Popular Sort = "popular"
	Recent  Sort = "popular"
)

func SearchBook(id int) (Book, error) {
	resp, err := http.Get(BaseAddress + fmt.Sprintf(BookEndpoint, id))
	if err != nil {
		return Book{}, err
	}
	defer resp.Body.Close()

	var book Book
	err = json.NewDecoder(resp.Body).Decode(&book)
	return book, err
}

func SearchBooks(query string, page int, sort Sort) (Search, error) {
	resp, err := http.Get(BaseAddress + fmt.Sprintf(BookSearchEndpoint, query, page, sort))
	if err != nil {
		return Search{}, err
	}
	defer resp.Body.Close()

	var bookSearch Search
	err = json.NewDecoder(resp.Body).Decode(&bookSearch)
	return bookSearch, err
}
