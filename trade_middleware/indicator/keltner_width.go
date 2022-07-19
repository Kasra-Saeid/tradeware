package indicator

import (
	"github.com/kasrasaeed/trade_vessel/trade_middleware/constants"
	"github.com/kasrasaeed/trade_vessel/trade_middleware/types"
)

type KeltnerWidth struct {
	Indicator
}

func NewKeltnerWidth() IIndicator {
	return &KeltnerWidth{
		Indicator{
			Name: constants.KeltnerWidth,
			Settings: &KeltnerWidthSettings{
				EmaLength:  14,
				Multiplier: 2,
				AtrLength:  10,
				Source:     constants.Close,
			},
			TimeFrame: constants.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type KeltnerWidthSettings struct {
	EmaLength  int
	Multiplier int
	AtrLength  int
	Source     types.Source
}
