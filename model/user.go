package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int32          `gorm:"type:int;primary_key;auto_increment"`
	Username  string         `gorm:"type:varchar(20);not null;index"`
	Nickname  string         `gorm:"type:varchar(50);not null;index"`
	Password  string         `gorm:"type:char(64);not null"`
	Avatar    string         `gorm:"type:varchar(255);not null"`
	Phone     string         `gorm:"type:varchar(20);not null"`
	Email     string         `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp;autoCreateTime;not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp"`
}

func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
