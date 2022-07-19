package indicator

import (
	types2 "github.com/kasrasaeed/trade_vessel/trade_middleware/types"
)

type IIndicator interface {
	GetName() types2.IndicatorName
	GetSettings() types2.IndicatorSettings
	GetTimeFrame() types2.TimeFrame
	GetValue() []float64
	SetName(name types2.IndicatorName)
	SetValue(value []float64)
	SetSettings(source *types2.Source, values ...IndicatorSettingsAttr)
	SetTimeFrame(timeFrame types2.TimeFrame)
}
