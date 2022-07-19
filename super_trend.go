package tradeware

type SuperTrend struct {
	Indicator
}

func NewSuperTrend() IIndicator {
	return &SuperTrend{
		Indicator{
			Name: SUPER_TREND,
			Settings: &SuperTrendSettings{
				AtrLength:  10,
				Multiplier: 4.4,
				Source:     Hlcc4,
			},
			TimeFrame: ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type SuperTrendSettings struct {
	AtrLength  int
	Multiplier float64
	Source     Source
}
