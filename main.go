package main

import (
	"encoding/json"
	"net/http"

	"github.com/berni69/go-archetype/consul"
	"github.com/berni69/go-archetype/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func helloWorld(resp http.ResponseWriter, req *http.Request) {

	log.Debug("Hello world")
	resp.Write([]byte("Hello world"))
	var generic []interface{}
	var err = utils.GetJson("https://api.github.com/users/hadley/orgs", &generic)
	if err != nil {
		log.Debug(err)
		return
	}
	j, err := json.Marshal(generic)
	if err != nil {
		log.Debug(err)
		return
	}
	resp.Write(j)

}

func main() {
	log.Info("Starting Server")
	var config configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	utils.InitLogger()

	addressPort := utils.GetEnv("ADDRESS_PORT", ":8000")
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld).Methods("GET")

	log.Debug(consul.LookupService("airport"))
	err = consul.LoadConsulConfig(utils.GetEnv("CONSUL_KV", ""), &config)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(addressPort, router))
}
