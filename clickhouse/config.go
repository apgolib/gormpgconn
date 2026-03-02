package pg

import "time"

type Config struct {
	Host               string
	Port               int
	User               string
	Password           string
	DbName             string
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    time.Duration
	Debug              bool
}

func NewDefaultConfig() Config {
	return Config{
		Port:               9000,
		DbName:             "default",
		MaxOpenConnections: 10,
		MaxIdleConnections: 1,
		ConnMaxLifetime:    30 * time.Second,
		Debug:              false,
	}
}
