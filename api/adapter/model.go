package adapter

// Extarnal API-format of the Article
type ArticleExternalFormat struct {
	ID    uint     `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

// Extarnal API-format of a list of Articles
// used in endpoint GET /tags/{tagName}/{date}
type ArticleListExternalFormat struct {
	Tag         string   `json:"tag"`
	Count       int64    `json:"count"`
	Articles    []uint   `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}
