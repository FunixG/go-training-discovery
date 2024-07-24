package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Crypto struct {
	Name  string `json:"symbol"`
	Value string `json:"price"`
}

func main() {
	fmt.Println("== Prix actuels des cryptomonnaies ==")

	var currencies = [2]string{"ETH", "BTC"}

	for i := 0; i < len(currencies); i++ {
		printCurrency(makeRequest(currencies[i]))
	}
}

func printCurrency(crypto Crypto) {
	fmt.Printf("Valeur du %v: %v\n", crypto.Name, crypto.Value)
}

func makeRequest(crypto string) Crypto {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=" + crypto + "EUR")

	if err != nil {
		log.Fatalf("Requête de récéption de la valeur de la cryptomonnaie %v. Erreur: %v", crypto, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Impossible de lire la réponse de la requête de la cryptomonnaie %v. Erreur: %s", crypto, err)
	}

	var cryptoPrice Crypto
	err = json.Unmarshal(body, &cryptoPrice)
	if err != nil {
		log.Fatalf("Erreur lors du décodage JSOn de la crypto %v. Erreur: %v", crypto, err)
	}

	cryptoPrice.Name = crypto
	return cryptoPrice
}
