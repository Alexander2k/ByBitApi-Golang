package client

import (
	"ByBitApi-Golang/internal/derivatives"
	"ByBitApi-Golang/internal/margin"
	"ByBitApi-Golang/internal/spot"
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
