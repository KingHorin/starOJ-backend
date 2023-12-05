package model

import (
	"gorm.io/gorm"
	"time"
)

type Submission struct {
	ID        int32     `gorm:"type:int;primary_key;auto_increment"`
	ProblemID int32     `gorm:"type:int;not null;index"`
	UserID    int32     `gorm:"type:int;not null;index"`
	Status    int32     `gorm:"type:int;not null"`
	Language  int32     `gorm:"type:int;not null"`
	Time      int32     `gorm:"type:int;not null"`
	Memory    int32     `gorm:"type:int;not null"`
	JudgedAt  time.Time `gorm:"type:timestamp;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
}

func MigrateSubmission(db *gorm.DB) {
	db.AutoMigrate(&Submission{})
}
