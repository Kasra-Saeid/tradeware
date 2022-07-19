package indicator

import (
	"github.com/kasrasaeed/trade_vessel/constants"
	"github.com/kasrasaeed/trade_vessel/types"
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
				Source:    constants.Close,
			},
			TimeFrame: constants.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type EmaSettings struct {
	EmaLength int
	Source    types.Source
}
