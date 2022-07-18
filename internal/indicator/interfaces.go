package indicator

import (
	"tradechef_backtest/types"
)

type IIndicator interface {
	GetName() types.IndicatorName
	GetSettings() types.IndicatorSettings
	GetTimeFrame() types.TimeFrame
	GetValue() []float64
	SetName(name types.IndicatorName)
	SetValue(value []float64)
	SetSettings(values ...IndicatorSettingsAttr)
	SetTimeFrame(timeFrame types.TimeFrame)
}
