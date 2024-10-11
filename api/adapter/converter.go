package adapter

import (
	"strings"
	"time"

	"github.com/nidoqueen1/article-api/api/helper"
	"github.com/nidoqueen1/article-api/entity"
)

// ConvertToArticleExternal converts internal DB-format of Article to its external API-format
func ConvertToArticleExternal(article *entity.Article) *ArticleExternalFormat {
	return &ArticleExternalFormat{
		ID:    article.ID,
		Title: article.Title,
		Date:  article.Date.Format("2006-01-02"),
		Body:  article.Body,
		Tags:  getTagNames(article.Tags),
	}
}

// ConvertToArticleListExternalFormat converts a list of Articles to the format proper for GET /tags/{tagName}/{date} API
func ConvertToArticleListExternalFormat(
	articles []*entity.Article, tagName string, count int64) *ArticleListExternalFormat {
	res := &ArticleListExternalFormat{
		Tag:   tagName,
		Count: count,
	}
	res.Articles, res.RelatedTags = extractIDsAndTags(articles, tagName)
	return res
}

// ConvertToInternalArticle onverts external API-format of Article to its internal DB-format
func ConvertToInternalArticle(extArticle *ArticleExternalFormat) (*entity.Article, error) {
	var (
		tags []*entity.Tag
		date time.Time
		err  error
	)

	inArticle := &entity.Article{
		Title: extArticle.Title,
		Body:  extArticle.Body,
	}

	for _, t := range extArticle.Tags {
		tags = append(tags, &entity.Tag{Name: strings.ToLower(t)})
	}

	if extArticle.Date != "" {
		date, err = helper.StringToDate(extArticle.Date)
		if err != nil {
			return nil, err
		}
	}

	inArticle.Date = date
	inArticle.Tags = tags

	return inArticle, nil
}

// extractIDsAndTags extracts Article IDs and related_Tag_names from a list of Articles
func extractIDsAndTags(articles []*entity.Article, tagName string) ([]uint, []string) {
	relatedTags := make(map[string]struct{})
	articleIDs := []uint{}

	for _, a := range articles {
		articleIDs = append(articleIDs, a.ID)
		for _, t := range a.Tags {
			relatedTags[t.Name] = struct{}{}
		}
	}
	delete(relatedTags, tagName)

	return articleIDs, helper.MapToList(relatedTags)
}

// getTagNames extracts Tag names from a list of Tags
func getTagNames(tags []*entity.Tag) []string {
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	return tagNames
}
