package test_test

import (
	"testing"
	"time"

	"github.com/nidoqueen1/article-api/api/helper"
	"github.com/nidoqueen1/article-api/entity"
	"github.com/stretchr/testify/assert"
)

func TestValidateArticle(t *testing.T) {
	t.Run("Valid Article", func(t *testing.T) {
		article := &entity.Article{
			Title: "Valid Title",
			Body:  "This is a valid body for the article.",
			Date:  time.Now(),
			Tags: []*entity.Tag{
				{
					Name: "ValidTag",
				},
			},
		}

		err := helper.ValidateArticle(article)
		assert.Nil(t, err)
	})

	t.Run("Missing Title", func(t *testing.T) {
		article := &entity.Article{
			Title: "",
			Body:  "This body is fine.",
			Date:  time.Now(),
			Tags: []*entity.Tag{
				{
					Name: "Tag",
				},
			},
		}

		err := helper.ValidateArticle(article)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Title")
	})

	t.Run("Missing Body", func(t *testing.T) {
		article := &entity.Article{
			Title: "Valid Title",
			Body:  "",
			Date:  time.Now(),
			Tags: []*entity.Tag{
				{
					Name: "Tag",
				},
			},
		}

		err := helper.ValidateArticle(article)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Body")
	})

	t.Run("Tag Name Too Short", func(t *testing.T) {
		article := &entity.Article{
			Title: "Valid Title",
			Body:  "Valid body",
			Date:  time.Now(),
			Tags: []*entity.Tag{
				{
					Name: "",
				},
			},
		}

		err := helper.ValidateArticle(article)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Name")
	})

	t.Run("Tag Name Contains Non-Alphanumeric Characters", func(t *testing.T) {
		article := &entity.Article{
			Title: "Valid Title",
			Body:  "Valid body",
			Date:  time.Now(),
			Tags: []*entity.Tag{
				{
					Name: "Invalid Tag<@>!", // non-alphanumeric characters
				},
			},
		}

		err := helper.ValidateArticle(article)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "Name")
	})

	t.Run("Valid Article with Multiple Tags", func(t *testing.T) {
		article := &entity.Article{
			Title: "Valid Title",
			Body:  "Valid body",
			Date:  time.Now(),
			Tags: []*entity.Tag{
				{
					Name: "Tag1",
				},
				{
					Name: "Tag2",
				},
			},
		}

		err := helper.ValidateArticle(article)
		assert.Nil(t, err)
	})

	t.Run("Article With Missing Tags", func(t *testing.T) {
		article := &entity.Article{
			Title: "Valid Title",
			Body:  "Valid body",
			Date:  time.Now(),
			Tags:  []*entity.Tag{},
		}

		err := helper.ValidateArticle(article)
		assert.Nil(t, err) // no error for empty tags
	})
}
