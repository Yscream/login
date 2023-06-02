package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func UnmarshalYAML(path string) *Config {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("couldn't load file: %s", err)
	}

	var config Config

	err = v.Unmarshal(&config)
	if err != nil {
		log.Fatalf("coulnd't unmarshal: %s", err)
	}

	return &config
}

type Config struct {
	DB   *DBConfig `yaml:"db"`
	Salt string    `yaml:"salt"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
	SSLMode  string `yaml:"sslmode"`
}

func (c *DBConfig) URL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.Username, c.Password, c.Host, c.Port, c.DBName, c.SSLMode,
	)
}
