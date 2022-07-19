package indicators

import (
	"github.com/kasrasaeed/trade_vessel/constants"
	indicator2 "github.com/kasrasaeed/trade_vessel/indicator"
	"github.com/kasrasaeed/trade_vessel/types"
)

type Ema struct {
	indicator2.Indicator
}

func NewEma() indicator2.IIndicator {
	return &Ema{
		Indicator: indicator2.Indicator{
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
