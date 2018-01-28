package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func myController(w http.ResponseWriter, r *http.Request) {

	var body = make([]Req, 0)
	resp, err := http.Get("https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=3")

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return

	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&body)
}