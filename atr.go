package trade_vessel

type Atr struct {
	Indicator
}

func NewAtr() IIndicator {
	return &Atr{
		Indicator{
			Name: ATR,
			Settings: &AtrSettings{
				AtrLength: 14,
			},
			TimeFrame: ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type AtrSettings struct {
	AtrLength int
}
