package persist

import (
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

var (
	sqlxDb *sqlx.DB
	once   sync.Once
)

type DataSourceConfig struct {
	DsType      string
	Url         string
	Username    string
	Password    string
	MaxIdle     int           // zero means defaultMaxIdleConns; negative means 0
	MaxOpen     int           // <= 0 means unlimited
	MaxLifetime time.Duration // maximum amount of time a connection may be reused
}

func CompleteConfig(c *DataSourceConfig) {
	db := NewDb(c)
	if db != nil {
		sqlxDb = db
	}
}

func GetDb() *sqlx.DB {
	return sqlxDb
}

/*
	As a quick ref to application.yaml
 */
func NewDb(cfg *DataSourceConfig) *sqlx.DB {
	db := sqlx.MustOpen(cfg.DsType, cfg.Url)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	if cfg.MaxIdle != 0 {
		db.SetMaxIdleConns(cfg.MaxIdle)
	}

	if cfg.MaxOpen != 0 {
		db.SetMaxOpenConns(cfg.MaxOpen)
	}

	if cfg.MaxLifetime != 0 {
		db.SetConnMaxLifetime(cfg.MaxLifetime * time.Second)
	}

	return db
}
