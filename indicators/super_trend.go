package indicators

import (
	"github.com/kasrasaeed/trade_vessel/constants"
	indicator2 "github.com/kasrasaeed/trade_vessel/indicator"
	"github.com/kasrasaeed/trade_vessel/types"
)

type SuperTrend struct {
	indicator2.Indicator
}

func NewSuperTrend() indicator2.IIndicator {
	return &SuperTrend{
		indicator2.Indicator{
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
