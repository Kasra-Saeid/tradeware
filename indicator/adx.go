package indicators

import (
	"github.com/kasrasaeed/trade_vessel/constants"
	indicator2 "github.com/kasrasaeed/trade_vessel/indicator"
)

type Adx struct {
	indicator2.Indicator
}

func NewAdx() indicator2.IIndicator {
	return &Adx{
		Indicator: indicator2.Indicator{
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
