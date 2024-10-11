package postgresql

import (
	"fmt"

	"github.com/nidoqueen1/article-api/entity"
	"github.com/nidoqueen1/article-api/repository/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresql struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func Init(logger *logrus.Logger) (db.IDatabase, error) {
	db, err := gorm.Open(postgres.Open(viper.GetString("db.url")), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect to database: ", err)
		return nil, fmt.Errorf("failed to connect to database: %s", err)
	}

	// configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error("failed to get sql DB: ", err)
		return nil, fmt.Errorf("failed to get sql DB: %s", err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * 60)

	err = db.AutoMigrate(&entity.Article{}, &entity.Tag{}, &entity.ArticleTag{})
	if err != nil {
		logger.Error("failed to migrate schema: ", err)
		return nil, fmt.Errorf("failed to migrate schema: %s", err)
	}

	logger.Info("Connected to the PostgreSQL database.")
	return &postgresql{
		db:     db,
		logger: logger,
	}, nil
}
