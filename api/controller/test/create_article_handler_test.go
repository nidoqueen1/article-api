package test_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nidoqueen1/article-api/api/adapter"
	"github.com/nidoqueen1/article-api/api/controller"
	"github.com/nidoqueen1/article-api/service/test/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateArticleHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := logrus.New()

	// Successful case
	t.Run("Success", func(t *testing.T) {
		mockService := &mocks.IService{}
		mockService.On("CreateArticle", context.TODO(), mock.Anything).Return(nil)

		ginEngine := gin.Default()
		handler := controller.Init(mockService, logger)
		controller.SetupRoutes(ginEngine, handler)

		articleExtFormat := adapter.ArticleExternalFormat{
			Title: "Sample Title",
			Body:  "Sample Body",
			Date:  "2023-01-01",
			Tags:  []string{"Tag1", "Tag2"},
		}
		body, err := json.Marshal(articleExtFormat)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		respRecorder := httptest.NewRecorder()
		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusCreated, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "success", responseBody["status"])
	})

	// Invalid JSON payload
	t.Run("Invalid JSON", func(t *testing.T) {
		ginEngine := gin.Default()
		handler := controller.Init(&mocks.IService{}, logger)
		controller.SetupRoutes(ginEngine, handler)

		req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer([]byte("invalid-json")))
		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		respRecorder := httptest.NewRecorder()
		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusBadRequest, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Invalid request", responseBody["error"])
	})

	// Not convertable payload
	t.Run("Not Convertable Payload", func(t *testing.T) {
		ginEngine := gin.Default()
		handler := controller.Init(&mocks.IService{}, logger)
		controller.SetupRoutes(ginEngine, handler)

		articleExtFormat := adapter.ArticleExternalFormat{
			Title: "Sample Title",
			Body:  "Sample Body",
			Date:  "invalid-date",
			Tags:  []string{"Tag1", "Tag2"},
		}
		body, err := json.Marshal(articleExtFormat)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		respRecorder := httptest.NewRecorder()
		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusBadRequest, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Not convertable payload", responseBody["error"])
	})

	// Validation failure
	t.Run("Validation Failure", func(t *testing.T) {
		ginEngine := gin.Default()
		handler := controller.Init(&mocks.IService{}, logger)
		controller.SetupRoutes(ginEngine, handler)

		articleExtFormat := adapter.ArticleExternalFormat{
			Title: "",
			Body:  "Sample Body",
			Date:  "2023-01-01",
			Tags:  []string{"Tag1", "Tag2"},
		}
		body, err := json.Marshal(articleExtFormat)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		respRecorder := httptest.NewRecorder()
		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusBadRequest, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Not allowed payload", responseBody["error"])
	})

	// Internal server error
	t.Run("Internal Server Error", func(t *testing.T) {
		mockService := &mocks.IService{}
		mockService.On("CreateArticle", context.TODO(), mock.Anything).Return(errors.New("database error"))

		ginEngine := gin.Default()
		handler := controller.Init(mockService, logger)
		controller.SetupRoutes(ginEngine, handler)

		articleExtFormat := adapter.ArticleExternalFormat{
			Title: "Sample Title",
			Body:  "Sample Body",
			Date:  "2023-01-01",
			Tags:  []string{"Tag1", "Tag2"},
		}
		body, err := json.Marshal(articleExtFormat)
		require.Nil(t, err)

		req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
		require.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")

		respRecorder := httptest.NewRecorder()
		ginEngine.ServeHTTP(respRecorder, req)

		require.Equal(t, http.StatusInternalServerError, respRecorder.Code)

		var responseBody map[string]interface{}
		err = json.Unmarshal(respRecorder.Body.Bytes(), &responseBody)
		require.Nil(t, err)
		require.Equal(t, "Failed to save article", responseBody["error"])
	})
}
