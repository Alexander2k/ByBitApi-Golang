package derivatives

type OrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderId     string `json:"orderId"`
		OrderLinkId string `json:"orderLinkId"`
	} `json:"result"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
