package groupbuy

import (
	"github.com/ikaruswill/koffeabot/internal/storage/order"
	"gorm.io/gorm"
)

type GroupBuy struct {
	gorm.Model
	UserID uint
	Orders []order.Order
}
