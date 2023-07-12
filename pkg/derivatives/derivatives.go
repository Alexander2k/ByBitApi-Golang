package derivatives

import (
	"encoding/json"
	"fmt"
	"github.com/Alexander2k/ByBitApi-Golang/config"
	"github.com/Alexander2k/ByBitApi-Golang/pkg"
	"net/http"
)

type Derivatives struct {
	Signer *pkg.Signer
}

func NewDerivatives(c *config.Config) *Derivatives {
	return &Derivatives{Signer: pkg.NewSigner(c)}
}

func (d *Derivatives) QueryBuild(params interface{}) string {
	var p map[string]string
	req, _ := http.NewRequest("GET", "", nil)
	q := req.URL.Query()

	m, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
	}
	_ = json.Unmarshal(m, &p)

	for k, v := range p {
		q.Add(k, v)
	}
	return q.Encode()
}
