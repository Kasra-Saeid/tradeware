package indicator

import (
	"github.com/kasrasaeed/trade_vessel/trade_middleware/constants"
)

type Atr struct {
	Indicator
}

func NewAtr() IIndicator {
	return &Atr{
		Indicator{
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
