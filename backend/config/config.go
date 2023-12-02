package config

import "crypto/rsa"

var Conf *Config

// for jwt
var PrivateKey *rsa.PrivateKey

type DatabaseConfig struct {
	Host     string `env:"HOST"`
	Port     uint16 `env:"PORT"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
}

type ServerConfig struct {
	Port         uint16 `env:"PORT"`
    AiHost string `env:"AI_HOST"`
	AllowOrigins string `env:"ALLOW_ORIGINS"`
}

type Config struct {
	DbConfig  DatabaseConfig `envPrefix:"DB_"`
	SrvConfig ServerConfig   `envPrefix:"SERVER_"`
}
