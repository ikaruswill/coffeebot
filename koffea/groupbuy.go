package koffea

import "time"

type GroupBuy struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Orders    []Order
}

type GroupBuyService interface {
	Create(*GroupBuy) error
	List() (*[]GroupBuy, error)
	Update(*GroupBuy) error
	Delete(*GroupBuy) error
}
