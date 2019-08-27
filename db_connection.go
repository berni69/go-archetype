package main

import (
	sql2 "database/sql"
	"fmt"

	sql "go.elastic.co/apm/module/apmsql"

	//_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

// DBPool Global object, it's a connection pool usable after call CreatePool
var DBPool *sql2.DB

// MAX_IDLE_CONNEXIONS constant:  maximum number of idle connections at same time
const MAX_IDLE_CONNEXIONS = 2

// MAX_OPEN_CONNEXIONS constant: maximum number of opened connections by the pool
const MAX_OPEN_CONNEXIONS = 5

// CreatePool This function is used to initialize the DB Pool, Configuration and Secrets must
// be filled before call this function
func CreatePool() {
	log.Info("Creating PG Pool")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
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
	log.Debug("Created PG Pool")

}
