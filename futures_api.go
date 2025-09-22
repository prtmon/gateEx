package gateEx

import (
	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"strings"
)

/*
ListFuturesAccounts Query futures account
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param settle Settle currency

@return FuturesAccount
*/
func (e *Exchange) ListFuturesAccounts(settle string) (out gateapi.FuturesAccount, err error) {
	out, _, err = e.ApiClient().FuturesApi.ListFuturesAccounts(e.ApiCtx(), strings.ToLower(settle))
	return out, err
}

func (e *Exchange) ListFuturesContract(settle string) (out []gateapi.Contract, err error) {
	out, _, err = e.ApiClient().FuturesApi.ListFuturesContracts(e.ApiCtx(), strings.ToLower(settle), nil)
	return out, err
}

/*
GetFuturesContract Get a single contract
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param settle Settle currency
  - @param contract Futures contract

@return Contract
*/
func (e *Exchange) GetFuturesContract(settle string, contract string) (out gateapi.Contract, err error) {
	out, _, err = e.ApiClient().FuturesApi.GetFuturesContract(e.ApiCtx(), strings.ToLower(settle), contract)
	return out, err
}

func (e *Exchange) ListFuturesTickers(settle string, contract string) (out []gateapi.FuturesTicker, err error) {
	localVarOptionals := &gateapi.ListFuturesTickersOpts{}
	if contract != "" {
		localVarOptionals.Contract = optional.NewString(contract)
	}
	out, _, err = e.ApiClient().FuturesApi.ListFuturesTickers(e.ApiCtx(), strings.ToLower(settle), localVarOptionals)
	return out, err
}

/*
ListFuturesCandlesticks Get futures candlesticks
Return specified contract candlesticks. If prefix &#x60;contract&#x60; with &#x60;mark_&#x60;, the contract&#39;s mark price candlesticks are returned; if prefix with &#x60;index_&#x60;, index price candlesticks will be returned.  Maximum of 2000 points are returned in one query. Be sure not to exceed the limit when specifying &#x60;from&#x60;, &#x60;to&#x60; and &#x60;interval&#x60;
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param settle Settle currency
  - @param contract Futures contract
  - @param optional nil or *ListFuturesCandlesticksOpts - Optional Parameters:
  - @param "From" (optional.Int64) -  Start time of candlesticks, formatted in Unix timestamp in seconds. Default to`to - 100 * interval` if not specified
  - @param "To" (optional.Int64) -  End time of candlesticks, formatted in Unix timestamp in seconds. Default to current time
  - @param "Limit" (optional.Int32) -  Maximum recent data points returned. `limit` is conflicted with `from` and `to`. If either `from` or `to` is specified, request will be rejected.
  - @param "Interval" (optional.String) -  Interval time between data points

@return []FuturesCandlestick
*/
func (e *Exchange) ListFuturesCandlesticks(settle string, contract string, from int64, to int64, limit int32, interval string) (out FuturesCandlesticks, err error) {
	localVarOptionals := &gateapi.ListFuturesCandlesticksOpts{}

	if from > 0 {
		localVarOptionals.From = optional.NewInt64(from)
	}
	if to > 0 {
		localVarOptionals.To = optional.NewInt64(to)
	}
	if limit > 0 {
		localVarOptionals.Limit = optional.NewInt32(limit)
	}
	if interval != "" {
		localVarOptionals.Interval = optional.NewString(interval)
	}

	out, _, err = e.ApiClient().FuturesApi.ListFuturesCandlesticks(e.ApiCtx(), strings.ToLower(settle), contract, localVarOptionals)
	return out, err
}

func (e *Exchange) ListFuturesOrderBook(settle string, contract string, limit int32, interval string) (out gateapi.FuturesOrderBook, err error) {
	localVarOptionals := &gateapi.ListFuturesOrderBookOpts{}

	if limit > 0 {
		localVarOptionals.Limit = optional.NewInt32(limit)
	}
	if interval != "" {
		localVarOptionals.Interval = optional.NewString(interval)
	}
	out, _, err = e.ApiClient().FuturesApi.ListFuturesOrderBook(e.ApiCtx(), strings.ToLower(settle), contract, localVarOptionals)
	return out, err
}

func (e *Exchange) GetDualModePosition(settle string, contract string) (out []gateapi.Position, err error) {
	out, _, err = e.ApiClient().FuturesApi.GetDualModePosition(e.ApiCtx(), strings.ToLower(settle), contract)
	return out, err
}

func (e *Exchange) GetPosition(settle string, contract string) (out gateapi.Position, err error) {
	out, _, err = e.ApiClient().FuturesApi.GetPosition(e.ApiCtx(), strings.ToLower(settle), contract)
	return out, err
}

func (e *Exchange) ListPositions(settle string, holding bool) (out []gateapi.Position, err error) {
	localVarOptionals := &gateapi.ListPositionsOpts{
		Holding: optional.NewBool(holding),
	}
	out, _, err = e.ApiClient().FuturesApi.ListPositions(e.ApiCtx(), strings.ToLower(settle), localVarOptionals)
	return out, err
}

func (e *Exchange) CancelFuturesOrders(settle string, contract string, side string) (out []gateapi.FuturesOrder, err error) {
	localVarOptionals := &gateapi.CancelFuturesOrdersOpts{}
	if side != "" {
		localVarOptionals.Side = optional.NewString(side)
	}
	out, _, err = e.ApiClient().FuturesApi.CancelFuturesOrders(e.ApiCtx(), strings.ToLower(settle), contract, localVarOptionals)
	return out, err
}

