package migrations

import (
	m "demoproject/Models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	user := m.User{}

	result := db.Limit(1).Find(&user)

	if result.RowsAffected == 0 || user.Username != "admin" {

		user.Username = "admin"
		user.Email = "admin@admin.com"
		user.Password = "admin"
		user.ActiveAt.Time = time.Now()
		user.ActiveAt.Valid = true
		user.Role = 3

		db.Create(&user)
	}

	fmt.Println()
}
