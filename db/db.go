package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func ConDB(config string, dbType string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch strings.ToLower(dbType) {
	case "mysql":
		db, err = sql.Open("mysql", config)
	case "postgis", "postgresql":
		db, err = sql.Open("postgres", config)
	default:
		return nil, errors.New("unsupported database type")
	}

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Printf("Connected into %s !", dbType)
	return db, nil
}
