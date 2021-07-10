package user

import (
	"github.com/ikaruswill/koffea/internal/storage/groupbuy"
	"github.com/ikaruswill/koffea/internal/storage/order"
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
