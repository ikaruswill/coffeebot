package config

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/ikaruswill/koffea/client/postgres"
	"github.com/ikaruswill/koffea/client/sqlite"
	"github.com/ikaruswill/koffea/client/telegram"
	"gopkg.in/yaml.v2"
)

type StorageDriver string

const (
	SqliteStorageDriver   StorageDriver = "sqlite"
	PostgresStorageDriver StorageDriver = "postgres"
)

type StorageConfig struct {
	Driver   StorageDriver   `yaml:"driver"`
	Sqlite   sqlite.Config   `yaml:"sqlite"`
	Postgres postgres.Config `yaml:"postgres"`
}

type Config struct {
	Telegram telegram.Config `yaml:"telegram"`
	Storage  StorageConfig   `yaml:"storage"`
}

func (s *StorageConfig) validateDriver() error {
	switch s.Driver {
	case SqliteStorageDriver, PostgresStorageDriver:
		return nil
	}
	return errors.New("invalid storage driver")
}

func (c *Config) Load(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	err = c.Storage.validateDriver()
	return err
}
