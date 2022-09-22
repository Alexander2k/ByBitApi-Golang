package spot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MarketData interface {
	GetAllSymbols() (AllSymbolsResponse, error)
	GetTicker() (TickerResponse, error)
}

func (s *Spot) GetAllSymbols() (AllSymbolsResponse, error) {
	var symbols AllSymbolsResponse

	point := fmt.Sprintf("/spot/v3/public/symbols")

	response, err := s.Signer.Get(http.MethodGet, point)
	if err != nil {
		return AllSymbolsResponse{}, err
	}

	_ = json.Unmarshal(response, &symbols)

	return symbols, err
}

func (s *Spot) GetTicker() (TickerResponse, error) {
	var t TickerResponse

	bytes, err := s.Signer.Get(http.MethodGet, "/spot/v3/public/quote/ticker/24hr")
	if err != nil {
		return TickerResponse{}, err
	}

	_ = json.Unmarshal(bytes, &t)
	return t, err
}
