package derivatives

type TickerResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string `json:"category"`
		List     []struct {
			Symbol                 string `json:"symbol"`
			BidPrice               string `json:"bidPrice"`
			AskPrice               string `json:"askPrice"`
			LastPrice              string `json:"lastPrice"`
			LastTickDirection      string `json:"lastTickDirection"`
			PrevPrice24H           string `json:"prevPrice24h"`
			Price24HPcnt           string `json:"price24hPcnt"`
			HighPrice24H           string `json:"highPrice24h"`
			LowPrice24H            string `json:"lowPrice24h"`
			PrevPrice1H            string `json:"prevPrice1h"`
			MarkPrice              string `json:"markPrice"`
			IndexPrice             string `json:"indexPrice"`
			OpenInterest           string `json:"openInterest"`
			Turnover24H            string `json:"turnover24h"`
			Volume24H              string `json:"volume24h"`
			FundingRate            string `json:"fundingRate"`
			NextFundingTime        string `json:"nextFundingTime"`
			PredictedDeliveryPrice string `json:"predictedDeliveryPrice"`
			BasisRate              string `json:"basisRate"`
			DeliveryFeeRate        string `json:"deliveryFeeRate"`
			DeliveryTime           string `json:"deliveryTime"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type OptionResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category               string `json:"category"`
		Symbol                 string `json:"symbol"`
		BidPrice               string `json:"bidPrice"`
		BidSize                string `json:"bidSize"`
		BidIv                  string `json:"bidIv"`
		AskPrice               string `json:"askPrice"`
		AskSize                string `json:"askSize"`
		AskIv                  string `json:"askIv"`
		LastPrice              string `json:"lastPrice"`
		HighPrice24H           string `json:"highPrice24h"`
		LowPrice24H            string `json:"lowPrice24h"`
		MarkPrice              string `json:"markPrice"`
		IndexPrice             string `json:"indexPrice"`
		MarkPriceIv            string `json:"markPriceIv"`
		UnderlyingPrice        string `json:"underlyingPrice"`
		OpenInterest           string `json:"openInterest"`
		Turnover24H            string `json:"turnover24h"`
		Volume24H              string `json:"volume24h"`
		TotalVolume            string `json:"totalVolume"`
		TotalTurnover          string `json:"totalTurnover"`
		Delta                  string `json:"delta"`
		Gamma                  string `json:"gamma"`
		Vega                   string `json:"vega"`
		Theta                  string `json:"theta"`
		PredictedDeliveryPrice string `json:"predictedDeliveryPrice"`
		Change24H              string `json:"change24h"`
	} `json:"result"`
	Time int64 `json:"time"`
}

type InstrumentInfoResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string `json:"category"`
		List     []struct {
			Symbol          string `json:"symbol"`
			ContractType    string `json:"contractType"`
			Status          string `json:"status"`
			BaseCoin        string `json:"baseCoin"`
			QuoteCoin       string `json:"quoteCoin"`
			LaunchTime      string `json:"launchTime"`
			DeliveryTime    string `json:"deliveryTime"`
			DeliveryFeeRate string `json:"deliveryFeeRate"`
			PriceScale      string `json:"priceScale"`
			LeverageFilter  struct {
				MinLeverage  string `json:"minLeverage"`
				MaxLeverage  string `json:"maxLeverage"`
				LeverageStep string `json:"leverageStep"`
			} `json:"leverageFilter"`
			PriceFilter struct {
				MinPrice string `json:"minPrice"`
				MaxPrice string `json:"maxPrice"`
				TickSize string `json:"tickSize"`
			} `json:"priceFilter"`
			LotSizeFilter struct {
				MaxTradingQty string `json:"maxTradingQty"`
				MinTradingQty string `json:"minTradingQty"`
				QtyStep       string `json:"qtyStep"`
			} `json:"lotSizeFilter"`
		} `json:"list"`
		NextPageCursor string `json:"nextPageCursor"`
	} `json:"result"`
	Time int64 `json:"time"`
}

type PriceKlineResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Symbol   string     `json:"symbol"`
		Category string     `json:"category"`
		List     [][]string `json:"list"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type OrderBookResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Symbol string     `json:"s"`
		Bid    [][]string `json:"b"`
		Ask    [][]string `json:"a"`
		Ts     int64      `json:"ts"`
		U      int        `json:"u"`
	} `json:"result"`
}

type KlineResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string   `json:"category"`
		Symbol   string   `json:"symbol"`
		Interval string   `json:"interval"`
		List     []string `json:"list"`
	} `json:"result"`
}

type OptionInfoResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		ResultTotalSize int    `json:"resultTotalSize"`
		Cursor          string `json:"cursor"`
		DataList        []struct {
			Category        string `json:"category"`
			Symbol          string `json:"symbol"`
			Status          string `json:"status"`
			BaseCoin        string `json:"baseCoin"`
			QuoteCoin       string `json:"quoteCoin"`
			SettleCoin      string `json:"settleCoin"`
			OptionsType     string `json:"optionsType"`
			LaunchTime      string `json:"launchTime"`
			DeliveryTime    string `json:"deliveryTime"`
			DeliveryFeeRate string `json:"deliveryFeeRate"`
			PriceFilter     struct {
				MinPrice string `json:"minPrice"`
				MaxPrice string `json:"maxPrice"`
				TickSize string `json:"tickSize"`
			} `json:"priceFilter"`
			LotSizeFilter struct {
				MaxOrderQty string `json:"maxOrderQty"`
				MinOrderQty string `json:"minOrderQty"`
				QtyStep     string `json:"qtyStep"`
			} `json:"lotSizeFilter"`
		} `json:"dataList"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type FoundingRateHistoryResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string `json:"category"`
		List     []struct {
			Symbol               string `json:"symbol"`
			FundingRate          string `json:"fundingRate"`
			FundingRateTimestamp string `json:"fundingRateTimestamp"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}

type RiskLimitResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string `json:"category"`
		List     []struct {
			Id             int      `json:"id"`
			Symbol         string   `json:"symbol"`
			Limit          string   `json:"limit"`
			MaintainMargin string   `json:"maintainMargin"`
			InitialMargin  string   `json:"initialMargin"`
			Section        []string `json:"section"`
			IsLowestRisk   int      `json:"isLowestRisk"`
			MaxLeverage    string   `json:"maxLeverage"`
		} `json:"list"`
	} `json:"result"`
	Time int64 `json:"time"`
}
