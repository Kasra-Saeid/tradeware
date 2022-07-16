package entities

import "tradechef_backtest/types"

type IIndicator interface {
	GetName() types.IndicatorName
	GetSettings() types.IndicatorSettings
	GetValue() []float64
	SetName(name types.IndicatorName)
	SetValue(value []float64)
	SetSettings(values ...types.IndicatorSettingsAttr)
}
