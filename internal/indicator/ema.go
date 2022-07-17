package indicator

import (
	"tradechef_backtest/constants"
)

type Ema struct {
	Indicator
}

func NewEma() IIndicator {
	return &Ema{
		Indicator: Indicator{
			Name: constants.Ema,
			Settings: &EmaSettings{
				EmaLength: 14,
			},
			Value: make([]float64, 0, 0),
		},
	}
}

type EmaSettings struct {
	EmaLength int
}
