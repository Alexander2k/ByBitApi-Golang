package margin

import (
	"github.com/Alexander2k/ByBitApi-Golang/config"
	"github.com/Alexander2k/ByBitApi-Golang/pkg"
)

type Margin struct {
	Signer *pkg.Signer
}

func NewMargin(c *config.Config) *Margin {
	return &Margin{Signer: pkg.NewSigner(c)}
}
