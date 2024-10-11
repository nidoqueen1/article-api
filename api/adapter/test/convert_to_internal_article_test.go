package test_test

import (
	"testing"
	"time"

	"github.com/nidoqueen1/article-api/api/adapter"
	"github.com/nidoqueen1/article-api/entity"
	"github.com/stretchr/testify/require"
)

func TestConvertToInternalArticle(t *testing.T) {
	t.Run("Conversion with tags and valid date", func(t *testing.T) {
		extArticle := &adapter.ArticleExternalFormat{
			Title: "Test article",
			Body:  "Test body",
			Date:  "2023-10-11",
			Tags:  []string{"Go", "Gin"},
		}

		expected := &entity.Article{
			Title: "Test article",
			Body:  "Test body",
			Date:  time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
			Tags: []*entity.Tag{
				{Name: "go"},
				{Name: "gin"},
			},
		}

		result, err := adapter.ConvertToInternalArticle(extArticle)

		require.NoError(t, err)
		require.Equal(t, expected.Title, result.Title)
		require.Equal(t, expected.Body, result.Body)
		require.Equal(t, expected.Date, result.Date)
		require.Len(t, result.Tags, len(expected.Tags))
		for i, tag := range result.Tags {
			require.Equal(t, expected.Tags[i].Name, tag.Name)
		}
	})

	t.Run("Conversion with no tags", func(t *testing.T) {
		extArticle := &adapter.ArticleExternalFormat{
			Title: "Test article",
			Body:  "Test body",
			Date:  "2023-10-11",
			Tags:  []string{},
		}

		expected := &entity.Article{
			Title: "Test article",
			Body:  "Test body",
			Date:  time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
			Tags:  []*entity.Tag{},
		}

		result, err := adapter.ConvertToInternalArticle(extArticle)

		require.NoError(t, err)
		require.Equal(t, expected.Title, result.Title)
		require.Equal(t, expected.Body, result.Body)
		require.Equal(t, expected.Date, result.Date)
		require.Len(t, result.Tags, 0)
	})

	t.Run("Invalid date format", func(t *testing.T) {
		extArticle := &adapter.ArticleExternalFormat{
			Title: "Test article",
			Body:  "Test body",
			Date:  "invalid date",
			Tags:  []string{"Go", "Gin"},
		}

		result, err := adapter.ConvertToInternalArticle(extArticle)

		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("Conversion with missing date", func(t *testing.T) {
		extArticle := &adapter.ArticleExternalFormat{
			Title: "Test article",
			Body:  "Test body",
			Date:  "",
			Tags:  []string{"Go", "Gin"},
		}

		expected := &entity.Article{
			Title: "Test article",
			Body:  "Test body",
			Tags: []*entity.Tag{
				{Name: "go"},
				{Name: "gin"},
			},
		}

		result, err := adapter.ConvertToInternalArticle(extArticle)

		require.NoError(t, err)
		require.Equal(t, expected.Title, result.Title)
		require.Equal(t, expected.Body, result.Body)
		require.Len(t, result.Tags, len(expected.Tags))
		for i, tag := range result.Tags {
			require.Equal(t, expected.Tags[i].Name, tag.Name)
		}
		require.True(t, result.Date.IsZero(), "expected zero date when date is not provided")
	})
}
