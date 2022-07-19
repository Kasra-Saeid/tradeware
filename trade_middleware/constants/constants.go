package constants

import (
	types2 "github.com/kasrasaeed/trade_vessel/trade_middleware/types"
)

const (
	Adx = types2.IndicatorName(iota)
	Ema
	SuperTrend
	KeltnerWidth
	Atr
)

const (
	AdxLength = types2.IndicatorSettingsAttrName(iota)
	EmaLength
	AtrLength
	Multiplier
)

const (
	OneMin       = types2.TimeFrame("1")
	ThreeMin     = types2.TimeFrame("3")
	FiveMin      = types2.TimeFrame("5")
	FifteenMin   = types2.TimeFrame("15")
	ThirtyMin    = types2.TimeFrame("30")
	FortyFiveMin = types2.TimeFrame("45")
	OneHour      = types2.TimeFrame("60")
	TwoHour      = types2.TimeFrame("120")
	ThreeHour    = types2.TimeFrame("180")
	FourHour     = types2.TimeFrame("240")
	OneDay       = types2.TimeFrame("1d")
	ThreeDay     = types2.TimeFrame("3d")
	OneWeek      = types2.TimeFrame("1w")
	OneMonth     = types2.TimeFrame("1m")
)

const (
	Buy   = types2.Side("Buy")
	Sell  = types2.Side("Sell")
	Long  = types2.Side("Long")
	Short = types2.Side("Short")
)

const (
	Open = types2.Source(iota)
	High
	Low
	Close
	Hl2
	Hlc3
	Ohlc4
	Hlcc4
)

const (
	InProgress = types2.Status("InProgress")
	Available  = types2.Status("Available")
)
