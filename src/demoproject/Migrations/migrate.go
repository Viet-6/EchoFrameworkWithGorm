package migrations

import (
	m "demoproject/Models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {

	db.AutoMigrate(
		&m.User{},
		&m.Tournament{},
		&m.Rank{},
		&m.Duel{},
		&m.UserTournament{},
		&m.UserDuel{},
	)
}
