package test_test

import (
	"testing"

	"github.com/nidoqueen1/article-api/api/helper"
	"github.com/nidoqueen1/article-api/entity"
	"github.com/stretchr/testify/assert"
)

func TestSanitizeArticle(t *testing.T) {
	t.Run("Sanitizes Article Title and Body", func(t *testing.T) {
		article := &entity.Article{
			Title: "<script>alert('Title XSS')</script>Sample Title",
			Body:  "A<tag>B'C'D",
		}

		helper.SanitizeArticle(article)

		assert.Equal(t, "Sample Title", article.Title)
		assert.Equal(t, "AB&#39;C&#39;D", article.Body)
	})

	t.Run("No Sanitization Needed", func(t *testing.T) {
		article := &entity.Article{
			Title: "Clean Title",
			Body:  "Clean Body",
		}

		helper.SanitizeArticle(article)

		assert.Equal(t, "Clean Title", article.Title)
		assert.Equal(t, "Clean Body", article.Body)
	})

	t.Run("Empty Fields", func(t *testing.T) {
		article := &entity.Article{
			Title: "",
			Body:  "",
		}

		helper.SanitizeArticle(article)

		assert.Equal(t, "", article.Title)
		assert.Equal(t, "", article.Body)
	})
}
