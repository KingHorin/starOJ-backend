package model

import (
	"gorm.io/gorm"
	"time"
)

type Sheet struct {
	ID          int32          `gorm:"type:int;primary_key;auto_increment"`
	Name        string         `gorm:"type:varchar(50);not null"`
	Cover       string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text;not null"`
	CreatedBy   int32          `gorm:"type:int;not null"`
	CreatedAt   time.Time      `gorm:"type:timestamp;autoCreateTime;not null"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp" json:"-"`
	Problems    []Problem      `gorm:"many2many:sheet_problem"`
}

func MigrateSheet(db *gorm.DB) {
	db.AutoMigrate(&Sheet{})
}
