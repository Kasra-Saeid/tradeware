package main

import (
	"fmt"
	"tradechef_backtest/constants"
	"tradechef_backtest/internal/indicator"
	"tradechef_backtest/internal/indicator/indicators"
)

func main() {

	ema := indicators.NewEma()
	emaSettings := indicator.IndicatorSettingsAttr{
		Attr:  constants.EmaLength,
		Value: 20,
	}
	emaSettings2 := indicator.IndicatorSettingsAttr{
		Attr:  constants.EmaLength,
		Value: 40,
	}
	ema.SetSettings(emaSettings, emaSettings2)
	fmt.Println(ema.GetSettings().(*indicators.EmaSettings).EmaLength)

}
