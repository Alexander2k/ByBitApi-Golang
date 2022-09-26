package derivatives

const (
	LINEAR  = "linear"
	INVERSE = "inverse"
	OPTION  = "option"
	GET     = "GET"
	POST    = "POST"
)

// Market Data Endpoints
const (
	GetOrderBook                  = "/derivatives/v3/public/order-book/L2?"
	GetKline                      = "/derivatives/v3/public/kline?"
	GetLatestInformationForSymbol = "/derivatives/v3/public/tickers?"
	GetInstrumentInfo             = "/derivatives/v3/public/instruments-info?"
	GetMarkPriceKline             = "/derivatives/v3/public/mark-price-kline?"
	GetIndexPriceKline            = "/derivatives/v3/public/index-price-kline?"
	GetFundingRateHistory         = "/derivatives/v3/public/funding/history-funding-rate?"
	GetRiskLimit                  = "/derivatives/v3/public/risk-limit/list?"
	GetOptionDeliveryPrice        = "/derivatives/v3/public/delivery-price?"
	GetPublicTradingHistory       = "/derivatives/v3/public/recent-trade?"
	GetOpenInterest               = "/derivatives/v3/public/open-interest?"
)
