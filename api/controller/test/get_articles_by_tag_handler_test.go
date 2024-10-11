package test_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nidoqueen1/article-api/api/adapter"
	"github.com/nidoqueen1/article-api/api/controller"
	"github.com/nidoqueen1/article-api/entity"
	"github.com/nidoqueen1/article-api/service/test/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetArticlesByTagHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := logrus.New()

	// Successful case
	t.Run("Success", func(t *testing.T) {
		tagName := "exampleTag"
		relatedTagName := "a related tag"
		dateStr := "2023-01-01"
		date, _ := time.Parse("2006-01-02", dateStr)

		returningArticles := []*entity.Article{
			{
				ID:    1,
				Title: "Sample title 1",
				Body:  "Sample body 1",
				Date:  date,
				Tags: []*entity.Tag{
					{ID: 1, Name: tagName},
					{ID: 2, Name: relatedTagName},
				},
			},
			{
				ID:    2,
				Title: "Sample title 2",
				Body:  "Sample body 2",
				Date:  date,
				Tags: []*entity.Tag{
					{ID: 1, Name: tagName},
				},
			},
		}

		mockService := &mocks.IService{}
		mockService.On("GetArticlesByTagAndDate", tagName, date).Return(returningArticles, int64(len(returningArticles)), nil)

		ginEngine := gin.Default()
		handler := controller.Init(mockService, logger)
		controller.SetupRoutes(ginEngine, handler)

		respRecorder := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/tags/exampleTag/2023-01-01", nil)
		require.Nil(t, err)

		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusOK, respRecorder.Code)

		var responseBody adapter.ArticleListExternalFormat
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)

		require.Equal(t, 2, len(responseBody.Articles))
		require.Equal(t, tagName, responseBody.Tag)
		require.Equal(t, int64(len(returningArticles)), responseBody.Count)
		require.Equal(t, []string{relatedTagName}, responseBody.RelatedTags)
		require.Contains(t, responseBody.Articles, returningArticles[0].ID)
		require.Contains(t, responseBody.Articles, returningArticles[1].ID)
	})

	// Missing Parameters
	t.Run("Missing Parameters", func(t *testing.T) {
		ginEngine := gin.Default()
		handler := controller.Init(&mocks.IService{}, logger)
		controller.SetupRoutes(ginEngine, handler)

		respRecorder := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/tags//someDate", nil) // Missing tagName and date
		require.Nil(t, err)

		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusBadRequest, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Tag name and date are required", responseBody["error"])
	})

	// Invalid Date Format
	t.Run("Invalid Date Format", func(t *testing.T) {
		ginEngine := gin.Default()
		handler := controller.Init(&mocks.IService{}, logger)
		controller.SetupRoutes(ginEngine, handler)

		respRecorder := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/tags/exampleTag/invalid-date", nil)
		require.Nil(t, err)

		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusBadRequest, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Invalid date format", responseBody["error"])
	})

	// Internal Server Error
	t.Run("Internal Server Error", func(t *testing.T) {
		tagName := "exampleTag"
		dateStr := "2023-01-01"
		date, _ := time.Parse("2006-01-02", dateStr)

		mockService := &mocks.IService{}
		mockService.On("GetArticlesByTagAndDate", tagName, date).Return(nil, int64(0), errors.New("database error"))

		ginEngine := gin.Default()
		handler := controller.Init(mockService, logger)
		controller.SetupRoutes(ginEngine, handler)

		respRecorder := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/tags/exampleTag/2023-01-01", nil)
		require.Nil(t, err)

		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusInternalServerError, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Internal error", responseBody["error"])
	})
}
