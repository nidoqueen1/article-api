package postgresql

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/nidoqueen1/article-api/entity"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Stores a new article into database
func (p *postgresql) CreateArticle(ctx context.Context, article *entity.Article) error {
	tagNames := getTagNames(*article)

	// insert new Tags and return both new and existing Tag IDs by updating existing Tags
	// raw query has been used because (*gorm.DB).Create returns ID only for the new entries by defaulf
	var tagIDs []uint
	query := `
        INSERT INTO tags (name)
        VALUES %s
        ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name 
        RETURNING id`

	valueStrings := make([]string, 0, len(tagNames))
	for _, tagName := range tagNames {
		valueStrings = append(valueStrings, fmt.Sprintf("('%s')", tagName))
	}

	query = fmt.Sprintf(query, strings.Join(valueStrings, ", "))
	if err := p.db.WithContext(ctx).Raw(query).Scan(&tagIDs).Error; err != nil {
		return err
	}

	// insert the article without tags first, clearing tags to avoid auto-insertion them
	article.Tags = nil
	if err := p.db.WithContext(ctx).Create(article).Error; err != nil {
		return err
	}

	// then insert article_tag relationships
	relationshipQuery := `
        INSERT INTO article_tags (article_id, tag_id)
        VALUES %s`

	articleTagValues := make([]string, 0, len(tagIDs))
	for _, tagID := range tagIDs {
		articleTagValues = append(articleTagValues, fmt.Sprintf("(%d, %d)", article.ID, tagID))
	}

	relationshipQuery = fmt.Sprintf(relationshipQuery, strings.Join(articleTagValues, ", "))
	if err := p.db.WithContext(ctx).Exec(relationshipQuery).Error; err != nil {
		return err
	}

	return nil
}

// Fetches an Article by its ID
func (p *postgresql) GetArticle(ctx context.Context, articleID uint) (*entity.Article, error) {
	var article *entity.Article
	err := p.db.WithContext(ctx).Preload("Tags").First(&article, articleID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		article = nil
	}

	return article, err
}

// Fetches a list of Articles filtered by their Tag name and Date
func (p *postgresql) GetArticlesByTagAndDate(ctx context.Context, tagName string, date time.Time) ([]*entity.Article, int64, error) {
	var articles []*entity.Article
	var totalCount int64

	subquery := p.db.WithContext(ctx).Model(&entity.Article{}).
		Joins("JOIN article_tags ON articles.id = article_tags.article_id").
		Joins("JOIN tags ON tags.id = article_tags.tag_id").
		Where("tags.name = ?", tagName).
		Where("articles.date = ?", date)

	err := subquery.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	err = subquery.Limit(viper.GetInt("list_articles.limit")).
		Preload("Tags").
		Order("articles.id DESC").
		Find(&articles).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	return articles, totalCount, err
}
