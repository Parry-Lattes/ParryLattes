package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)


var db *sql.DB

func ConnectDB() error {
	//defer db.Close()
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "1234"
	cfg.Net = "tcp"
	cfg.Addr = "ParryDB:3306"
	cfg.DBName = "mydb"

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}

	fmt.Println("Connected")

	return nil
}

func GetDBHandle() *sql.DB {
	return db
}

func CloseDB() {
	db.Close()
}
