package service

import (
	"github.com/nidoqueen1/article-api/repository/db"
	"github.com/sirupsen/logrus"
)

type service struct {
	db     db.IDatabase
	logger *logrus.Logger
}

// Initialized service instance
func Init(db db.IDatabase, l *logrus.Logger) IService {
	return &service{
		db:     db,
		logger: l,
	}
}
