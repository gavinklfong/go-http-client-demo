package main

import (
	"fmt"

	"github.com/gavinklfong/go-http-client-demo/forex"
)

const serverPort = 8080

func main() {
	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	client := forex.NewForexApiClient(requestURL)
	rates := client.GetLatestRates()

	fmt.Printf("rates: %v\n", rates)
}
