package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/bmizerany/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "powerFootballTeams"
)

var db *sql.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	tmpDb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = tmpDb
	fmt.Println("Connecting to db successfully")
}
