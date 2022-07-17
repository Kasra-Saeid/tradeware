package strategy

import "tradechef_backtest/internal/indicator"

type TwoEmaCrossUp struct {
	FirstEma  indicator.Ema
	SecondEma indicator.Ema
}

func (t *TwoEmaCrossUp) CanEnter() bool {
	firstIndicatorValue := t.FirstEma.GetValue()
	secondIndicatorValue := t.SecondEma.GetValue()
	if firstIndicatorValue[len(firstIndicatorValue)-1] > secondIndicatorValue[len(secondIndicatorValue)-1] &&
		firstIndicatorValue[len(firstIndicatorValue)-2] < secondIndicatorValue[len(secondIndicatorValue)-2] {
		return true
	}
	return false
}
