package test_test

import (
	"sort"
	"testing"
	"time"

	"github.com/nidoqueen1/article-api/api/adapter"
	"github.com/nidoqueen1/article-api/entity"
	"github.com/stretchr/testify/require"
)

func TestConvertToArticleListExternalFormat(t *testing.T) {
	t.Run("Articles with related tags", func(t *testing.T) {
		articles := []*entity.Article{
			{
				ID:    1,
				Title: "First Article",
				Body:  "Body of first article",
				Date:  time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
				Tags: []*entity.Tag{
					{Name: "go"},
					{Name: "testing"},
				},
			},
			{
				ID:    2,
				Title: "Second Article",
				Body:  "Body of second article",
				Date:  time.Date(2023, 10, 12, 0, 0, 0, 0, time.UTC),
				Tags: []*entity.Tag{
					{Name: "gin"},
					{Name: "testing"},
				},
			},
		}

		tagName := "testing"
		count := int64(2)

		expected := &adapter.ArticleListExternalFormat{
			Tag:         "testing",
			Count:       2,
			Articles:    []uint{1, 2},
			RelatedTags: []string{"go", "gin"},
		}

		result := adapter.ConvertToArticleListExternalFormat(articles, tagName, count)

		// sort slices' elements to prevent mismatching in require.Equal
		sort.Strings(expected.RelatedTags)
		sort.Strings(result.RelatedTags)
		var intIDs []int
		for _, id := range result.Articles {
			intIDs = append(intIDs, int(id))
		}
		sort.Ints(intIDs)
		for i, id := range intIDs {
			result.Articles[i] = uint(id)
		}

		require.Equal(t, expected, result)
	})

	t.Run("Conversion with no related tags", func(t *testing.T) {
		articles := []*entity.Article{
			{
				ID:    1,
				Title: "First Article",
				Body:  "Body of first article",
				Date:  time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
				Tags: []*entity.Tag{
					{Name: "go"},
				},
			},
		}

		tagName := "go"
		count := int64(1)

		expected := &adapter.ArticleListExternalFormat{
			Tag:         "go",
			Count:       1,
			Articles:    []uint{1},
			RelatedTags: []string{},
		}

		result := adapter.ConvertToArticleListExternalFormat(articles, tagName, count)
		require.Equal(t, expected, result)
	})

	t.Run("Conversion with no articles", func(t *testing.T) {
		articles := []*entity.Article{}
		tagName := "go"
		count := int64(0)

		expected := &adapter.ArticleListExternalFormat{
			Tag:         "go",
			Count:       0,
			Articles:    []uint{},
			RelatedTags: []string{},
		}

		result := adapter.ConvertToArticleListExternalFormat(articles, tagName, count)
		require.Equal(t, expected, result)
	})
}
