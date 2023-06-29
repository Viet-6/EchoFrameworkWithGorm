package models

import "gorm.io/gorm"

type Duel struct {
	gorm.Model
	Duration uint8 `gorm:"not null;size:30"`
	WinnerID uint
	Status   bool    `gorm:"not null;default:true"`
	Users    []*User `gorm:"many2many:user_duels"`
}
