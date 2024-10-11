package db

import (
	"time"

	"github.com/nidoqueen1/article-api/entity"
)

// Interface type of general database related methods
type IDatabase interface {
	CreateArticle(article *entity.Article) error
	GetArticle(articleID uint) (*entity.Article, error)
	GetArticlesByTagAndDate(tagName string, date time.Time) ([]*entity.Article, int64, error)
}
