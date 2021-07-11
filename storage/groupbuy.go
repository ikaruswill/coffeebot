package storage

import (
	"time"

	"github.com/ikaruswill/koffea/koffea"
	"gorm.io/gorm"
)

type PaymentMethod string

const (
	CashPaymentMethod   PaymentMethod = "cash"
	PayNowPaymentMethod PaymentMethod = "paynow"
)

type GroupBuy struct {
	gorm.Model
	RoasterName          string
	LocationID           uint
	Location             Location
	Expiry               time.Time
	PaymentMethod        PaymentMethod
	AllowAdvancePayment  bool
	CollectionPostalCode uint
	Comment              string
	UserID               uint
	Orders               []Order
}

type GroupBuyService struct {
	Storage Storage
}

func NewGroupBuyService(storage Storage) *GroupBuyService {
	s := GroupBuyService{
		Storage: storage,
	}
	s.Storage.AutoMigrate(&GroupBuy{})
	return &s
}

func (g *GroupBuyService) Create(k *koffea.GroupBuy) error {
	panic("not implemented")
}
func (g *GroupBuyService) List() (*[]koffea.GroupBuy, error) {
	panic("not implemented")
}
func (g *GroupBuyService) Update(k *koffea.GroupBuy) error {
	panic("not implemented")
}
func (g *GroupBuyService) Delete(k *koffea.GroupBuy) error {
	panic("not implemented")
}
