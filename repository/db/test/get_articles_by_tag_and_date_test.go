package test_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nidoqueen1/article-api/repository/db/postgresql"
	"github.com/nidoqueen1/article-api/repository/db/test/mocks"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestGetArticlesByTagAndDate_success(t *testing.T) {
	sqlDB, db, mock := mocks.DbMock(t)
	defer sqlDB.Close()
	logger := logrus.New()
	mockDB := postgresql.New(db, logger)

	viper.Set("list_articles.limit", 2)

	tagName := "test_tag"
	date := time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC)
	articleID := uint(1)

	mock.ExpectQuery(`SELECT count\(\*\) FROM "articles"`).
		WithArgs(tagName, date).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	articleRows := sqlmock.NewRows([]string{"id", "title", "body", "date"}).
		AddRow(articleID, "Title article 1", "Body article 1", date).
		AddRow(articleID+1, "Title article 2", "Body article 2", date)
	mock.ExpectQuery(`
	SELECT "articles"."id","articles"."title","articles"."body","articles"."date" FROM "articles"`).
		WithArgs(tagName, date, 2).
		WillReturnRows(articleRows)

	// mock preloading tags
	articleTagsRows := sqlmock.NewRows([]string{"article_id", "tag_id"}).
		AddRow(articleID, 1).
		AddRow(articleID, 2).
		AddRow(articleID, 3)
	mock.ExpectQuery(`SELECT \* FROM "article_tags" WHERE "article_tags"."article_id" IN \(\$1,\$2\)`).
		WithArgs(articleID, articleID+1).
		WillReturnRows(articleTagsRows)

	tagRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "tag1").
		AddRow(2, "tag2").
		AddRow(3, "tag3")
	mock.ExpectQuery(`SELECT \* FROM "tags" WHERE "tags"."id" IN \(\$1,\$2,\$3\)`).
		WithArgs(1, 2, 3).
		WillReturnRows(tagRows)

	articles, totalCount, err := mockDB.GetArticlesByTagAndDate(tagName, date)
	require.NoError(t, err)
	require.Len(t, articles, 2)
	require.Equal(t, int64(2), totalCount)
}

func TestGetArticlesByTagAndDate_NotFound(t *testing.T) {
	sqlDB, db, mock := mocks.DbMock(t)
	defer sqlDB.Close()
	logger := logrus.New()
	mockDB := postgresql.New(db, logger)

	viper.Set("list_articles.limit", 2)

	tagName := "non_existent_tag"
	date := time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC)

	mock.ExpectQuery(`SELECT count\(\*\) FROM "articles"`).
		WithArgs(tagName, date).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	mock.ExpectQuery(`
	SELECT "articles"."id","articles"."title","articles"."body","articles"."date" FROM "articles"`).
		WithArgs(tagName, date, 2).
		WillReturnRows(sqlmock.NewRows([]string{}))

	articles, totalCount, err := mockDB.GetArticlesByTagAndDate(tagName, date)
	require.NoError(t, err)
	require.Empty(t, articles)
	require.Equal(t, int64(0), totalCount)
}
