package test_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nidoqueen1/article-api/api/controller"
	"github.com/nidoqueen1/article-api/entity"
	"github.com/nidoqueen1/article-api/service/test/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetArticleHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := logrus.New()

	// Success case
	t.Run("Success", func(t *testing.T) {
		returningArticle := &entity.Article{
			ID:    1,
			Title: "Sample title",
			Body:  "Sample body",
			Date:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
			Tags: []*entity.Tag{
				{ID: 1, Name: "Tag name"},
				{ID: 2, Name: "Another Tag name"},
			},
		}

		mockService := &mocks.IService{}
		mockService.On("GetArticle", context.TODO(), mock.AnythingOfType("uint")).Return(returningArticle, nil)

		ginEngine := gin.Default()
		handler := controller.Init(mockService, logger)
		controller.SetupRoutes(ginEngine, handler)

		respRecorder := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/articles/1", nil)
		require.Nil(t, err)

		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusOK, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)

		require.Equal(t, returningArticle.ID, uint(responseBody["id"].(float64)))
		require.Equal(t, returningArticle.Title, responseBody["title"].(string))
		require.Equal(t, returningArticle.Body, responseBody["body"].(string))
		require.Equal(t, returningArticle.Date.Format("2006-01-02"), responseBody["date"].(string))

		expectedTags := []string{"Tag name", "Another Tag name"}
		var responseTags []string
		for _, tag := range responseBody["tags"].([]interface{}) {
			responseTags = append(responseTags, tag.(string))
		}
		require.Equal(t, expectedTags, responseTags)
	})

	// Invalid ID Format
	t.Run("Invalid ID Format", func(t *testing.T) {
		mockService := &mocks.IService{}
		ginEngine := gin.Default()
		handler := controller.Init(mockService, logger)
		controller.SetupRoutes(ginEngine, handler)

		respRecorder := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/articles/invalid-id", nil)
		require.Nil(t, err)

		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusBadRequest, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Invalid ID", responseBody["error"])
	})

	// Article Not Found
	t.Run("Article Not Found", func(t *testing.T) {
		mockService := &mocks.IService{}
		mockService.On("GetArticle", context.TODO(), mock.AnythingOfType("uint")).Return(nil, nil) // simulating article not found

		ginEngine := gin.Default()
		handler := controller.Init(mockService, logger)
		controller.SetupRoutes(ginEngine, handler)

		respRecorder := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/articles/99999999", nil) // non-existing ID
		require.Nil(t, err)

		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusNotFound, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Article not found", responseBody["error"])
	})

	// Internal Server Error
	t.Run("Internal Server Error", func(t *testing.T) {
		mockService := &mocks.IService{}
		mockService.On("GetArticle", context.TODO(), mock.AnythingOfType("uint")).Return(nil, errors.New("database error"))

		ginEngine := gin.Default()
		handler := controller.Init(mockService, logger)
		controller.SetupRoutes(ginEngine, handler)

		respRecorder := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/articles/1", nil)
		require.Nil(t, err)

		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusInternalServerError, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Internal error", responseBody["error"])
	})
}
