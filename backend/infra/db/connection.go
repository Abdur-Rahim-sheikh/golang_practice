package db

import (
	"ecommerce/config"
	"fmt"
	"strconv"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(conf config.Config) string {
	return "postgres://" + conf.UserName + ":" + conf.Password + "@" + conf.Host + ":" + strconv.Itoa(conf.Port) + "/" + conf.DbName + "?sslmode=disable"
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