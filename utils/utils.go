package utils

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// GetEnv allows to retrieve a env variable or get the default value if not exists
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetJSON This functions wrap a simple HTTP GET request
func GetJSON(url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// Jsonify Helper used to convert a DB Output  to JSON (bytes) ready to be served by MUX
func Jsonify(rows *sql.Rows) ([]byte, error) {
	cols, _ := rows.Columns()
	// 2. Iterate
	list := make([]map[string]interface{}, 0)
	for rows.Next() {
		vals := make([]interface{}, len(cols))
		for i := range cols {
			// Previously you assigned vals[i] a pointer to a column name cols[i].
			// This meant that everytime you did rows.Scan(vals),
			// rows.Scan would see pointers to cols and modify them
			// Since cols are the same for all rows, they shouldn't be modified.

			// Here we assign a pointer to an empty string to vals[i],
			// so rows.Scan can fill it.
			var s string
			vals[i] = &s

			// This is effectively like saying:
			// var string1, string2 string
			// rows.Scan(&string1, &string2)
			// Except the above only scans two string columns
			// and we allow as many string columns as the query returned us — len(cols).
		}

		err := rows.Scan(vals...)

		// Don't forget to check errors.
		if err != nil {
			log.Fatal(err)
		}

		// Make a new map before appending it.
		// Remember maps aren't copied by value, so if we declared
		// the map m outside of the rows.Next() loop, we would be appending
		// and modifying the same map for each row, so all rows in list would look the same.
		m := make(map[string]interface{})
		for i, val := range vals {
			m[cols[i]] = val
		}
		list = append(list, m)
	}
	return json.MarshalIndent(list, "", "\t")

}

// GetRandomInt generate a random integer between 0 and max
func GetRandomInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - 0 + 1)
}

// InitLogger Used to init the common config of the logger in the entire project
func InitLogger() {
	lvl := GetEnv("LOG_LEVEL", "debug")
	// parse string, this is built-in feature of logrus
	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.DebugLevel
	}
	// set global log level
	log.SetLevel(ll)
	log.SetFormatter(&log.JSONFormatter{})

}
