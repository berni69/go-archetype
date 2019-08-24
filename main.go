package main

import (
	"net/http"

	"github.com/berni69/go-archetype/consul"
	"github.com/berni69/go-archetype/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func helloWorld(resp http.ResponseWriter, req *http.Request) {

	log.Debug("Hello world")
	resp.Write([]byte("Hello world"))
}

func discoverAirport(resp http.ResponseWriter, req *http.Request) {
	url, err := consul.LookupService("airport")
	if err != nil {
		resp.Write([]byte(err.Error()))
	}

	log.Debug(url)
	resp.Write([]byte(url))
}

func getConfig() {

	err := consul.LoadConsulConfig(utils.GetEnv("CONSUL_KV", ""), &Configuration)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Info("Starting Server")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Logging options
	utils.InitLogger()
	/** Loading Config **/
	c := cron.New()
	getConfig()
	c.AddFunc("0 * * * * *", func() { getConfig() }) //Retrieve the config every minut
	c.Start()

	/* Starting new Router */
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld).Methods("GET")
	router.HandleFunc("/discover", discoverAirport).Methods("GET")
	addressPort := utils.GetEnv("ADDRESS_PORT", ":8000")
	log.Fatal(http.ListenAndServe(addressPort, router))

}
