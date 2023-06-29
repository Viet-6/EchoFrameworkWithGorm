package models

import (
	"time"

	"gorm.io/gorm"
)

type UserDuel struct {
	UserID   uint `gorm:"primaryKey"`
	DuelID   uint `gorm:"primaryKey"`
	Point    uint `gorm:"not null;default:0"`
	CreateAt time.Time
	DeleteAt gorm.DeletedAt `gorm:"index"`
}
