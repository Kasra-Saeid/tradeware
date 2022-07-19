package indicators

import (
	"github.com/kasrasaeed/trade_vessel/constants"
	indicator2 "github.com/kasrasaeed/trade_vessel/indicator"
)

type Atr struct {
	indicator2.Indicator
}

func NewAtr() indicator2.IIndicator {
	return &Atr{
		indicator2.Indicator{
			Name: constants.Atr,
			Settings: &AtrSettings{
				AtrLength: 14,
			},
			TimeFrame: constants.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type AtrSettings struct {
	AtrLength int
}
