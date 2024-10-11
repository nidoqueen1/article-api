package test_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nidoqueen1/article-api/entity"
	"github.com/nidoqueen1/article-api/repository/db/test/mocks"
	"github.com/nidoqueen1/article-api/service"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetArticlesByTagAndDate_Success(t *testing.T) {
	mockDB := new(mocks.IDatabase)
	logger := logrus.New()
	svc := service.Init(mockDB, logger)

	tagName := "health"
	date := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	expectedArticles := []*entity.Article{
		{ID: 1, Title: "Article 1", Body: "Body of article 1"},
		{ID: 2, Title: "Article 2", Body: "Body of article 2"},
	}
	expectedCount := int64(len(expectedArticles))

	mockDB.On("GetArticlesByTagAndDate", context.TODO(), tagName, date).Return(expectedArticles, expectedCount, nil)
	articles, count, err := svc.GetArticlesByTagAndDate(context.TODO(), tagName, date)

	assert.NoError(t, err)
	assert.NotNil(t, articles)
	assert.Equal(t, expectedArticles, articles)
	assert.Equal(t, expectedCount, count)
	mockDB.AssertExpectations(t)
}

func TestGetArticlesByTagAndDate_ErrorFetching(t *testing.T) {
	mockDB := new(mocks.IDatabase)
	logger := logrus.New()
	svc := service.Init(mockDB, logger)

	tagName := "fitness"
	date := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	expectedErr := errors.New("database error")

	mockDB.On("GetArticlesByTagAndDate", context.TODO(), tagName, date).Return(nil, int64(0), expectedErr)
	articles, count, err := svc.GetArticlesByTagAndDate(context.TODO(), tagName, date)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, articles)
	assert.Equal(t, int64(0), count)
	mockDB.AssertExpectations(t)
}

func TestGetArticlesByTagAndDate_NoArticlesFound(t *testing.T) {
	mockDB := new(mocks.IDatabase)
	logger := logrus.New()
	svc := service.Init(mockDB, logger)

	tagName := "science"
	date := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)

	// simulate no articles found
	mockDB.On("GetArticlesByTagAndDate", context.TODO(), tagName, date).Return([]*entity.Article{}, int64(0), nil)
	articles, count, err := svc.GetArticlesByTagAndDate(context.TODO(), tagName, date)

	assert.NoError(t, err)
	assert.NotNil(t, articles)
	assert.Equal(t, []*entity.Article{}, articles) // empty slice should be returned
	assert.Equal(t, int64(0), count)
	mockDB.AssertExpectations(t)
}

func TestGetArticlesByTagAndDate_InvalidDate(t *testing.T) {
	mockDB := new(mocks.IDatabase)
	logger := logrus.New()
	svc := service.Init(mockDB, logger)

	tagName := "technology"
	invalidDate := time.Time{} // an invalid zero date

	// simulate no articles found for invalid date
	mockDB.On("GetArticlesByTagAndDate", context.TODO(), tagName, invalidDate).Return([]*entity.Article{}, int64(0), nil)
	articles, count, err := svc.GetArticlesByTagAndDate(context.TODO(), tagName, invalidDate)

	assert.NoError(t, err)
	assert.NotNil(t, articles)
	assert.Equal(t, []*entity.Article{}, articles) // empty slice should be returned
	assert.Equal(t, int64(0), count)
	mockDB.AssertExpectations(t)
}
