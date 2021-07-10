package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Client struct {
	DB *gorm.DB
}

func (c *Client) init(config Config) error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.Database, config.Port)
	c.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func NewConnection(config Config) (*Client, error) {
	c := &Client{}
	err := c.init(config)
	return c, err
}
