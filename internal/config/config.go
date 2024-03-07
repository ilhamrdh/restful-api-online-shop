package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App      AppConfig `yaml:"app"`
	DBConfig DBConfig  `yaml:"database"`
}

type AppConfig struct {
	Name       string           `yaml:"name"`
	Port       string           `yaml:"port"`
	Encryption EncryptionConfig `yaml:"encryption"`
}

type EncryptionConfig struct {
	Salt      uint8  `yaml:"salt"`
	JWTSecret string `yaml:"jwt_secret"`
}

type DBConfig struct {
	Host           string                `yaml:"host"`
	Port           string                `yaml:"port"`
	User           string                `yaml:"user"`
	Pass           string                `yaml:"pass"`
	Name           string                `yaml:"name"`
	ConnectionPool DBConnctionPoolConfig `yaml:"connection_pool"`
}

type DBConnctionPoolConfig struct {
	MaxIdleConnection     uint8 `yaml:"max_idle_connection"`
	MaxOpenConnection     uint8 `yaml:"max_open_connection"`
	MaxLifetimeConnection uint8 `yaml:"max_lifetime_connection"`
	MaxIdletimeConnection uint8 `yaml:"max_idletime_connection"`
}

var Cfg Config

func LoadConfig(filename string) error {
	configByte, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(configByte, &Cfg)
}
