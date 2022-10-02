package derivatives

import (
	"encoding/json"
	"fmt"
)

type Account interface {
	PlaceOrder(params Order) (OrderResponse, error)
}

func (d *Derivatives) PlaceOrder(params Order) (OrderResponse, error) {
	var orderResponse OrderResponse
	var ord map[string]string

	orderM, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(orderM, &ord)

	q := d.QueryBuild(params)
	bytes, err := d.Signer.Post(POST, q+PlaceOrder, ord)
	if err != nil {
		return OrderResponse{}, err
	}

	err = json.Unmarshal(bytes, &orderResponse)
	if err != nil {
		return OrderResponse{}, err
	}

	return orderResponse, err
}

func (d *Derivatives) ReplaceOrder() {

}
