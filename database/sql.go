package database

import (
	"database/sql"
	"sync"
)

var (
	mu           = &sync.Mutex{}
	DBConnection *sql.DB
)

func GetDBConnection() *sql.DB {
	mu.Lock()
	defer mu.Unlock()
	return DBConnection
}

func InitDBConnection(driver, connectUrl string) (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()
	if DBConnection == nil {
		DBConnection, err := sql.Open(driver, connectUrl)
		return DBConnection, err
	}
	return DBConnection, nil
}
