package gateEx

import (
	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
)

func (e *Exchange) GetTotalBalance() (out gateapi.TotalBalance, err error) {
	localVarOptionals := &gateapi.GetTotalBalanceOpts{}
	localVarOptionals.Currency = optional.NewString("USDT")
	out, _, err = e.ApiClient().WalletApi.GetTotalBalance(e.ApiCtx(), localVarOptionals)
	return out, err
}

func (e *Exchange) Transfer(settle, currency, from, to, amount string) (out gateapi.TransactionId, err error) {
	localVarOptionals := gateapi.Transfer{}
	localVarOptionals.Currency = currency
	localVarOptionals.From = from
	localVarOptionals.To = to
	localVarOptionals.Amount = amount
	localVarOptionals.Settle = settle
	out, _, err = e.ApiClient().WalletApi.Transfer(e.ApiCtx(), localVarOptionals)
	return out, err
}
