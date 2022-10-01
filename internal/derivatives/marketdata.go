package derivatives

import (
	"encoding/json"
	"fmt"
)

type MarketData interface {
	GetOrderBook(params OrderBookParams) (OrderBookResponse, error)
	GetKline(params KlineParams) (KlineResponse, error)
	GetTicker(params GetTickerParams) (TickerResponse, error)
	GetTickerOption(params GetTickerParams) (OptionResponse, error)
	GetInstrumentInfo(params GetInstrumentInfoParams) InstrumentInfoResponse
	GetOptionInfo(params GetInstrumentInfoParams) (OptionInfoResponse, error)
	GetMarkPriceKline(params KlineParams) (PriceKlineResponse, error)
	GetIndexPriceKline(params KlineParams) (PriceKlineResponse, error)
}

func (d *Derivatives) GetOrderBook(params OrderBookParams) (OrderBookResponse, error) {
	var orb OrderBookResponse
	query := d.QueryBuild(params)
	bytes, err := d.Signer.Get(GET, GetOrderBook+query)
	if err != nil {
		return OrderBookResponse{}, nil
	}
	err = json.Unmarshal(bytes, &orb)
	if err != nil {
		return OrderBookResponse{}, err
	}

	return orb, err
}

func (d *Derivatives) GetKline(params KlineParams) (KlineResponse, error) {
	var kl KlineResponse
	query := d.QueryBuild(params)

	bytes, err := d.Signer.Get(GET, GetKline+query)
	if err != nil {
		return KlineResponse{}, nil
	}

	err = json.Unmarshal(bytes, &kl)
	if err != nil {
		return KlineResponse{}, err
	}
	return kl, err
}

func (d *Derivatives) GetTicker(params GetTickerParams) (TickerResponse, error) {
	var t TickerResponse
	query := d.QueryBuild(params)

	switch {
	case params.Category == INVERSE:
		bytes, err := d.Signer.Get(GET, GetLatestInformationForSymbol+query)
		if err != nil {
			return TickerResponse{}, err
		}
		_ = json.Unmarshal(bytes, &t)
		return t, err
	case params.Category == LINEAR:
		bytes, err := d.Signer.Get(GET, GetLatestInformationForSymbol+query)
		if err != nil {
			return TickerResponse{}, err
		}
		_ = json.Unmarshal(bytes, &t)
		return t, err
	default:
		fmt.Println("ERROR: check parameters")
		return TickerResponse{}, nil
	}

}

func (d *Derivatives) GetTickerOption(params GetTickerParams) (OptionResponse, error) {
	var o OptionResponse
	query := d.QueryBuild(params)

	bytes, err := d.Signer.Get(GET, GetLatestInformationForSymbol+query)
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(bytes, &o)
	return o, err
}

func (d *Derivatives) GetInstrumentInfo(params GetInstrumentInfoParams) InstrumentInfoResponse {
	var inst InstrumentInfoResponse

	query := d.QueryBuild(params)

	switch {
	case params.Category == LINEAR:
		bytes, err := d.Signer.Get("GET", GetInstrumentInfo+query)
		if err != nil {
			return InstrumentInfoResponse{}
		}

		_ = json.Unmarshal(bytes, &inst)
		return inst
	case params.Category == INVERSE:

		bytes, err := d.Signer.Get("GET", GetInstrumentInfo+query)
		if err != nil {
			return InstrumentInfoResponse{}
		}

		_ = json.Unmarshal(bytes, &inst)
		return inst
	}
	return InstrumentInfoResponse{}
}

func (d *Derivatives) GetOptionInfo(params GetInstrumentInfoParams) (OptionInfoResponse, error) {
	var opt OptionInfoResponse
	q := d.QueryBuild(params)
	bytes, err := d.Signer.Get(GET, GetInstrumentInfo+q)
	if err != nil {
		return OptionInfoResponse{}, err
	}

	err = json.Unmarshal(bytes, &opt)
	if err != nil {
		return OptionInfoResponse{}, err
	}
	return opt, err
}

func (d *Derivatives) GetMarkPriceKline(params KlineParams) (PriceKlineResponse, error) {
	var mpk PriceKlineResponse
	query := d.QueryBuild(params)
	bytes, err := d.Signer.Get(GET, GetMarkPriceKline+query)
	if err != nil {

		return PriceKlineResponse{}, err
	}
	err = json.Unmarshal(bytes, &mpk)
	if err != nil {
		return PriceKlineResponse{}, err
	}

	return mpk, err
}

func (d *Derivatives) GetIndexPriceKline(params KlineParams) (PriceKlineResponse, error) {
	var ind PriceKlineResponse
	q := d.QueryBuild(params)
	bytes, err := d.Signer.Get(GET, GetIndexPriceKline+q)
	if err != nil {
		return PriceKlineResponse{}, err
	}
	err = json.Unmarshal(bytes, &ind)
	if err != nil {
		return PriceKlineResponse{}, err
	}
	return ind, err
}

func (d *Derivatives) GetFoundingRateHistory(params FoundingRateHistoryParams) (FoundingRateHistoryResponse, error) {
	var fr FoundingRateHistoryResponse
	q := d.QueryBuild(params)
	bytes, err := d.Signer.Get(GET, GetFundingRateHistory+q)
	if err != nil {
		return FoundingRateHistoryResponse{}, nil
	}

	err = json.Unmarshal(bytes, &fr)
	if err != nil {
		return FoundingRateHistoryResponse{}, err
	}

	return fr, err

}
