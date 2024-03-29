package spot

type RetExtInfo struct {
}
type GetTickersResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string `json:"category"`
		List     []struct {
			Symbol        string `json:"symbol"`
			Bid1Price     string `json:"bid1Price"`
			Bid1Size      string `json:"bid1Size"`
			Ask1Price     string `json:"ask1Price"`
			Ask1Size      string `json:"ask1Size"`
			LastPrice     string `json:"lastPrice"`
			PrevPrice24H  string `json:"prevPrice24h"`
			Price24HPcnt  string `json:"price24hPcnt"`
			HighPrice24H  string `json:"highPrice24h"`
			LowPrice24H   string `json:"lowPrice24h"`
			Turnover24H   string `json:"turnover24h"`
			Volume24H     string `json:"volume24h"`
			UsdIndexPrice string `json:"usdIndexPrice"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type AllSymbolsResponse struct {
	RetCode    int           `json:"retCode"`
	RetMsg     string        `json:"retMsg"`
	Result     SymbolsResult `json:"result"`
	RetExtMap  RetExtMap     `json:"retExtMap"`
	RetExtInfo RetExtInfo    `json:"retExtInfo"`
	Time       int64         `json:"time"`
}
type Symbol struct {
	Name              string `json:"name"`
	Alias             string `json:"alias"`
	BaseCoin          string `json:"baseCoin"`
	QuoteCoin         string `json:"quoteCoin"`
	BasePrecision     string `json:"basePrecision"`
	QuotePrecision    string `json:"quotePrecision"`
	MinTradeQty       string `json:"minTradeQty"`
	MinTradeAmt       string `json:"minTradeAmt"`
	MaxTradeQty       string `json:"maxTradeQty"`
	MaxTradeAmt       string `json:"maxTradeAmt"`
	MinPricePrecision string `json:"minPricePrecision"`
	Category          string `json:"category"`
	ShowStatus        string `json:"showStatus"`
	Innovation        string `json:"innovation"`
}
type SymbolsResult struct {
	Symbols []Symbol `json:"list"`
}
type RetExtMap struct {
}

type OrderBookResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Time int64      `json:"time"`
		Bids [][]string `json:"bids"`
		Asks [][]string `json:"asks"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type TradingRecordsResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			Price        string `json:"price"`
			Time         int64  `json:"time"`
			Qty          string `json:"qty"`
			IsBuyerMaker int    `json:"isBuyerMaker"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type TickerResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			Timestamp   int64  `json:"t"`
			Symbol      string `json:"s"`
			Last        string `json:"lp"`
			High        string `json:"h"`
			Low         string `json:"l"`
			Open        string `json:"o"`
			BestBid     string `json:"bp"`
			BestAsk     string `json:"ap"`
			Volume      string `json:"v"`
			QuoteVolume string `json:"qv"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type LastTradePriceResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			Symbol string `json:"symbol"`
			Price  string `json:"price"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}

type BookTickerResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		List []struct {
			Symbol   string `json:"symbol"`
			BidPrice string `json:"bidPrice"`
			BidQty   string `json:"bidQty"`
			AskPrice string `json:"askPrice"`
			AskQty   string `json:"askQty"`
			Time     int64  `json:"time"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
	} `json:"retExtInfo"`
	Time int64 `json:"time"`
}
