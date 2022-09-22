package client

import (
	"ByBitApiV2/internal/margin"
	"ByBitApiV2/internal/spot"
)

type API struct {
	Spot   *spot.Spot
	Margin *margin.Margin
}

func NewAPI() *API {
	return &API{Spot: spot.NewSpot(), Margin: margin.NewMargin()}
}
