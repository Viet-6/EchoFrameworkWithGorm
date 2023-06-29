package models

import (
	"time"

	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	NumberOfCompetetor uint8     `gorm:"not null;default:0;size:100" json:"number_of_competetor" validate:"required,number,min=10,max=100"`
	StartAt            time.Time `gorm:"not null" json:"start_at" validate:"required,datetime"`
	EndRegister        time.Time `gorm:"not null" json:"end_register" validate:"required,datetime,gtcsfield=start_at"`
	Duration           uint8     `gorm:"not null;size:30" json:"duration" validate:"required,number,min=10,max=30"`
	Status             bool      `gorm:"not null;default:true"`
	Users              []*User   `gorm:"many2many:user_tournaments"`
}
