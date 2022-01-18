package article

import (
	"strconv"
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

type Result struct{
	Tag string
	Count int
	Articles pq.StringArray `gorm:"type:text[]"`
	RelatedTags pq.StringArray `gorm:"type:text[]"`

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
func (s *Service) GetArticleByTagDate(tag string, date string) (Result, error)  {
	var articles []Article
	var articlesArray []string = nil
	var relatedTags = make(map[string]bool)
	var uniqueRelatedFlags = []string{}
	var count int
	const layout = "2006-01-02"
	slicedDate := date[0:4] + "-" + date[4:6] + "-" + date[6:8]
	tm, _ := time.Parse(layout,slicedDate)
	if result := s.DB.Where("Date = ? ",tm).Find(&articles); result.Error != nil{
		return Result{}, result.Error
	}
	for _,article := range articles{
		articlesArray = append(articlesArray, strconv.FormatInt(int64(article.ID),10))
		for _,tagLoop := range article.Tags{
			count = count + len(article.Tags)
			if tagLoop == tag{
				for _,selectedTags := range article.Tags{
					if (!relatedTags[selectedTags]) && selectedTags != tag {
						uniqueRelatedFlags = append(uniqueRelatedFlags, selectedTags)
						relatedTags[selectedTags] = true
					}
				}
			}
		}
	}
	var result Result
	result.Tag = tag
	result.Count = count
	firstTen := len(articlesArray)
	if firstTen > 10{
		firstTen = 10
	}
	result.Articles = articlesArray[0:firstTen]
	result.RelatedTags = uniqueRelatedFlags
	return result, nil
}

// PostArticle - adds a new article 
func (s *Service) PostArticle(article Article) (Article, error)  {
	if result := s.DB.Save(&article); result.Error != nil{
		return Article{}, result.Error
	}
	return article, nil
	
}
