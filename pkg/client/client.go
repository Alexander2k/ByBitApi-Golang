package client

import (
	"github.com/Alexander2k/ByBitApi-Golang/pkg/derivatives"
	"github.com/Alexander2k/ByBitApi-Golang/pkg/margin"
	"github.com/Alexander2k/ByBitApi-Golang/pkg/spot"
)

type API struct {
	Spot        *spot.Spot
	Margin      *margin.Margin
	Derivatives *derivatives.Derivatives
}

func NewByBitAPI() *API {
	return &API{
		Spot:        spot.NewSpot(),
		Margin:      margin.NewMargin(),
		Derivatives: derivatives.NewDerivatives(),
	}
}
