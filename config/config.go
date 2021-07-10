package config

import (
	"io/ioutil"
	"log"

	"github.com/ikaruswill/koffea/client/postgres"
	"github.com/ikaruswill/koffea/client/sqlite"
	"github.com/ikaruswill/koffea/client/telegram"
	"gopkg.in/yaml.v2"
)

type StorageDriver string

const (
	SqliteStorageDriver   = "sqlite"
	PostgresStorageDriver = "postgres"
)

type StorageConfig struct {
	Driver   StorageDriver   `yaml:"driver"`
	Sqlite   sqlite.Config   `yaml:"sqlite"`
	postgres postgres.Config `yaml:"postgres"`
}

type Config struct {
	Telegram telegram.Config `yaml:"telegram"`
	Storage  StorageConfig   `yaml:"storage"`
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
	return nil
}
