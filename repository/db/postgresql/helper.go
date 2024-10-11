package postgresql

import "github.com/nidoqueen1/article-api/entity"

func getTagNames(article entity.Article) []string {
	tagNames := []string{}
	for _, tag := range article.Tags {
		tagNames = append(tagNames, tag.Name)
	}
	return tagNames
}
