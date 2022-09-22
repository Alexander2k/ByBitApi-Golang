package margin

import "ByBitApiV2/pkg"

type Margin struct {
	Signer *pkg.Signer
}

func NewMargin() *Margin {
	return &Margin{Signer: pkg.NewSigner()}
}
