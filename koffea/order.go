package koffea

import "time"

type Order struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ProductName    string
	Quantity       uint8
	AdvancePayment bool
	HasPaid        bool
	GroupBuy       GroupBuy
	User           User
}

type OrderService interface {
	Create(*Order) error
	GetByUser(string) (*Order, error)
	Update(*Order) error
	Delete(*Order) error
}
