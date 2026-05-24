package database

import "time"

type User struct {
	ID           int64     `gorm:"primaryKey"`
	Email        string    `gorm:"not null;uniqueIndex"`
	Name         string    `gorm:"not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}
