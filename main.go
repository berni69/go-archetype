package main

import (

	//"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/berni69/go-archetype/consul"
	"github.com/berni69/go-archetype/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func helloWorld(resp http.ResponseWriter, req *http.Request) {

	fmt.Println("Hello world")
	resp.Write([]byte("Hello world"))
}

func main() {
	fmt.Println("Starting Server")

	consul.Create_tls_config("", "", "")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addressPort := utils.GetEnv("ADDRESS_PORT", ":8000")
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(addressPort, router))

}
