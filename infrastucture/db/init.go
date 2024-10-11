package db

import (
	"fmt"

	"github.com/nidoqueen1/article-api/repository/db"
	"github.com/nidoqueen1/article-api/repository/db/postgresql"
	"github.com/sirupsen/logrus"
)

func InitDatabase(dbType string, l *logrus.Logger) (db.IDatabase, error) {
	switch dbType {
	case "postgresql":
		return postgresql.Init(l)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
