package database

import (
	"github.com/jinzhu/gorm"
	"github.com/singfield/rest-api-golang/internal/comment"
)

// migrates our databases and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}