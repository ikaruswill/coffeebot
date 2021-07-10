package groupbuy

import (
	"github.com/ikaruswill/koffea/storage/order"
	"gorm.io/gorm"
)

type GroupBuy struct {
	gorm.Model
	UserID uint
	Orders []order.Order
}
