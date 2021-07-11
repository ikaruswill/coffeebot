package storage

import (
	"gorm.io/gorm"
)

// User stores Telegram user metadata
type User struct {
	gorm.Model
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Username     string `gorm:"not null"`
	LanguageCode string
	IsBot        bool
	GroupBuys    []GroupBuy
	Orders       []Order
}
