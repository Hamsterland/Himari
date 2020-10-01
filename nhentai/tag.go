package nhentai

type Tag struct {
	Id    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Url   string `json:"url"`
	Count int    `json:"count"`
}
