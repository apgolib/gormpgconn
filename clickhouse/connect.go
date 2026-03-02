package pg

import (
	"fmt"
	"log"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

func Init(db **gorm.DB, cfg Config) {
	dsn := fmt.Sprintf(
		"clickhouse://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName,
	)

	var err error
	*db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if cfg.Debug {
		*db = (*db).Debug()
	}

	sqlDB, err := (*db).DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}

	// Configure connection pool
	if cfg.MaxOpenConnections > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	}
	if cfg.MaxIdleConnections > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	}
	if cfg.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}
}

func Get(db *gorm.DB) *gorm.DB {
	if db == nil {
		log.Fatal("database not initialized")
	}
	return db
}
