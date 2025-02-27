package model

import "testing"

func TestForexRateResponseUnmarshalJSON(t *testing.T) {

	// Timestamp       time.Time
	// BaseCurrency    string
	// CounterCurrency string
	// BuyRate         float32
	// SellRate        float32
	// Spread          float32

	json := `{
        "timestamp":"2025-02-22T01:16:07.427722",
        "baseCurrency":"GBP",
        "counterCurrency":"JPY",
        "buyRate":147.9909,
        "sellRate":147.9901,
        "spread":8.0E-4
    }`

	response := ForexRateResponse{}

	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
