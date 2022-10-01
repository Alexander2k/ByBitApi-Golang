package margin

import "github.com/Alexander2k/ByBitApi-Golang/pkg"

type Margin struct {
	Signer *pkg.Signer
}

func NewMargin() *Margin {
	return &Margin{Signer: pkg.NewSigner()}
}
