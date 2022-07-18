package indicators

import (
	"tradechef_backtest/constants"
	"tradechef_backtest/internal/indicator"
)

type Ema struct {
	indicator.Indicator
}

func NewEma() indicator.IIndicator {
	return &Ema{
		Indicator: indicator.Indicator{
			Name: constants.Ema,
			Settings: &EmaSettings{
				EmaLength: 14,
			},
			TimeFrame: constants.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type EmaSettings struct {
	EmaLength int
}
