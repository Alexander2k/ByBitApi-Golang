package spot

import "ByBitApiV2/pkg"

type Spot struct {
	Signer *pkg.Signer
}

func NewSpot() *Spot {
	return &Spot{Signer: pkg.NewSigner()}
}
