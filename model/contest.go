package model

import (
	"gorm.io/gorm"
	"time"
)

type Contest struct {
	ID          int32          `gorm:"type:int;primary_key;auto_increment"`
	Name        string         `gorm:"type:varchar(50);not null"`
	Host        string         `gorm:"type:varchar(50);not null"`
	StartTime   string         `gorm:"type:timestamp;not null"`
	EndTime     string         `gorm:"type:timestamp;not null"`
	Description string         `gorm:"type:text;not null"`
	CreatedBy   int32          `gorm:"type:int;not null"`
	CreatedAt   time.Time      `gorm:"type:timestamp;autoCreateTime;not null" json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp" json:"-"`
	Problems    []Problem      `gorm:"many2many:contest_problem"`
}

func MigrateContest(db *gorm.DB) {
	db.AutoMigrate(&Contest{})
}
