package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Path string `yaml:"path"`
}

type Client struct {
	DB *gorm.DB
}

func (c *Client) init(config Config) error {
	var err error
	c.DB, err = gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
	return err
}

func NewConnection(config Config) (*Client, error) {
	c := &Client{}
	err := c.init(config)
	return c, err
}
