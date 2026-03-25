package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID        int64   `gorm:"type:bigserial;primaryKey" db:"id"`
	Name      string  `json:"name" gorm:"not null"`
	Email     string  `json:"email" gorm:"unique;not null"`
	Phone     string  `json:"phone" gorm:"not null"`
	Address   string  `json:"address" gorm:"not null"`
	Longitude float64 `json:"longitude" gorm:"not null"`
	Latitude  float64 `json:"latitude" gorm:"not null"`
	Status    string  `json:"status" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
