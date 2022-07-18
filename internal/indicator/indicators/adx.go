package indicators

import (
	"tradechef_backtest/constants"
	"tradechef_backtest/internal/indicator"
)

type Adx struct {
	indicator.Indicator
}

func NewAdx() indicator.IIndicator {
	return &Adx{
		Indicator: indicator.Indicator{
			Name: constants.Adx,
			Settings: &AdxSettings{
				AdxLength: 14,
			},
			TimeFrame: constants.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type AdxSettings struct {
	AdxLength int
}
