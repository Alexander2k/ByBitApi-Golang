package spot

import "github.com/Alexander2k/ByBitApi-Golang/pkg"

type Spot struct {
	Signer *pkg.Signer
}

func NewSpot() *Spot {
	return &Spot{Signer: pkg.NewSigner()}
}
