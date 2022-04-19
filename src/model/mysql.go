package model

import (
	"fmt"
	"github.com/FlyInThesky10/antiNCP-backend/config"
	"github.com/FlyInThesky10/antiNCP-backend/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s",
			config.C.MySQL.Username, config.C.MySQL.Password, config.C.MySQL.Addr, config.C.MySQL.Db))
	if err != nil {
		log.Logger.Println("Cannot connect to MySQL.")
		log.Logger.Panic(err)
		return
	}
	Db = database
	log.Logger.Println("Successfully connected to MySQL.")
}
