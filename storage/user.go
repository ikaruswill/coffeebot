package storage

import (
	"github.com/ikaruswill/koffea/koffea"
	"gorm.io/gorm"
)

// User stores Telegram user metadata
type User struct {
	gorm.Model
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Username     string `gorm:"not null"`
	LanguageCode string
	IsBot        bool
	GroupBuys    []GroupBuy
	Orders       []Order
}

type UserService struct {
	Storage Storage
}

func NewUserService(storage Storage) *UserService {
	s := UserService{
		Storage: storage,
	}
	s.Storage.AutoMigrate(&User{})
	return &s
}

func (u *UserService) CreateIfNotExists(k *koffea.User) error {
	user := &User{
		FirstName:    k.FirstName,
		LastName:     k.LastName,
		Username:     k.Username,
		LanguageCode: k.LanguageCode,
		IsBot:        k.IsBot,
	}
	u.Storage.FirstOrCreate(&user)
	return nil
}
