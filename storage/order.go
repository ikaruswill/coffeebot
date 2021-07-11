package storage

import (
	"github.com/ikaruswill/koffea/koffea"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductName    string `gorm:"not null"`
	Quantity       uint8  `gorm:"not null"`
	AdvancePayment bool
	HasPaid        bool
	GroupBuyID     uint `gorm:"not null"`
	UserID         uint `gorm:"not null"`
	User           User
}

type OrderService struct {
	Storage Storage
}

func NewOrderService(storage Storage) *OrderService {
	s := OrderService{
		Storage: storage,
	}
	s.Storage.AutoMigrate(&Order{})
	return &s
}

func (o *OrderService) Create(k *koffea.Order) error {
	panic("not implemented")
}
func (o *OrderService) GetByUser(string) (*koffea.Order, error) {
	panic("not implemented")
}
func (o *OrderService) Update(k *koffea.Order) error {
	panic("not implemented")
}
func (o *OrderService) Delete(k *koffea.Order) error {
	panic("not implemented")
}
