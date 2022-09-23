package spot

type OrderResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderId       string `json:"orderId"`
		OrderLinkId   string `json:"orderLinkId"`
		Symbol        string `json:"symbol"`
		CreateTime    string `json:"createTime"`
		OrderPrice    string `json:"orderPrice"`
		OrderQty      string `json:"orderQty"`
		OrderType     string `json:"orderType"`
		Side          string `json:"side"`
		Status        string `json:"status"`
		TimeInForce   string `json:"timeInForce"`
		AccountId     string `json:"accountId"`
		ExecQty       string `json:"execQty"`
		OrderCategory int    `json:"orderCategory"`
	} `json:"result"`
	RetExtMap struct {
	} `json:"retExtMap"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
}
