package main

import (
	"fmt"
	"encoding/json"
	"strconv"
	"net/http"
)

func calculateExchange(w http.ResponseWriter, r *http.Request) {
	request := new(CalculateRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var exchangeRates = make([]Req, 0)
	resp, err := http.Get("https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=3")

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&exchangeRates); err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return

	}

	defer r.Body.Close()

	result := new(CalculateResponse)
	result.Amount = request.Amount

	for _, rate := range exchangeRates {

		rateAmount, err := strconv.ParseFloat(rate.Buy, 64)
		if err != nil {
			continue
		}

		if rate.Ccy == "USD" {
			result.Usd = request.Amount / float32(rateAmount)
		}
		if rate.Ccy == "EUR" {
			result.Eur = request.Amount / float32(rateAmount)
		}
		if rate.Ccy == "BTC" {
			result.Btc = request.Amount / float32(rateAmount)
		}
		if rate.Ccy == "RUR" {
			result.Rur = request.Amount / float32(rateAmount)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&result)
}
