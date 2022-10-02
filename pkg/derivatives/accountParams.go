package derivatives

type Order struct {
	Category       string `json:"category"`
	Symbol         string `json:"symbol"`
	Side           string `json:"side"`
	PositionIdx    string `json:"positionIdx,omitempty"`
	OrderType      string `json:"orderType"`
	Qty            string `json:"qty"`
	Price          string `json:"price,omitempty"`
	BasePrice      string `json:"basePrice,omitempty"`
	TriggerPrice   string `json:"triggerPrice,omitempty"`
	TriggerBy      string `json:"triggerBy,omitempty"`
	Iv             string `json:"iv,omitempty"`
	TimeInForce    string `json:"timeInForce"`
	OrderLinkId    string `json:"orderLinkId,omitempty"`
	TakeProfit     string `json:"takeProfit,omitempty"`
	StopLoss       string `json:"stopLoss,omitempty"`
	TpTriggerBy    string `json:"tpTriggerBy,omitempty"`
	SlTriggerBy    string `json:"slTriggerBy,omitempty"`
	ReduceOnly     string `json:"reduceOnly,omitempty"`
	CloseOnTrigger string `json:"closeOnTrigger,omitempty"`
	Mmp            string `json:"mmp,omitempty"`
}
