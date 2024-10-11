package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nidoqueen1/article-api/api/controller"
	"github.com/nidoqueen1/article-api/infrastucture/config"
	"github.com/nidoqueen1/article-api/infrastucture/db"
	"github.com/nidoqueen1/article-api/infrastucture/logging"
	"github.com/nidoqueen1/article-api/service"
	"github.com/spf13/viper"
)

// Initializes the application's elements and creates a Router to serve the endpoints
func New() *gin.Engine {
	logger := logging.InitLog()

	err := config.InitConfig()
	if err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := db.InitDatabase(viper.GetString("db.type"), logger)
	if err != nil {
		logger.Fatalf("Failed to init database: %v", err)
	}

	service := service.Init(db, logger)
	handler := controller.Init(service, logger)
	router := gin.Default()
	controller.SetupRoutes(router, handler)
	logger.Infof("Listening on %s . . .", viper.GetString("server.addr"))

	return router
}
