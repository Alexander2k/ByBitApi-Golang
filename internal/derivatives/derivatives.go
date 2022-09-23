package derivatives

import "ByBitApiV2/pkg"

type Derivatives struct {
	Signer *pkg.Signer
}

func NewDerivatives() *Derivatives {
	return &Derivatives{Signer: pkg.NewSigner()}
}
