package storage

import (
	"github.com/ikaruswill/koffea/internal/config"
	"github.com/ikaruswill/koffea/internal/storage/groupbuy"
	"github.com/ikaruswill/koffea/internal/storage/order"
	"github.com/ikaruswill/koffea/internal/storage/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func (c *Client) Init(config.DBConfig) error {
	var err error
	c.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	c.DB.AutoMigrate(
		&user.User{},
		&groupbuy.GroupBuy{},
		&order.Order{},
	)
	return err
}
