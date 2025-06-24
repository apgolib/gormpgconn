package pg

import "time"

type Config struct {
	Host               string
	Port               int
	User               string
	Password           string
	DbName             string
	SSLMode            string
	Timezone           string
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    time.Duration
}

func NewDefaultConfig() Config {
	return Config{
		Port:               5432,
		DbName:             "postgres",
		SSLMode:            "disable",
		Timezone:           "UTC",
		MaxOpenConnections: 10,
		MaxIdleConnections: 1,
		ConnMaxLifetime:    30 * time.Second,
	}
}
