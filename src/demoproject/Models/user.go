package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"not null;size:255" json:"username" validate:"required"`
	Email       string `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password    string `gorm:"not null;size:255" json:"password" validate:"required"`
	ActiveAt    sql.NullTime
	Role        uint8         `gorm:"not null;default:0;size:5" json:"role"`
	Tournaments []*Tournament `gorm:"many2many:user_tournaments"`
	Duel        []*Duel       `gorm:"many2many:user_duels"`
	Rank        Rank
}

type UpdateUser struct {
	ID        uint   `json:"-"`
	Username  string `json:"username" validate:"required"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (UpdateUser) TableName() string { return "users" }

func (User) HasPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (User) CheckHasPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password, err = user.HasPassword(user.Password)

	return
}
