package indicators

import (
	"github.com/kasrasaeed/trade_vessel/constants"
	"github.com/kasrasaeed/trade_vessel/internal/indicator"
	"github.com/kasrasaeed/trade_vessel/types"
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
