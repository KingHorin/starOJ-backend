package model

import (
	"gorm.io/gorm"
	"time"
)

type Problem struct {
	ID           int32          `gorm:"type:int;primary_key;auto_increment"`
	Name         string         `gorm:"type:varchar(50);not null;index"`
	TimeLimits   int32          `gorm:"type:int;not null"`
	MemoryLimits int32          `gorm:"type:int;not null"`
	Description  string         `gorm:"type:text;not null"`
	InputFormat  string         `gorm:"type:text;not null"`
	OutputFormat string         `gorm:"type:text;not null"`
	Note         string         `gorm:"type:text;not null"`
	CreatedBy    int32          `gorm:"type:int;not null"`
	CreatedAt    time.Time      `gorm:"type:timestamp;autoCreateTime;not null"`
	UpdatedAt    time.Time      `gorm:"type:timestamp;autoUpdateTime;not null"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamp;index"`
}

type ProblemTag struct {
	ID   int32  `gorm:"type:int;primary_key;auto_increment"`
	Name string `gorm:"type:varchar(50);not null;index"`
}

type Submission struct {
	ID        int32     `gorm:"type:int;primary_key;auto_increment"`
	ProblemID int32     `gorm:"type:int;not null;index"`
	UserID    int32     `gorm:"type:int;not null;index"`
	Status    int32     `gorm:"type:int;not null;index"`
	Time      int32     `gorm:"type:int;not null;index"`
	Memory    int32     `gorm:"type:int;not null;index"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime;not null"`
}

func MigrateProblem(db *gorm.DB) {
	db.AutoMigrate(&Problem{})
	db.AutoMigrate(&ProblemTag{})
	db.AutoMigrate(&Submission{})
}
