package derivatives

type OrderBookParams struct {
	Category string `json:"category,omitempty"`
	Symbol   string `json:"symbol"`
	Limit    string `json:"limit,omitempty"`
}

type GetInstrumentInfoParams struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol,omitempty"`
	Limit    string `json:"limit,omitempty"`
	Cursor   string `json:"cursor,omitempty"`
}

type GetTickerParams struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol,omitempty"`
}

type KlineParams struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
	Interval string `json:"interval"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Limit    string `json:"limit,omitempty"`
}

type FoundingRateHistoryParams struct {
	Category  string `json:"category"`
	Symbol    string `json:"symbol"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Limit     string `json:"limit,omitempty"`
}

type RiskLimitParams struct {
	Category string `json:"category"`
	Symbol   string `json:"symbol"`
}

type OptionDeliveryPriceParams struct {
	Category  string `json:"category,omitempty"`
	Symbol    string `json:"symbol,omitempty"`
	BaseCoin  string `json:"base_coin,omitempty"`
	Direction string `json:"direction,omitempty"`
	Limit     string `json:"limit,omitempty"`
	Cursor    string `json:"cursor,omitempty"`
}

type PublicTradingHistoryParams struct {
	Category   string `json:"category"`
	Symbol     string `json:"symbol"`
	BaseCoin   string `json:"base_coin,omitempty"`
	OptionType string `json:"option_type,omitempty"`
	Limit      string `json:"limit,omitempty"`
}

type OpenInterest struct {
	Category  string `json:"category"`
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Limit     string `json:"limit,omitempty"`
}
