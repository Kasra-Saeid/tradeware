package trade_vessel

type Adx struct {
	Indicator
}

func NewAdx() IIndicator {
	return &Adx{
		Indicator: Indicator{
			Name: ADX,
			Settings: &AdxSettings{
				AdxLength: 14,
			},
			TimeFrame: ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type AdxSettings struct {
	AdxLength int
}
