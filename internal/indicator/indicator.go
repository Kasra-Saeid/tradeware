package indicator

import (
	"tradechef_backtest/constants"
	"tradechef_backtest/internal/indicator/indicators"
	"tradechef_backtest/types"
)

type Indicator struct {
	Name      types.IndicatorName
	Settings  types.IndicatorSettings
	TimeFrame types.TimeFrame
	Value     []float64
}

func (i *Indicator) GetValue() []float64 {
	return i.Value
}

func (i *Indicator) GetName() types.IndicatorName {
	return i.Name
}

func (i *Indicator) GetSettings() types.IndicatorSettings {
	switch i.Settings.(type) {
	case *indicators.AdxSettings:
		return i.Settings.(*indicators.AdxSettings)
	case *indicators.EmaSettings:
		return i.Settings.(*indicators.EmaSettings)
	default:
		return i.Settings
	}
}

func (i *Indicator) GetTimeFrame() types.TimeFrame {
	return i.TimeFrame
}

func (i *Indicator) SetValue(value []float64) {
	i.Value = value
}

func (i *Indicator) SetName(name types.IndicatorName) {
	i.Name = name
}

func (i *Indicator) SetSettings(attrs ...IndicatorSettingsAttr) {
	switch i.Settings.(type) {
	case *indicators.AdxSettings:
		for _, v := range attrs {
			if v.Attr == constants.AdxLength {
				i.Settings.(*indicators.AdxSettings).AdxLength = (int)(v.Value)
			}
		}
	case *indicators.EmaSettings:
		for _, v := range attrs {
			if v.Attr == constants.EmaLength {
				i.Settings.(*indicators.EmaSettings).EmaLength = (int)(v.Value)
			}
		}
	}
}

func (i *Indicator) SetTimeFrame(timeFrame types.TimeFrame) {
	i.TimeFrame = timeFrame
}
