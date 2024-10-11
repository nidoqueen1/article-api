package service

import (
	"time"

	"github.com/nidoqueen1/article-api/entity"
)

// Interface type of service related methods
type IService interface {
	CreateArticle(article *entity.Article) error
	GetArticle(articleID uint) (*entity.Article, error)
	GetArticlesByTagAndDate(tagName string, date time.Time) ([]*entity.Article, int64, error)
}
