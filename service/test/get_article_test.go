package test_test

import (
	"context"
	"errors"
	"testing"

	"github.com/nidoqueen1/article-api/entity"
	"github.com/nidoqueen1/article-api/repository/db/test/mocks"
	"github.com/nidoqueen1/article-api/service"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetArticle_Success(t *testing.T) {
	mockDB := new(mocks.IDatabase)
	logger := logrus.New()
	svc := service.Init(mockDB, logger)

	articleID := uint(1)
	expectedArticle := &entity.Article{
		ID:    articleID,
		Title: "Sample Article",
		Body:  "This is a sample article.",
	}

	mockDB.On("GetArticle", context.TODO(), articleID).Return(expectedArticle, nil)
	result, err := svc.GetArticle(context.TODO(), articleID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedArticle, result)
	mockDB.AssertExpectations(t)
}

func TestGetArticle_ErrorFetching(t *testing.T) {
	mockDB := new(mocks.IDatabase)
	logger := logrus.New()
	svc := service.Init(mockDB, logger)

	articleID := uint(2)
	expectedErr := errors.New("database error")

	mockDB.On("GetArticle", context.TODO(), articleID).Return(nil, expectedErr)
	result, err := svc.GetArticle(context.TODO(), articleID)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, result)
	mockDB.AssertExpectations(t)
}

func TestGetArticle_NoArticleFound(t *testing.T) {
	mockDB := new(mocks.IDatabase)
	logger := logrus.New()
	svc := service.Init(mockDB, logger)

	articleID := uint(3)

	mockDB.On("GetArticle", context.TODO(), articleID).Return(nil, nil) // Simulate no error but also no article found
	result, err := svc.GetArticle(context.TODO(), articleID)

	assert.NoError(t, err)
	assert.Nil(t, result)
	mockDB.AssertExpectations(t)
}
