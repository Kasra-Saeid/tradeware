package tradeware

const (
	ADX = IndicatorName(iota)
	EMA
	SUPER_TREND
	KELTNER_WIDTH
	ATR
)

const (
	AdxLength = IndicatorSettingsAttrName(iota)
	EmaLength
	AtrLength
	Multiplier
)

const (
	OneMin       = TimeFrame("1")
	ThreeMin     = TimeFrame("3")
	FiveMin      = TimeFrame("5")
	FifteenMin   = TimeFrame("15")
	ThirtyMin    = TimeFrame("30")
	FortyFiveMin = TimeFrame("45")
	OneHour      = TimeFrame("60")
	TwoHour      = TimeFrame("120")
	ThreeHour    = TimeFrame("180")
	FourHour     = TimeFrame("240")
	OneDay       = TimeFrame("1d")
	ThreeDay     = TimeFrame("3d")
	OneWeek      = TimeFrame("1w")
	OneMonth     = TimeFrame("1m")
)

const (
	Buy   = Side("Buy")
	Sell  = Side("Sell")
	Long  = Side("Long")
	Short = Side("Short")
)

const (
	Open = Source(iota)
	High
	Low
	Close
	Hl2
	Hlc3
	Ohlc4
	Hlcc4
)

const (
	InProgress = Status("InProgress")
	Available  = Status("Available")
)
