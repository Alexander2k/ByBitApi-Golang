package spot

import (
	"encoding/json"
	"fmt"
)

type Account interface {
	NewOrder(symbol string, orderQty string, side string, orderType string, timeInForce string, orderPrice string, orderLinkId string, orderCategory int, triggerPrice string) *Order
	BuyLimitOrder(symbol string, orderQty string, orderPrice string) *Order
	SellLimitOrder(symbol string, orderQty string, orderPrice string) *Order
	BuyMarketOrder(symbol string, orderQty string) *Order
	SellMarketOrder(symbol string, orderQty string) *Order
	PlaceOrder(order *Order) (OrderResponse, error)
}
type Order struct {
	Symbol        string `json:"symbol,omitempty"`
	OrderQty      string `json:"orderQty,omitempty"`
	Side          string `json:"side,omitempty"`
	OrderType     string `json:"orderType,omitempty"`
	TimeInForce   string `json:"timeInForce,omitempty"`
	OrderPrice    string `json:"orderPrice,omitempty"`
	OrderLinkId   string `json:"orderLinkId,omitempty"`
	OrderCategory int    `json:"orderCategory,omitempty"`
	TriggerPrice  string `json:"triggerPrice,omitempty"`
}

func (s *Spot) NewOrder(symbol string, orderQty string, side string, orderType string, timeInForce string, orderPrice string, orderLinkId string, orderCategory int, triggerPrice string) *Order {
	return &Order{Symbol: symbol, OrderQty: orderQty, Side: side, OrderType: orderType, TimeInForce: timeInForce, OrderPrice: orderPrice, OrderLinkId: orderLinkId, OrderCategory: orderCategory, TriggerPrice: triggerPrice}
}

func (s *Spot) BuyLimitOrder(symbol string, orderQty string, orderPrice string) *Order {
	return &Order{Symbol: symbol, OrderQty: orderQty, Side: BUY, OrderType: LIMITORDER, OrderPrice: orderPrice}
}

func (s *Spot) SellLimitOrder(symbol string, orderQty string, orderPrice string) *Order {
	return &Order{Symbol: symbol, OrderQty: orderQty, Side: SELL, OrderType: LIMITORDER, OrderPrice: orderPrice}
}

func (s *Spot) BuyMarketOrder(symbol string, orderQty string) *Order {
	return &Order{Symbol: symbol, OrderQty: orderQty, Side: BUY, OrderType: MARKETORDER}
}

func (s *Spot) SellMarketOrder(symbol string, orderQty string) *Order {
	return &Order{Symbol: symbol, OrderQty: orderQty, Side: SELL, OrderType: MARKETORDER}
}

func (s *Spot) PlaceOrder(order *Order) (OrderResponse, error) {
	var orderResponse OrderResponse
	var ord map[string]string

	orderM, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(orderM, &ord)

	point := "/spot/v3/private/order"

	response, err := s.Signer.Post(GET, point, ord)
	if err != nil {
		return OrderResponse{}, err
	}

	_ = json.Unmarshal(response, &orderResponse)

	return orderResponse, err

}
