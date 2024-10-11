package service

import (
	"time"

	"github.com/nidoqueen1/article-api/entity"
)

type IService interface {
	CreateArticle(article *entity.Article) error
	GetArticle(articleID uint) (*entity.Article, error)
	GetArticlesByTagAndDate(tagName string, date time.Time) ([]*entity.Article, int64, error)
}
