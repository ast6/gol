package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", myController)
	router.HandleFunc("/exchange", calculateExchange).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
