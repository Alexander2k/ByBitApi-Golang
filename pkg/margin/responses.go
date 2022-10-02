package margin

type TickersResponse struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  []struct {
		Symbol                 string `json:"symbol"`
		BidPrice               string `json:"bid_price"`
		AskPrice               string `json:"ask_price"`
		LastPrice              string `json:"last_price"`
		LastTickDirection      string `json:"last_tick_direction"`
		PrevPrice24H           string `json:"prev_price_24h"`
		Price24HPcnt           string `json:"price_24h_pcnt"`
		HighPrice24H           string `json:"high_price_24h"`
		LowPrice24H            string `json:"low_price_24h"`
		PrevPrice1H            string `json:"prev_price_1h"`
		Price1HPcnt            string `json:"price_1h_pcnt"`
		MarkPrice              string `json:"mark_price"`
		IndexPrice             string `json:"index_price"`
		OpenInterest           int    `json:"open_interest"`
		OpenValue              string `json:"open_value"`
		TotalTurnover          string `json:"total_turnover"`
		Turnover24H            string `json:"turnover_24h"`
		TotalVolume            int64  `json:"total_volume"`
		Volume24H              int    `json:"volume_24h"`
		FundingRate            string `json:"funding_rate"`
		PredictedFundingRate   string `json:"predicted_funding_rate"`
		NextFundingTime        string `json:"next_funding_time"`
		CountdownHour          int    `json:"countdown_hour"`
		DeliveryFeeRate        string `json:"delivery_fee_rate"`
		PredictedDeliveryPrice string `json:"predicted_delivery_price"`
		DeliveryTime           string `json:"delivery_time"`
	} `json:"result"`
	TimeNow string `json:"time_now"`
}
