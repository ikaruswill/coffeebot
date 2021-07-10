package user

import (
	"github.com/ikaruswill/koffeabot/internal/storage/groupbuy"
	"github.com/ikaruswill/koffeabot/internal/storage/order"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not null"`
	GroupBuys    []groupbuy.GroupBuy
	Orders       []order.Order
}
