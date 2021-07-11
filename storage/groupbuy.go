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
	RoasterName          string `gorm:"not null"`
	LocationID           uint   `gorm:"not null"`
	Location             Location
	Expiry               time.Time
	PaymentMethod        PaymentMethod `gorm:"not null"`
	AllowAdvancePayment  bool          `gorm:"not null"`
	CollectionPostalCode uint          `gorm:"not null"`
	Comment              string
	UserID               uint `gorm:"not null"`
	User                 User
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
