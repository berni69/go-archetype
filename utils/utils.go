package utils

import (
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
