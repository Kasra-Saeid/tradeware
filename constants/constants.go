package constants

import "tradechef_backtest/types"

const (
	Adx = types.IndicatorName(iota)
	Ema
	SuperTrend
	KeltnerWidth
	Atr
)

const (
	AdxLength = types.IndicatorSettingsAttrName(iota)
	EmaLength
	AtrLength
	Multiplier
)

const (
	OneMin       = types.TimeFrame("1")
	ThreeMin     = types.TimeFrame("3")
	FiveMin      = types.TimeFrame("5")
	FifteenMin   = types.TimeFrame("15")
	ThirtyMin    = types.TimeFrame("30")
	FortyFiveMin = types.TimeFrame("45")
	OneHour      = types.TimeFrame("60")
	TwoHour      = types.TimeFrame("120")
	ThreeHour    = types.TimeFrame("180")
	FourHour     = types.TimeFrame("240")
	OneDay       = types.TimeFrame("1d")
	ThreeDay     = types.TimeFrame("3d")
	OneWeek      = types.TimeFrame("1w")
	OneMonth     = types.TimeFrame("1m")
)

const (
	Buy   = types.Side("Buy")
	Sell  = types.Side("Sell")
	Long  = types.Side("Long")
	Short = types.Side("Short")
)

const (
	Open = types.Source(iota)
	High
	Low
	Close
	Hl2
	Hlc3
	Ohlc4
	Hlcc4
)

const (
	InProgress = types.Status("InProgress")
	Available  = types.Status("Available")
)
