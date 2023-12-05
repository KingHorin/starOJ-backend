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
	SPJ          bool           `gorm:"type:tinyint;not null"`
	TestCase     int32          `gorm:"type:int;not null" json:"-"`
	Difficulty   int32          `gorm:"type:int;not null;index"`
	CreatedBy    int32          `gorm:"type:int;not null"`
	CreatedAt    time.Time      `gorm:"type:timestamp;autoCreateTime;not null" json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamp;index" json:"-"`
	Tags         []Tag          `gorm:"many2many:problem_tag"`
}

type Tag struct {
	ID    int32  `gorm:"type:int;primary_key;auto_increment"`
	Name  string `gorm:"type:varchar(50);not null;index"`
	Color string `gorm:"type:char(7);not null"`
}

func MigrateProblem(db *gorm.DB) {
	db.AutoMigrate(&Problem{})
	db.AutoMigrate(&Tag{})
}
