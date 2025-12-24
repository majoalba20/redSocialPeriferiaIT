package models

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Message   string `gorm:"type:text;not null"`
	ProfileID uint
	Profile   Profile `gorm:"constraint:OnDelete:CASCADE;"`
	Likes     int     `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
