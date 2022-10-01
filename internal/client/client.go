package client

import (
	"ByBitApiV2/internal/derivatives"
	"ByBitApiV2/internal/margin"
	"ByBitApiV2/internal/spot"
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
