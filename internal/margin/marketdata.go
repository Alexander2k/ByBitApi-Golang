package margin

import (
	"encoding/json"
	"net/http"
)

type MarketData interface {
	GetTicker() (TickersResponse, error)
}

func (m *Margin) GetTicker() (TickersResponse, error) {
	var t TickersResponse
	bytes, err := m.Signer.Get(http.MethodGet, "/v2/public/tickers")
	if err != nil {
		return TickersResponse{}, nil
	}

	_ = json.Unmarshal(bytes, &t)
	return t, err
}
