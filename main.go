package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const serverPort = 8080

type ForexRate struct {
	Timestamp       time.Time
	BaseCurrency    string
	CounterCurrency string
	BuyRate         float32
	SellRate        float32
	Spread          float32
}

func (f *ForexRate) UnmarshalJSON(data []byte) error {
	type Alias ForexRate
	aux := &struct {
		Timestamp string
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var t, err = time.Parse("2006-01-02T15:04:05", aux.Timestamp)
	if err != nil {
		return err
	}

	fmt.Printf("timestamp in parser: %v\n", t)
	f.Timestamp = t
	return nil
}

func main() {
	requestURL := fmt.Sprintf("http://localhost:%d/rates/latest", serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("status: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("body: %s\n", resBody)

	var rates []ForexRate
	err = json.Unmarshal(resBody, &rates)
	if err != nil {
		fmt.Printf("error parsing response body: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Forex Rates: %v\n", rates)
}
