package main

import (
	"fmt"
	"log"

	"github.com/gavinklfong/go-http-client-demo/forex"
)

const serverPort = 8080

func main() {
	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	client := forex.NewForexApiClient(requestURL)
	rates, err := client.GetLatestRates()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("rates: %v\n", rates)
}
