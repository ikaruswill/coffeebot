package koffea

type User struct {
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	IsBot        bool
	GroupBuys    []GroupBuy
	Orders       []Order
}

type UserService interface {
	CreateIfNotExists(*User) error
}
