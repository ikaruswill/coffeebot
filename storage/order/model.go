package order

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	GroupBuyID uint
	UserID     uint
}
