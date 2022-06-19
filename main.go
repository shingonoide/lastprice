package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type PriceResult struct {
	Result struct {
		Price float64 `json:"price"`
	} `json:"result"`
}

func main() {

	var pair string = "btcusdt"
	var p PriceResult
	argLength := len(os.Args[1:])
	if argLength > 0 {
		pair = os.Args[1]
	}

	resp, err := http.Get("https://api.cryptowat.ch/markets/binance/" + pair + "/price")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &p)

	log.Printf("%s: %f", pair, p.Result.Price)

}
