package test_test

import (
	"sort"
	"testing"
	"time"

	"github.com/nidoqueen1/article-api/api/adapter"
	"github.com/nidoqueen1/article-api/entity"
	"github.com/stretchr/testify/require"
)

func TestConvertToArticleExternal(t *testing.T) {
	t.Run("Valid conversion of article with tags", func(t *testing.T) {
		article := &entity.Article{
			ID:    1,
			Title: "Sample Article",
			Body:  "This is a sample article.",
			Date:  time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
			Tags: []*entity.Tag{
				{Name: "go"},
				{Name: "testing"},
			},
		}

		expected := &adapter.ArticleExternalFormat{
			ID:    1,
			Title: "Sample Article",
			Date:  "2023-10-11",
			Body:  "This is a sample article.",
			Tags:  []string{"go", "testing"},
		}

		result := adapter.ConvertToArticleExternal(article)

		// sort slices' elements to prevent mismatching in require.Equal
		sort.Strings(expected.Tags)
		sort.Strings(result.Tags)
		
		require.Equal(t, expected, result)
	})

	t.Run("Valid conversion of article without tags", func(t *testing.T) {
		article := &entity.Article{
			ID:    2,
			Title: "Article without tags",
			Body:  "This article has no tags.",
			Date:  time.Date(2023, 10, 12, 0, 0, 0, 0, time.UTC),
			Tags:  []*entity.Tag{},
		}

		expected := &adapter.ArticleExternalFormat{
			ID:    2,
			Title: "Article without tags",
			Date:  "2023-10-12",
			Body:  "This article has no tags.",
			Tags:  []string(nil),
		}

		result := adapter.ConvertToArticleExternal(article)
		require.Equal(t, expected, result)
	})
}
