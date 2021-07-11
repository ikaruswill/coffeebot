package storage

import (
	"fmt"

	"github.com/ikaruswill/koffea/client/postgres"
	"github.com/ikaruswill/koffea/client/sqlite"
	"github.com/ikaruswill/koffea/config"
	"gorm.io/gorm"
)

type Storage struct {
	*gorm.DB
}

func NewStorage(cfg config.StorageConfig) (Storage, error) {
	var (
		db  *gorm.DB
		err error
	)
	switch cfg.Driver {
	case config.SqliteStorageDriver:
		db, err = sqlite.NewConnection(cfg.Sqlite)
	case config.PostgresStorageDriver:
		db, err = postgres.NewConnection(cfg.Postgres)
	default:
		err = fmt.Errorf("invalid driver: %s", cfg.Driver)
	}

	return Storage{db}, err
}
