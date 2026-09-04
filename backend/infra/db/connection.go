package db

import (
	"ecommerce/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"strconv"
)

func GetConnectionString(conf config.Config) string {
	return "postgres://" + conf.POSTGRES_USER + ":" + conf.POSTGRES_PASSWORD + "@" + conf.POSTGRES_HOST + ":" + strconv.Itoa(conf.POSTGRES_PORT) + "/" + conf.POSTGRES_DB + "?sslmode=disable"
}

func NewConnection(conf config.Config) (*sqlx.DB, error) {
	connStr := GetConnectionString(conf)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
