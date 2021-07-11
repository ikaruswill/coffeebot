package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Path string `yaml:"path"`
}

func NewConnection(config Config) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
}
