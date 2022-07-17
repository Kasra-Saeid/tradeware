package main

import (
	"fmt"
	"tradechef_backtest/constants"
	"tradechef_backtest/internal/indicator"
)

func main() {

	ema := indicator.NewEma()
	emaSettings := indicator.IndicatorSettingsAttr{
		Attr:  constants.EmaLength,
		Value: 20,
	}
	emaSettings2 := indicator.IndicatorSettingsAttr{
		Attr:  constants.EmaLength,
		Value: 40,
	}
	ema.SetSettings(emaSettings, emaSettings2)
	fmt.Println(ema.GetSettings().(*indicator.EmaSettings).EmaLength)

}
