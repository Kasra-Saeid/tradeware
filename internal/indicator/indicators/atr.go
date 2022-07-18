package indicators

import (
	"tradechef_backtest/constants"
	"tradechef_backtest/internal/indicator"
)

type Atr struct {
	indicator.Indicator
}

func NewAtr() indicator.IIndicator {
	return &Atr{
		indicator.Indicator{
			Name: constants.Atr,
			Settings: &AtrSettings{
				AtrLength: 14,
			},
			TimeFrame: constants.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type AtrSettings struct {
	AtrLength int
}
