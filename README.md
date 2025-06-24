# gormpgconn

A lightweight wrapper for managing PostgreSQL connections using GORM.

## Simple Usage

```go
package main

import (
    "os"
    "github.com/apgolib/gormpgconn/pg"
    "gorm.io/gorm"
)

var db *gorm.DB

func main() {
    // Initialize database configuration from environment variables
    cfg := pg.NewDefaultConfig()
    cfg.Host = os.Getenv("DB_HOST")
    cfg.User = os.Getenv("DB_USER")
    cfg.Password = os.Getenv("DB_PASSWORD")
    
    // Initialize the GORM connection
    pg.Init(&db, cfg)

    // Use the connection
    gormResult := db.Create(user)
}
```

## Advansed usage 

### 1. Implement a config loader from environment variables
```go
package conf

import (
    "github.com/apgolib/gormpgconn/pg"
    "github.com/caarlos0/env/v10"
    "github.com/joho/godotenv"
    "log"
    "time"
)

type Env struct {
    PgHost               string        `env:"DB_HOST,required"`
    PgPort               int           `env:"DB_PORT" envDefault:"5432"`
    PgUser               string        `env:"DB_USER,required"`
    PgPassword           string        `env:"DB_PASSWORD,required"`
    PgSchema             string        `env:"DB_DB" envDefault:"postgres"`
    PgSSLMode            string        `env:"DB_SSLMODE" envDefault:"disable"`
    PgTimezone           string        `env:"DB_TIMEZONE" envDefault:"UTC"`
    PgMaxOpenConnections int           `env:"DB_MAX_OPEN_CONNECTIONS" envDefault:"10"`
    PgMaxIdleConnections int           `env:"DB_MAX_IDLE_CONNECTIONS" envDefault:"1"`
    PgConnMaxLifetime    time.Duration `env:"DB_CONN_MAX_LIFETIME" envDefault:"30s"`
}

var envs *Env

func (e *Env) GetPGConfig() pg.Config {
    return pg.Config{
        Host:               e.PgHost,
        Port:               e.PgPort,
        User:               e.PgUser,
        Password:           e.PgPassword,
        DbName:             e.PgSchema,
        SSLMode:            e.PgSSLMode,
        Timezone:           e.PgTimezone,
        MaxOpenConnections: e.PgMaxOpenConnections,
        MaxIdleConnections: e.PgMaxIdleConnections,
        ConnMaxLifetime:    e.PgConnMaxLifetime,
    }
}

func InitEnvs() {
    _ = godotenv.Load()
    envs = &Env{}
    if err := env.Parse(envs); err != nil {
        log.Fatalf("Failed to parse environment variables: %v", err)
    }
}

func GetEnvs() *Env {
    if envs == nil {
        log.Fatal("Environment variables not initialized")
    }
    return envs
}

```

### 2. Implement the connector

```go
package database

import (
    "github.com/apgolib/gormpgconn/pg"
    "gorm.io/gorm"
)

var db *gorm.DB

func InitDb(cfg pg.Config) {
    pg.Init(&db, cfg)
}

func GetDb() *gorm.DB {
    return pg.Get(db)
}
```


### 3. Set up and use the connection
```go
func main() {
    // Initialize environment and database
    conf.InitEnvs()
    database.InitDb(conf.GetEnvs().GetPGConfig())

    // Use the GORM connection
    gormResult := database.GetDb().Create(user)
}
```
