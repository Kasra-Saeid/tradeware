package constants

import "tradechef_backtest/types"

const (
	Adx = types.IndicatorName(iota)
	Ema
	SuperTrend
	KeltnerWidth
)

const (
	AdxLength = types.IndicatorSettingsAttrName(iota)
	EmaLength
)
