package models

import (
	"time"

	"gorm.io/gorm"
)

type UserTournament struct {
	UserID          uint  `gorm:"primaryKey"`
	TournamentID    uint  `gorm:"primaryKey"`
	ParticipateRole uint8 `gorm:"not null;default:0;size:5"`
	CreateAt        time.Time
	DeleteAt        gorm.DeletedAt `gorm:"index"`
}
