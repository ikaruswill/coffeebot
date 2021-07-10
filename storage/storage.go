package storage

import (
	"fmt"

	"github.com/ikaruswill/koffea/config"
	"github.com/ikaruswill/koffea/storage/groupbuy"
	"github.com/ikaruswill/koffea/storage/order"
	"github.com/ikaruswill/koffea/storage/user"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func (c *Storage) Init(cfg config.StorageConfig) error {
	var err error
	switch cfg.Driver {
	case config.SqliteStorageDriver:
		c.DB, err = gorm.Open(sqlite.Open(cfg.Sqlite.Path), &gorm.Config{})
	case config.PostgresStorageDriver:
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database, cfg.Postgres.Port)
		c.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
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
