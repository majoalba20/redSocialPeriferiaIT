package models

import "time"

type Profile struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"uniqueIndex"`
	Name      string
	LastName  string
	Alias     string `gorm:"uniqueIndex"`
	BirthDate time.Time
	Bio       string
	AvatarURL string

	Posts []Post
}
