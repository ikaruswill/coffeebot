package koffea

import "time"

type PaymentMethod string

const (
	CashPaymentMethod   = "Cash"
	PayNowPaymentMethod = "PayNow"
)

type GroupBuy struct {
	RoasterName          string
	StartedAt            time.Time
	EndedAt              time.Time
	PaymentMethod        PaymentMethod
	AllowAdvancePayment  bool
	CollectionPostalCode uint
	Comment              string
	Organizer            User
	Orders               []Order
}

type GroupBuyService interface {
	Create(*GroupBuy) error
	List() (*[]GroupBuy, error)
	Update(*GroupBuy) error
	Delete(*GroupBuy) error
}
