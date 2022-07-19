package trade_vessel

type Ema struct {
	Indicator
}

func NewEma() IIndicator {
	return &Ema{
		Indicator: Indicator{
			Name: EMA,
			Settings: &EmaSettings{
				EmaLength: 14,
				Source:    Close,
			},
			TimeFrame: ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type EmaSettings struct {
	EmaLength int
	Source    Source
}
