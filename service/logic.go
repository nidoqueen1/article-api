package service

import (
	"context"
	"time"

	"github.com/nidoqueen1/article-api/entity"
)

// Service logic of creation an Article
func (s *service) CreateArticle(ctx context.Context, article *entity.Article) error {
	s.logger.Infof("Creating article with title: %s", article.Title)
	if article.Date.IsZero() {
		todayDate := time.Date(
			time.Now().Year(), time.Now().Month(), time.Now().Day(),
			0, 0, 0, 0, time.UTC)
		article.Date = todayDate
	}

	if err := s.db.CreateArticle(ctx, article); err != nil {
		s.logger.Errorf("Error creating article: %v", err)
		return err
	}
	s.logger.Infof("Successfully created article with ID: %d", article.ID)
	return nil
}

// Service logic of fetching an Article by its ID
func (s *service) GetArticle(ctx context.Context, articleID uint) (*entity.Article, error) {
	s.logger.Infof("Fetching article with ID: %d", articleID)
	article, err := s.db.GetArticle(ctx, articleID)

	if err != nil {
		s.logger.Errorf("Error fetching article with ID %d: %v", articleID, err)
		return nil, err
	}
	if article == nil {
		s.logger.Warnf("No article found with ID: %d", articleID)
	}
	return article, nil
}

// Service logic of fetching a list of Articles filtered by their Tag name and Date
func (s *service) GetArticlesByTagAndDate(ctx context.Context, tagName string, date time.Time) ([]*entity.Article, int64, error) {
	s.logger.Infof("Fetching articles for tag: %s on date: %s", tagName, date.Format("2006-01-02"))
	articles, count, err := s.db.GetArticlesByTagAndDate(ctx, tagName, date)

	if err != nil {
		s.logger.Errorf("Error fetching articles for tag %s on date %s: %v", tagName, date.Format("2006-01-02"), err)
		return nil, 0, err
	}
	s.logger.Infof("Found %d articles for tag: %s on date: %s", len(articles), tagName, date.Format("2006-01-02"))
	return articles, count, nil
}
