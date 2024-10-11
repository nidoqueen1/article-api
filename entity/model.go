package entity

import "time"

// Internal DB-format of Article
type Article struct {
	ID    uint      `gorm:"primaryKey"`
	Title string    `gorm:"column:title" validate:"required,min=1,max=200"`
	Body  string    `gorm:"column:body" validate:"required,min=1,max=2000"`
	Date  time.Time `gorm:"column:date"` // todo create index
	Tags  []*Tag    `gorm:"many2many:article_tags;"`
}

// Internal DB-format of Tag
type Tag struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique" validate:"required,alphanum,min=1,max=50"` // todo create index
}

// Database specific model to serve many-two-many relationship between Article and Tag
type ArticleTag struct {
	ArticleID uint `gorm:"column:article_id"`
	TagID     uint `gorm:"column:tag_id"`
}