func (e *Exchange) CreateFuturesOrder(settle string, futuresOrder gateapi.FuturesOrder) (out gateapi.FuturesOrder, err error) {
	out, _, err = e.ApiClient().FuturesApi.CreateFuturesOrder(e.ApiCtx(), strings.ToLower(settle), futuresOrder, nil)
	return out, err
}

func (e *Exchange) ListFuturesOrders(settle string, contract string, status string, limit int32, offset int32, lastId string) (out []gateapi.FuturesOrder, err error) {
	localVarOptionals := &gateapi.ListFuturesOrdersOpts{}
	if limit > 0 {
		localVarOptionals.Limit = optional.NewInt32(limit)
	}
	if offset > 0 {
		localVarOptionals.Offset = optional.NewInt32(offset)
	}
	if lastId != "" {
		localVarOptionals.LastId = optional.NewString(lastId)
	}
	if contract != "" {
		localVarOptionals.LastId = optional.NewString(contract)
	}
	out, _, err = e.ApiClient().FuturesApi.ListFuturesOrders(e.ApiCtx(), strings.ToLower(settle), status, localVarOptionals)
	return out, err
}

func (e *Exchange) CancelFuturesOrder(settle string, orderId string) (out gateapi.FuturesOrder, err error) {
	out, _, err = e.ApiClient().FuturesApi.CancelFuturesOrder(e.ApiCtx(), strings.ToLower(settle), orderId, nil)
	return out, err
}

func (e *Exchange) UpdatePositionLeverage(settle, contract, leverage string, isCross bool) (out gateapi.Position, err error) {
	localVarOptionals := &gateapi.UpdatePositionLeverageOpts{}
	if isCross {
		localVarOptionals.CrossLeverageLimit = optional.NewString(leverage)
		leverage = "0"
	}
	out, _, err = e.ApiClient().FuturesApi.UpdatePositionLeverage(e.ApiCtx(), strings.ToLower(settle), contract, leverage, localVarOptionals)
	return out, err
}

func (e *Exchange) SetDualMode(settle string, dualMode bool) (out gateapi.FuturesAccount, err error) {
	out, _, err = e.ApiClient().FuturesApi.SetDualMode(e.ApiCtx(), strings.ToLower(settle), dualMode)
	return out, err
}

// ListFuturesAccountBook 查询合约账户变更历史
func (e *Exchange) ListFuturesAccountBook(settle, contract, type_ string, limit, offset int32, from, to int64) (out []gateapi.FuturesAccountBook, err error) {
	localVarOptionals := &gateapi.ListFuturesAccountBookOpts{}
	if limit > 0 {
		localVarOptionals.Limit = optional.NewInt32(limit)
	}
	if offset > 0 {
		localVarOptionals.Offset = optional.NewInt32(offset)
	}
	if from > 0 {
		localVarOptionals.From = optional.NewInt64(from)
	}
	if to > 0 {
		localVarOptionals.To = optional.NewInt64(to)
	}
	if contract != "" {
		localVarOptionals.Contract = optional.NewString(contract)
	}
	if type_ != "" {
		localVarOptionals.Type_ = optional.NewString(type_)
	}
	out, _, err = e.ApiClient().FuturesApi.ListFuturesAccountBook(e.ApiCtx(), strings.ToLower(settle), localVarOptionals)
	return out, err
}

/*
ListPositionClose 查询平仓历史
- @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
- @param settle Settle currency
- @param optional nil or *ListPositionCloseOpts - Optional Parameters:
- @param "Contract" (optional.String) -  Futures contract, return related data only if specified
- @param "Limit" (optional.Int32) -  Maximum number of records to be returned in a single list
- @param "Offset" (optional.Int32) -  List offset, starting from 0
- @param "From" (optional.Int64) -  Start timestamp
- @param "To" (optional.Int64) -  End timestamp
- @param "Side" (optional.String) -  Query side.  long or shot
- @param "Pnl" (optional.String) -  Query profit or loss
*/
func (e *Exchange) ListPositionClose(settle, contract string, limit, offset int32, from, to int64, side, pnl string) (out []gateapi.PositionClose, err error) {
	localVarOptionals := &gateapi.ListPositionCloseOpts{}
	if limit > 0 {
		localVarOptionals.Limit = optional.NewInt32(limit)
	}
	if offset > 0 {
		localVarOptionals.Offset = optional.NewInt32(offset)
	}
	if from > 0 {
		localVarOptionals.From = optional.NewInt64(from)
	}
	if to > 0 {
		localVarOptionals.To = optional.NewInt64(to)
	}
	if contract != "" {
		localVarOptionals.Contract = optional.NewString(contract)
	}
	if side != "" {
		localVarOptionals.Side = optional.NewString(side)
	}
	if pnl != "" {
		localVarOptionals.Pnl = optional.NewString(pnl)
	}
	out, _, err = e.ApiClient().FuturesApi.ListPositionClose(e.ApiCtx(), strings.ToLower(settle), localVarOptionals)
	return out, err
}

//一键平仓功能
//获取持仓列表,for列表仓位平仓,time.sleep以免限速
