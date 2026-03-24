package models

type User struct {
	ID           string `gorm:"type:uuid;primaryKey" db:"id"`
	Email        string `json:"email" gorm:"unique;not null"`
	PasswordHash string `json:"password" gorm:"not null"`
}
