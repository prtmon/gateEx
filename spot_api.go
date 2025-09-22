package gateEx

import (
	"github.com/gateio/gateapi-go/v6"
)

func (e *Exchange) ListCurrencyPairs() (out []gateapi.CurrencyPair, err error) {
	out, _, err = e.ApiClient().SpotApi.ListCurrencyPairs(e.ApiCtx())
	return out, err
}
