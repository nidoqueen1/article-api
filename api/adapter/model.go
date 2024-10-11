package adapter

// Custom struct to format the response with a properly formatted date
type ArticleExternalFormat struct {
	ID    uint     `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

type articleListExternalFormat struct {
	Tag         string   `json:"tag"`
	Count       int64    `json:"count"`
	Articles    []uint   `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}
