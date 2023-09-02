package main

import (
	"log"

	"github.com/go-pg/pg/v10"
)

func main() {

	//Options User GreePlum
	dns := `postgres://User:Password@127.0.0.1:9999/Database?sslmode=disable&options=-c%20gp_session_role%3dutility`
	log.Println(pg.ParseURL(dns))

	db := pg.Connect(&pg.Options{
		User:    "postgres",
		Options: "-c gp_session_role=utility",
	})
	log.Println(db.Options())
	defer db.Close()
}
