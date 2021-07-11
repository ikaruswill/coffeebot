package storage

import (
	"gorm.io/gorm"
)

type GroupBuy struct {
	gorm.Model
	UserID uint
	Orders []Order
}
