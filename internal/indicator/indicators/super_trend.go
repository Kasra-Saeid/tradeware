package indicators

import (
	"github.com/kasrasaeed/trade_vessel/constants"
	"github.com/kasrasaeed/trade_vessel/internal/indicator"
	"github.com/kasrasaeed/trade_vessel/types"
)

type SuperTrend struct {
	indicator.Indicator
}

func NewSuperTrend() indicator.IIndicator {
	return &SuperTrend{
		indicator.Indicator{
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
