package koffea

import "time"

type Order struct {
	CreatedAt time.Time
	GroupBuy  GroupBuy
	User      User
}

type OrderService interface {
	Create(*Order) error
	GetByUser(string) (*Order, error)
	Update(*Order) error
	Delete(*Order) error
}
