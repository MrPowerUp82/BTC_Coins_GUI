package main

import (
	"bitcoin/gui"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://blockchain.info/ticker")

	if err != nil {
		log.Fatal(err.Error())
	}

	resData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var resObj map[string]gui.Coin

	json.Unmarshal(resData, &resObj)

	var resArray []gui.Coin

	for _, coin := range resObj {
		// fmt.Printf("Moeda: %v, Compra: %v, Venda: %v\n", coin.Symbol, coin.Buy, coin.Sell)
		resArray = append(resArray, coin)
	}

	gui.App(resArray)
}
