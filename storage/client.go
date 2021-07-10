package storage

import (
	"github.com/ikaruswill/koffea/config"
	"github.com/ikaruswill/koffea/storage/groupbuy"
	"github.com/ikaruswill/koffea/storage/order"
	"github.com/ikaruswill/koffea/storage/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func (c *Client) Init(cfg config.StorageConfig) error {
	var err error
	c.DB, err = gorm.Open(sqlite.Open(cfg.Sqlite.Path), &gorm.Config{})
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
