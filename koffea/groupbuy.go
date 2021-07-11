package koffea

import "time"

type PaymentMethod string

const (
	CashPaymentMethod   PaymentMethod = "Cash"
	PayNowPaymentMethod PaymentMethod = "PayNow"
)

type GroupBuy struct {
	RoasterName          string
	Location             string
	StartedAt            time.Time
	Expiry               time.Time
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
