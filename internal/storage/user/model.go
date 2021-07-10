package user

import (
	"github.com/ikaruswill/koffeabot/internal/storage/groupbuy"
	"github.com/ikaruswill/koffeabot/internal/storage/order"
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
	GroupBuys    []groupbuy.GroupBuy
	Orders       []order.Order
}
