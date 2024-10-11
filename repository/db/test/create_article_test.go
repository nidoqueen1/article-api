package test_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nidoqueen1/article-api/entity"
	"github.com/nidoqueen1/article-api/repository/db/postgresql"
	"github.com/nidoqueen1/article-api/repository/db/test/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestCreateArticle(t *testing.T) {
	sqlDB, db, mock := mocks.DbMock(t)
	defer sqlDB.Close()
	logger := logrus.New()
	mockDB := postgresql.New(db, logger)

	article := &entity.Article{
		ID:    1,
		Title: "Test title",
		Body:  "Test body",
		Tags: []*entity.Tag{
			{Name: "tag1"},
			{Name: "tag2"},
		},
	}

	tagRows := sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2)

	mock.ExpectQuery(`INSERT INTO tags`).WillReturnRows(tagRows)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "articles"`).
		WithArgs(article.Title, article.Body, article.Date, article.ID).
		WillReturnRows(sqlmock.NewRows([]string{"LastInsertId", "RowsAffected"}).AddRow(1, 1))
	mock.ExpectCommit()

	mock.ExpectExec(`INSERT INTO article_tags`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := mockDB.CreateArticle(article)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
