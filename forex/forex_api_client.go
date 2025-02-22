package forex

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gavinklfong/go-http-client-demo/forex/model"
)

type ForexApiClient struct {
	url string
}

func NewForexApiClient(url string) *ForexApiClient {
	return &ForexApiClient{url: url}
}

func (c *ForexApiClient) GetLatestRates() ([]model.ForexRateResponse, error) {
	requestURL := fmt.Sprintf("%s/rates/latest", c.url)
	res, err := http.Get(requestURL)
	if err != nil {
		log.Fatalf("error making http request: %s\n", err)
		return nil, err
	}

	fmt.Printf("status: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading response body: %s\n", err)
		return nil, err
	}
	fmt.Printf("body: %s\n", resBody)

	var rates []model.ForexRateResponse
	err = json.Unmarshal(resBody, &rates)
	if err != nil {
		log.Fatalf("error parsing response body: %s", err)
		return nil, err
	}

	return rates, nil
}

// func (c *ForexApiClient) BookRate(booking *model.ForexRateBookingRequest) *model.ForexRateBookingResponse {

// }
