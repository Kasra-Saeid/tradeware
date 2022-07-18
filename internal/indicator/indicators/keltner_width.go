package indicators

import (
	"tradechef_backtest/constants"
	"tradechef_backtest/internal/indicator"
	"tradechef_backtest/types"
)

type KeltnerWidth struct {
	indicator.Indicator
}

func NewKeltnerWidth() indicator.IIndicator {
	return &KeltnerWidth{
		indicator.Indicator{
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
