package tradeware

type KeltnerWidth struct {
	Indicator
}

func NewKeltnerWidth() IIndicator {
	return &KeltnerWidth{
		Indicator{
			Name: KELTNER_WIDTH,
			Settings: &KeltnerWidthSettings{
				EmaLength:  14,
				Multiplier: 2,
				AtrLength:  10,
				Source:     Close,
			},
			TimeFrame: ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type KeltnerWidthSettings struct {
	EmaLength  int
	Multiplier int
	AtrLength  int
	Source     SourceName
}
