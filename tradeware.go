package tradeware

type TradeWare struct {
	timeFrames []TimeFrame
	indicators []IIndicator
	Sources    []SourceName
	Symbol     string
}

func NewTradeWare(timeFrames []TimeFrame, indicators []IIndicator, sources []SourceName, symbol string) *TradeWare {
	return &TradeWare{
		timeFrames: timeFrames,
		indicators: indicators,
		Sources:    sources,
		Symbol:     symbol,
	}
}
