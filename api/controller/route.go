package controller

import "github.com/gin-gonic/gin"

// Setups endpoints with their handles
func SetupRoutes(router *gin.Engine, h *handler) {
	router.POST("/articles", h.CreateArticleHandler)
	router.GET("/articles/:id", h.GetArticleHandler)
	router.GET("/tags/:tagName/:date", h.GetArticlesByTagHandler)
}
