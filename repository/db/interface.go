package db

import (
	"context"
	"time"

	"github.com/nidoqueen1/article-api/entity"
)

// Interface type of general database related methods
type IDatabase interface {
	CreateArticle(ctx context.Context, article *entity.Article) error
	GetArticle(ctx context.Context, articleID uint) (*entity.Article, error)
	GetArticlesByTagAndDate(ctx context.Context, tagName string, date time.Time) ([]*entity.Article, int64, error)
}
