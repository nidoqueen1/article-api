package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nidoqueen1/article-api/api/adapter"
	"github.com/nidoqueen1/article-api/api/helper"
)

// Handler for endpoint POST /articles
func (h *handler) CreateArticleHandler(c *gin.Context) {
	var articleExtFormat adapter.ArticleExternalFormat
	if err := c.ShouldBindJSON(&articleExtFormat); err != nil {
		h.logger.Errorf("Invalid request payload: %v, error: %s", c.Request.Body, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	article, err := adapter.ConvertToInternalArticle(&articleExtFormat)
	if err != nil {
		h.logger.Errorf("Not convertable payload: %v, error: %s", articleExtFormat, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not convertable payload"})
		return
	}

	if err := helper.ValidateArticle(article); err != nil {
		h.logger.Errorf("Not allowed request payload: %v, error: %s", article, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not allowed payload"})
		return
	}

	ctx := c.Request.Context()
	if err := h.service.CreateArticle(ctx, article); err != nil {
		h.logger.Errorf("Failed to save article: %v, error: %s", article, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save article"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

// Handler for endpoint GET /articles/{id}
func (h *handler) GetArticleHandler(c *gin.Context) {
	id := c.Param("id")
	uintID, err := helper.StringToUint(id)
	if err != nil {
		h.logger.Errorf("Invalid ID: %s, error: %s", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ctx := c.Request.Context()
	article, err := h.service.GetArticle(ctx, uintID)
	if err != nil {
		h.logger.Errorf("Article not found, internal error: %s, ID: %s", err, id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}
	if article == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Sanitize the article's content
	helper.SanitizeArticle(article)

	ArticleExternalFormat := adapter.ConvertToArticleExternal(article)
	c.JSON(http.StatusOK, ArticleExternalFormat)
}

// Handler for endpoint GET /tags/{tagName}/{date}
func (h *handler) GetArticlesByTagHandler(c *gin.Context) {
	tagName := c.Param("tagName")
	dateStr := c.Param("date")

	if tagName == "" || dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag name and date are required"})
		return
	}

	date, err := helper.StringToDate(dateStr)
	if err != nil {
		h.logger.Errorf("Invalid date format %s, error: %s", dateStr, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	ctx := c.Request.Context()
	articles, count, err := h.service.GetArticlesByTagAndDate(ctx, tagName, date)
	if err != nil {
		h.logger.Errorf("Articles not found, internal error: %s, tag name: %s, date: %s",
			err, tagName, dateStr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	response := adapter.ConvertToArticleListExternalFormat(articles, tagName, count)
	c.JSON(http.StatusOK, response)
}
