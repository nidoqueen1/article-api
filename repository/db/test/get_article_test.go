package test_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nidoqueen1/article-api/entity"
	"github.com/nidoqueen1/article-api/repository/db/postgresql"
	"github.com/nidoqueen1/article-api/repository/db/test/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetArticle_success(t *testing.T) {
	sqlDB, db, mock := mocks.DbMock(t)
	defer sqlDB.Close()
	logger := logrus.New()
	mockDB := postgresql.New(db, logger)

	articleID := uint(1)
	article := entity.Article{
		ID:    articleID,
		Title: "Test title",
		Body:  "Test body.",
		Tags: []*entity.Tag{
			{ID: 1, Name: "tag1"},
			{ID: 2, Name: "tag2"},
		},
		Date: time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
	}

	rows := sqlmock.NewRows([]string{"id", "title", "boddy", "date"}).
		AddRow(article.ID, article.Title, article.Body, article.Date)
	mock.ExpectQuery(`SELECT \* FROM "articles" WHERE "articles"."id" = \$1 ORDER BY "articles"."id" LIMIT \$2`).
		WithArgs(articleID, 1).
		WillReturnRows(rows)

	// mock preloading tags
	articleTagsRows := sqlmock.NewRows([]string{"article_id", "tag_id"}).
		AddRow(articleID, article.Tags[0].ID).
		AddRow(articleID, article.Tags[1].ID)
	mock.ExpectQuery(`SELECT \* FROM "article_tags" WHERE "article_tags"."article_id" = \$1`).
		WithArgs(articleID).
		WillReturnRows(articleTagsRows)

	tagRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(article.Tags[0].ID, article.Tags[0].Name).
		AddRow(article.Tags[1].ID, article.Tags[1].Name)
	mock.ExpectQuery(`SELECT \* FROM "tags" WHERE "tags"."id" IN \(\$1,\$2\)`).
		WithArgs(article.Tags[0].ID, article.Tags[1].ID).
		WillReturnRows(tagRows)

	result, err := mockDB.GetArticle(articleID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, article.ID, result.ID)
	require.Equal(t, article.Title, result.Title)
	require.Equal(t, len(article.Tags), len(result.Tags))
}

func TestGetArticle_NotFound(t *testing.T) {
	sqlDB, db, mock := mocks.DbMock(t)
	defer sqlDB.Close()
	logger := logrus.New()
	mockDB := postgresql.New(db, logger)

	articleID := uint(999999)

	mock.ExpectQuery(`SELECT \* FROM "articles" WHERE "articles"."id" = \$1 ORDER BY "articles"."id" LIMIT \$2`).
		WithArgs(articleID, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	result, err := mockDB.GetArticle(articleID)
	require.NoError(t, err)
	require.Nil(t, result)
}
