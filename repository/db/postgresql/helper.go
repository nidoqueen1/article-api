package postgresql

import "github.com/nidoqueen1/article-api/entity"

// Converts Tag names of an Article to a list
func getTagNames(article entity.Article) []string {
	tagNames := []string{}
	for _, tag := range article.Tags {
		tagNames = append(tagNames, tag.Name)
	}
	return tagNames
}
