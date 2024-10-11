package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/microcosm-cc/bluemonday"
	"github.com/nidoqueen1/article-api/entity"
)

// SanitizeArticle sanitizes the title and body of an article to prevent XSS attacks
func SanitizeArticle(article *entity.Article) {
	// UGCPolicy() allows only user-generated content that is safe
	p := bluemonday.UGCPolicy()

	article.Title = p.Sanitize(article.Title)
	article.Body = p.Sanitize(article.Body)
}

// Validates an Article based on its validate tags
func ValidateArticle(article *entity.Article) error {
	validate := validator.New()

	if err := validate.Struct(article); err != nil {
		return err
	}

	for _, tag := range article.Tags {
		if err := validate.Struct(tag); err != nil {
			return err
		}
	}

	return nil
}
