package models

type Rank struct {
	UserID uint
	Point  uint `gorm:"not null;default:0"`
}
