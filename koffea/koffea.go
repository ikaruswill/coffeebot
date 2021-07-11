package koffea

import (
	"github.com/ikaruswill/koffea/client/telegram"
)

type Koffea struct {
	UserService     UserService
	GroupBuyService GroupBuyService
	OrderService    OrderService
	TelegramService *telegram.Client
}
