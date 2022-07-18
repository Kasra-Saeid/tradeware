package indicator

import (
	"tradechef_backtest/constants"
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
	case *AdxSettings:
		return i.Settings.(*AdxSettings)
	case *EmaSettings:
		return i.Settings.(*EmaSettings)
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
	case *AdxSettings:
		for _, v := range attrs {
			if v.Attr == constants.AdxLength {
				i.Settings.(*AdxSettings).AdxLength = (int)(v.Value)
			}
		}
	case *EmaSettings:
		for _, v := range attrs {
			if v.Attr == constants.EmaLength {
				i.Settings.(*EmaSettings).EmaLength = (int)(v.Value)
			}
		}
	}
}

func (i *Indicator) SetTimeFrame(timeFrame types.TimeFrame) {
	i.TimeFrame = timeFrame
}
