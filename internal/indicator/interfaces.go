package indicator

import (
	"github.com/kasrasaeed/trade_vessel/types"
)

type IIndicator interface {
	GetName() types.IndicatorName
	GetSettings() types.IndicatorSettings
	GetTimeFrame() types.TimeFrame
	GetValue() []float64
	SetName(name types.IndicatorName)
	SetValue(value []float64)
	SetSettings(source *types.Source, values ...IndicatorSettingsAttr)
	SetTimeFrame(timeFrame types.TimeFrame)
}
