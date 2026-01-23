package entities

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string    `gorm:"not null"`
	Password string    `json:"-" gorm:"not null"`
	Email    string    `gorm:"uniqueIndex;not null"`
	Avatar   *string
}
