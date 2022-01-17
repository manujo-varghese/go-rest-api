package database

import (
	"github.com/jinzhu/gorm"
	"github.com/manujo-varghese/go-rest-api/internal/article"
)

// MigrateDB - migrates database and create article table
func MigrateDB(db *gorm.DB) error  {
	if result := db.AutoMigrate(&article.Article{}); result.Error != nil{
		return result.Error
	}
	return nil
	
}