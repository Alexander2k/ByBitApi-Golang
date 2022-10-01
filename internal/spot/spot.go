package spot

import "ByBitApi-Golang/pkg"

type Spot struct {
	Signer *pkg.Signer
}

func NewSpot() *Spot {
	return &Spot{Signer: pkg.NewSigner()}
}
