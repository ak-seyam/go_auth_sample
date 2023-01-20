package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Password  string
	Username  string `gorm:"size:256"`
	Groups    string `gorm:"default:'users'"`
}
