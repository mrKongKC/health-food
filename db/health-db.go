package db

import (
	"database/sql"
	_ "github.com/proullon/ramsql/driver"
	"log"
)

var Db *sql.DB

func createHeathFood() {
	createTb := `
	CREATE TABLE IF NOT EXISTS HealthFood (
	healthFoodId INT AUTO_INCREMENT,
	title TEXT NOT NULL,
	ingredient TEXT NOT NULL,
	recipe TEXT NOT NULL,
	year INT NOT NULL,
	isApprove BOOLEAN NOT NULL,
	PRIMARY KEY (healthFoodId)
	);
	`
	if _, err := Db.Exec(createTb); err != nil {
		log.Fatal("create table error", err)
	}
	log.Println("Already created DB")

}

func Conn() {
	var err error
	Db, err = sql.Open("ramsql", "Health-Food")
	if err != nil {
		log.Fatal(err)
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	createHeathFood()
}
