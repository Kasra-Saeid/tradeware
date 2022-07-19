package indicator

import (
	"github.com/kasrasaeed/trade_vessel/trade_middleware/constants"
	"github.com/kasrasaeed/trade_vessel/trade_middleware/types"
)

type SuperTrend struct {
	Indicator
}

func NewSuperTrend() IIndicator {
	return &SuperTrend{
		Indicator{
			Name: constants.SuperTrend,
			Settings: &SuperTrendSettings{
				AtrLength:  10,
				Multiplier: 4.4,
				Source:     constants.Hlcc4,
			},
			TimeFrame: constants.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type SuperTrendSettings struct {
	AtrLength  int
	Multiplier float64
	Source     types.Source
}
