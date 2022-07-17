package indicator

import (
	"tradechef_backtest/constants"
)

type Adx struct {
	Indicator
}

func NewAdx() IIndicator {
	return &Adx{
		Indicator: Indicator{
			Name:  constants.Adx,
			Value: make([]float64, 0, 0),
			Settings: &AdxSettings{
				AdxLength: 14,
			},
		},
	}
}

type AdxSettings struct {
	AdxLength int
}
