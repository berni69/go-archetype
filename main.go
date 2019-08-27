package main

import (
	"net/http"

	"github.com/berni69/go-archetype/consul"
	"github.com/berni69/go-archetype/utils"
	"github.com/berni69/go-archetype/vault"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgorilla"
	"go.elastic.co/apm/transport"
)

func helloWorld(resp http.ResponseWriter, req *http.Request) {

	log.Debug("Hello world")
	resp.Write([]byte("Hello world"))
}

func records(resp http.ResponseWriter, req *http.Request) {

	log.Debug("Records")
	rows, err := DBPool.QueryContext(req.Context(), `SELECT id, data2 FROM test`)

	if err != nil {
		log.Error(err)
		return
	}
	if err, ok := err.(*pq.Error); ok {
		log.Error("pq error:", err.Code.Name())
		return
	}
	if rows == nil {
		resp.Write([]byte("Empty rows"))
		return
	}
	js, err := utils.Jsonify(rows)
	if err != nil {
		log.Error(err)
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	resp.Write(js)
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

	consulKv := utils.GetEnv("CONSUL_KV", "")
	if consulKv != "" {
		err := consul.LoadConsulConfig(consulKv, &Configuration)
		if err != nil {
			log.Fatal(err)
		}
	}
	vaultKv := utils.GetEnv("VAULT_KV", "")
	if vaultKv != "" {
		err := vault.LoadVaultConfig(vaultKv, &Secrets)
		if err != nil {
			log.Fatal(err)

		}
	}
}

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Warning("Error loading .env file")
	}

	// Logging options
	utils.InitLogger()

	/** Loading Config **/
	c := cron.New()
	getConfig()
	c.AddFunc("0 * * * * *", func() { getConfig() }) //Retrieve the config every minut
	c.Start()

	//APM Init
	transport.InitDefault()

	// Connection Pool
	CreatePool()

}

func main() {
	log.Info("Starting Server")

	/* Starting new Router */
	router := mux.NewRouter()
	tracer, err := apm.NewTracer("", "")
	if err != nil {
		log.Fatal(err)
	}
	apmgorilla.Instrument(router, apmgorilla.WithTracer(tracer))
	router.HandleFunc("/hello", helloWorld).Methods("GET")
	router.HandleFunc("/discover", discoverAirport).Methods("GET")
	router.HandleFunc("/records", records).Methods("GET")

	addressPort := utils.GetEnv("ADDRESS_PORT", ":8000")
	log.Fatal(http.ListenAndServe(addressPort, router))

}
