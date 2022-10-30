package db

import (
	"database/sql"
	"fmt"
	"time"
	"vandyahmad24/privy/app/util"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {

	fmt.Println("Trying to connect database :" + util.GetEnvVariable("MYSQL_DBNAME"))
	fmt.Println("Trying to connect MYSQL_HOST :" + util.GetEnvVariable("MYSQL_HOST"))

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		util.GetEnvVariable("MYSQL_USER"),
		util.GetEnvVariable("MYSQL_PASSWORD"),
		util.GetEnvVariable("MYSQL_HOST"),
		util.GetEnvVariable("MYSQL_PORT"),
		util.GetEnvVariable("MYSQL_DBNAME"),
	))

	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(6 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
