package test_test

import (
	"testing"
	"time"

	"github.com/nidoqueen1/article-api/entity"
	"github.com/nidoqueen1/article-api/repository/db/test/mocks"
	"github.com/nidoqueen1/article-api/service"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// TestCreateArticle tests the CreateArticle function
func TestCreateArticle(t *testing.T) {
	mockDB := new(mocks.IDatabase)
	logger := logrus.New()
	svc := service.Init(mockDB, logger)

	tests := []struct {
		name           string
		inputArticle   *entity.Article
		mockCreateResp error
		expectedError  error
	}{
		{
			name: "Create article with zero date, should set to today",
			inputArticle: &entity.Article{
				Title: "New Title",
				Body:  "Some body content",
				Tags:  nil,
			},
			mockCreateResp: nil,
			expectedError:  nil,
		},
		{
			name: "Create article with provided date",
			inputArticle: &entity.Article{
				Title: "Another Title",
				Body:  "Some body content",
				Date:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				Tags:  nil,
			},
			mockCreateResp: nil,
			expectedError:  nil,
		},
		{
			name: "Error creating article",
			inputArticle: &entity.Article{
				Title: "Title That Will Fail",
				Body:  "This will fail",
			},
			mockCreateResp: assert.AnError,
			expectedError:  assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB.On("CreateArticle", tt.inputArticle).Return(tt.mockCreateResp).Once()
			err := svc.CreateArticle(tt.inputArticle)

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockDB.AssertExpectations(t)
		})
	}
}
