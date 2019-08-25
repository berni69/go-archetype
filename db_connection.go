package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var DBPool *sql.DB

const MAX_IDLE_CONNEXIONS = 2
const MAX_OPEN_CONNEXIONS = 5

func CreatePool() {
	log.Info("Creating PG Pool")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		Configuration.Database.User, Secrets.DatabasePassword,
		Configuration.Database.Host, Configuration.Database.Port,
		Configuration.Database.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(MAX_OPEN_CONNEXIONS)
	db.SetMaxIdleConns(MAX_IDLE_CONNEXIONS)
	DBPool = db // Export as global variable

}
