package indicator

import (
	"github.com/kasrasaeed/trade_vessel/constants"
)

type Adx struct {
	Indicator
}

func NewAdx() IIndicator {
	return &Adx{
		Indicator: Indicator{
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
