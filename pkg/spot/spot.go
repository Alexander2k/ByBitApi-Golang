package spot

import (
	"github.com/Alexander2k/ByBitApi-Golang/config"
	"github.com/Alexander2k/ByBitApi-Golang/pkg"
)

type Spot struct {
	Signer *pkg.Signer
}

func NewSpot(c *config.Config) *Spot {
	return &Spot{Signer: pkg.NewSigner(c)}
}
