package database

import (
	"github.com/HasanShahjahan/go-rest-api/internal/comment"
	"github.com/eatigo/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
