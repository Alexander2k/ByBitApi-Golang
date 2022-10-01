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

// Order endpoints
const (
	PlaceOrder         = "/unified/v3/private/order/create?"
	ReplaceOrder       = "/unified/v3/private/order/replace?"
	CancelOrder        = "/unified/v3/private/order/cancel?"
	GetOpenOrders      = "/unified/v3/private/order/unfilled-orders?"
	GetOrders          = "/unified/v3/private/order/list?"
	BatchPlaceOrders   = "/unified/v3/private/order/create-batch?"
	BatchReplaceOrders = "/unified/v3/private/order/replace-batch?"
	BatchCancelOrders  = "/unified/v3/private/order/cancel-batch?"
	CancelAllOrders    = "/unified/v3/private/order/cancel-all?"
)

// Position endpoints
const (
	GetPosition                   = "/unified/v3/private/position/list?"
	ModifyLeverage                = "/unified/v3/private/position/set-leverage?"
	SwitchTPSLMode                = "/unified/v3/private/position/tpsl/switch-mode?"
	SetRiskLimit                  = "/unified/v3/private/position/set-risk-limit?"
	SetTPSL                       = "/unified/v3/private/position/trading-stop?"
	GetSevenDayTradingHistory     = "/unified/v3/private/execution/list?"
	GetSettlementHistoryOptions   = "/unified/v3/private/delivery-record?"
	GetSettlementHistoryPerpetual = "/unified/v3/private/settlement-record?"
)

// Account
const (
	GetWalletBalance              = "/unified/v3/private/account/wallet/balance?"
	UpgradeToUnifiedMarginAccount = "/unified/v3/private/account/upgrade-unified-account"
)
