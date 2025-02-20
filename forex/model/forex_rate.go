package model

import (
	"encoding/json"
	"fmt"
	"time"
)

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
