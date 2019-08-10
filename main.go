package main

import (

	//"encoding/json"
	"encoding/json"
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
	//var generic map[string]interface{}
	var generic []interface{}
	var err = utils.GetJson("https://api.github.com/users/hadley/orgs", &generic)
	if err != nil {
		fmt.Println(err)
		return
	}
	j, err := json.Marshal(generic)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Write(j)

}

func main() {
	fmt.Println("Starting Server")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addressPort := utils.GetEnv("ADDRESS_PORT", ":8000")
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld).Methods("GET")

	log.Fatal(http.ListenAndServe(addressPort, router))

	consul.Create_tls_config("", "", "")

}
