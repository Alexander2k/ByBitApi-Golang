package derivatives

import (
	"encoding/json"
	"fmt"
)

type MarketData interface {
	GetTicker(category string) (TickerResponse, error)
	GetTickerOption(category, symbol string) OptionResponse
	GetInstrumentInfo(category, symbol, limit string) InstrumentInfoResponse
}

func (d *Derivatives) GetTicker(category string) (TickerResponse, error) {
	var t TickerResponse
	switch {
	case category == INVERSE:

		point := fmt.Sprintf("/derivatives/v3/public/tickers?category=%s", category)
		bytes, err := d.Signer.Get(GET, point)
		if err != nil {
			return TickerResponse{}, err
		}
		_ = json.Unmarshal(bytes, &t)
		return t, err
	case category == LINEAR:

		point := fmt.Sprintf("/derivatives/v3/public/tickers?category=%s", category)
		bytes, err := d.Signer.Get(GET, point)
		if err != nil {
			return TickerResponse{}, err
		}
		_ = json.Unmarshal(bytes, &t)
		return t, err
	}
	return t, nil

}

func (d *Derivatives) GetTickerOption(category, symbol string) OptionResponse {
	var o OptionResponse
	point := fmt.Sprintf("/derivatives/v3/public/tickers?category=%s&symbol=%s", category, symbol)
	bytes, err := d.Signer.Get(GET, point)
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(bytes, &o)
	return o
}

func (d *Derivatives) GetInstrumentInfo(category, symbol, limit string) InstrumentInfoResponse {
	var inst InstrumentInfoResponse

	switch {
	case category == LINEAR:
		point := fmt.Sprintf("/derivatives/v3/public/instruments-info?category=%s&symbol=%s&limit=%s", category, symbol, limit)
		bytes, err := d.Signer.Get("GET", point)
		if err != nil {
			return InstrumentInfoResponse{}
		}

		_ = json.Unmarshal(bytes, &inst)
		return inst
	case category == INVERSE:
		point := fmt.Sprintf("/derivatives/v3/public/instruments-info?category=%s&limit=%s", category, limit)
		bytes, err := d.Signer.Get("GET", point)
		if err != nil {
			return InstrumentInfoResponse{}
		}

		_ = json.Unmarshal(bytes, &inst)
		return inst
	}
	return InstrumentInfoResponse{}
}
