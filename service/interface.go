package service

import (
	"context"
	"time"

	"github.com/nidoqueen1/article-api/entity"
)

// Interface type of service related methods
type IService interface {
	CreateArticle(ctx context.Context, article *entity.Article) error
	GetArticle(ctx context.Context, articleID uint) (*entity.Article, error)
	GetArticlesByTagAndDate(ctx context.Context, tagName string, date time.Time) ([]*entity.Article, int64, error)
}
