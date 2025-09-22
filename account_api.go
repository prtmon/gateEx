package gateEx

import (
	"github.com/gateio/gateapi-go/v6"
)

func (e *Exchange) GetAccountDetail() (out gateapi.AccountDetail, err error) {
	out, _, err = e.ApiClient().AccountApi.GetAccountDetail(e.ApiCtx())
	return out, err
}
