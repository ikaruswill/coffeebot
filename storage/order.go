package storage

import (
	"github.com/ikaruswill/koffea/koffea"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint
	GroupBuyID uint
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
