package article

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// Service - the struct for article service
type Service struct{
	DB *gorm.DB
}

// Article - the model structure for article 
type Article struct{
	gorm.Model
	Title string
	Date time.Time
	Body string
	Tags pq.StringArray `gorm:"type:text[]"`

}

// ArticleService - the interface for article service
type ArticleService interface{
	GetArticle(ID uint) (Article, error)
	GetArticleByTagDate(tag string, date string) ([]Article, error)
	PostArticle(article Article) (Article, error)
}

// NewService - returns a new article service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
	
}

// GetArticle - retrives article by its ID from the database
func (s *Service) GetArticle(ID uint) (Article, error)  {
	var article Article
	if result := s.DB.First(&article, ID); result.Error != nil{
		return Article{}, result.Error
	}
	return article, nil
}

// GetArticleByTagDate - retrives all articles by tag and date (path -/api/article/{tag}/{date})
func (s *Service) GetArticleByTagDate(tag string, date string) ([]Article, error)  {
	var articles []Article
	if result := s.DB.Find(&articles).Where("tag = ?", tag); result.Error != nil{
		return []Article{}, result.Error
	}	
	return articles, nil
}

// PostArticle - adds a new article 
func (s *Service) PostArticle(article Article) (Article, error)  {
	if result := s.DB.Save(&article); result.Error != nil{
		return Article{}, result.Error
	}
	return article, nil
	
}