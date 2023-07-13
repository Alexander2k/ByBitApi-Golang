package spot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MarketData interface {
	GetAllSymbols() (AllSymbolsResponse, error)
	GetOrderBook(symbol, limit string) (OrderBookResponse, error)
	GetMergedOrderBook(symbol, scale, limit string) (OrderBookResponse, error)
	GetPublicTradingRecords(symbol string) (TradingRecordsResponse, error)
	GetTickers(category, symbol string) (GetTickersResponse, error)
	GetLastTradePrice(symbol string) (LastTradePriceResponse, error)
	GetBookTicker(symbol string) (BookTickerResponse, error)
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

func (s *Spot) GetOrderBook(symbol, limit string) (OrderBookResponse, error) {
	var orderBook OrderBookResponse

	point := fmt.Sprintf("/spot/v3/public/quote/depth?symbol=%s&limit=%s", symbol, limit)

	response, err := s.Signer.Get(http.MethodGet, point)
	if err != nil {
		return OrderBookResponse{}, err
	}

	_ = json.Unmarshal(response, &orderBook)

	return orderBook, err
}

func (s *Spot) GetMergedOrderBook(symbol, scale, limit string) (OrderBookResponse, error) {
	var orderBook OrderBookResponse

	point := fmt.Sprintf("/spot/v3/public/quote/depth?symbol=%s&limit=%s&scale=%s", symbol, limit, scale)

	response, err := s.Signer.Get(http.MethodGet, point)
	if err != nil {
		return OrderBookResponse{}, err
	}

	_ = json.Unmarshal(response, &orderBook)

	return orderBook, err
}

func (s *Spot) GetPublicTradingRecords(symbol string) (TradingRecordsResponse, error) {
	var tr TradingRecordsResponse
	point := fmt.Sprintf("/spot/v3/public/quote/trades?symbol=%s", symbol)

	response, err := s.Signer.Get(http.MethodGet, point)
	if err != nil {
		return TradingRecordsResponse{}, nil
	}
	_ = json.Unmarshal(response, &tr)

	return tr, err
}

// GetTickers Query for the latest price snapshot, best bid/ask price, and trading volume in the last 24 hours.
// Category is Required
// if not need symbol use ""
func (s *Spot) GetTickers(category, symbol string) (GetTickersResponse, error) {
	var t GetTickersResponse
	var point string
	if symbol != "" {
		point = fmt.Sprintf("/v5/market/tickers?category=%s&symbol=%s", category, symbol)
	} else {
		point = fmt.Sprintf("/v5/market/tickers?category=%s", category)
	}

	b, err := s.Signer.Get(http.MethodGet, point)
	if err != nil {
		return GetTickersResponse{}, nil
	}

	err = json.Unmarshal(b, &t)
	if err != nil {
		return GetTickersResponse{}, err
	}

	return t, err
}

func (s *Spot) GetLastTradePrice(symbol string) (LastTradePriceResponse, error) {
	var lt LastTradePriceResponse
	var point string
	if symbol != "" {
		point = fmt.Sprintf("/spot/v3/public/quote/ticker/price?symbol=%s", symbol)
	} else {
		point = "/spot/v3/public/quote/ticker/price"
	}

	bytes, err := s.Signer.Get(http.MethodGet, point)
	if err != nil {
		return LastTradePriceResponse{}, err
	}

	err = json.Unmarshal(bytes, &lt)
	if err != nil {
		return LastTradePriceResponse{}, err
	}

	return lt, err
}

func (s *Spot) GetBookTicker(symbol string) (BookTickerResponse, error) {
	var point string
	var bt BookTickerResponse
	if symbol != "" {
		point = fmt.Sprintf("/spot/v3/public/quote/ticker/bookTicker?symbol=%s", symbol)
	} else {
		point = "/spot/v3/public/quote/ticker/bookTicker"
	}

	bytes, err := s.Signer.Get(http.MethodGet, point)
	if err != nil {
		return BookTickerResponse{}, err
	}

	err = json.Unmarshal(bytes, &bt)
	if err != nil {
		return BookTickerResponse{}, err
	}
	return bt, err
}
