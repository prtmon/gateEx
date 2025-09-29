package gateEx

import (
	"github.com/gateio/gateapi-go/v6"
	"github.com/prtmon/finance/common"
	"github.com/prtmon/tools"
)

type FuturesCandlesticks []gateapi.FuturesCandlestick

func (f FuturesCandlesticks) ToCandlesticks() common.Candlesticks {
	candlesticks := make(common.Candlesticks, len(f))
	for i, k := range f {
		candlesticks[i].Time = k.T
		candlesticks[i].Open = tools.ToFloat64(k.O)
		candlesticks[i].High = tools.ToFloat64(k.H)
		candlesticks[i].Low = tools.ToFloat64(k.L)
		candlesticks[i].Close = tools.ToFloat64(k.C)
		candlesticks[i].Volume = tools.ToFloat64(k.V)
	}
	return candlesticks
}

func (f FuturesCandlesticks) ToOhlcv() common.OHLCV {
	ohlcv := common.OHLCV{
		Time:   make([]float64, len(f)),
		Open:   make([]float64, len(f)),
		High:   make([]float64, len(f)),
		Low:    make([]float64, len(f)),
		Close:  make([]float64, len(f)),
		Volume: make([]float64, len(f)),
	}

	for i, k := range f {
		ohlcv.Time[i] = k.T
		ohlcv.Open[i] = tools.ToFloat64(k.O)
		ohlcv.High[i] = tools.ToFloat64(k.H)
		ohlcv.Low[i] = tools.ToFloat64(k.L)
		ohlcv.Close[i] = tools.ToFloat64(k.C)
		ohlcv.Volume[i] = tools.ToFloat64(k.V)
	}

	return ohlcv
}
