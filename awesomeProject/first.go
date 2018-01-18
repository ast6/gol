package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Req struct {
	username string `json:"username"`
    password string `json:"password"`
}

func main() {
git
	router := mux.NewRouter()
	router.HandleFunc("/", myController)
	fmt.Println("1111")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func myController(w http.ResponseWriter, r *http.Request) {
	req := Req{}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("welcome"))
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(req)

}
