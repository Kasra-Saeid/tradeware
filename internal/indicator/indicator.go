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
	case *indicators.AtrSettings:
		return i.Settings.(*indicators.AtrSettings)
	case *indicators.KeltnerWidthSettings:
		return i.Settings.(*indicators.KeltnerWidthSettings)
	case *indicators.SuperTrendSettings:
		return i.Settings.(*indicators.SuperTrendSettings)
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

func (i *Indicator) SetSettings(source *types.Source, attrs ...IndicatorSettingsAttr) {
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
	case *indicators.AtrSettings:
		for _, v := range attrs {
			if v.Attr == constants.AtrLength {
				i.Settings.(*indicators.AtrSettings).AtrLength = (int)(v.Value)
			}
		}
	case *indicators.KeltnerWidthSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == constants.EmaLength {
				i.Settings.(*indicators.KeltnerWidthSettings).EmaLength = (int)(v.Value)
			} else if v.Attr == constants.AtrLength {
				i.Settings.(*indicators.KeltnerWidthSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == constants.Multiplier {
				i.Settings.(*indicators.KeltnerWidthSettings).Multiplier = (int)(v.Value)
			}
		}
		i.Settings.(*indicators.KeltnerWidthSettings).Source = *source

	case *indicators.SuperTrendSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == constants.AtrLength {
				i.Settings.(*indicators.SuperTrendSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == constants.Multiplier {
				i.Settings.(*indicators.SuperTrendSettings).Multiplier = v.Value
			}
		}
		i.Settings.(*indicators.SuperTrendSettings).Source = *source

	}
}

func (i *Indicator) SetTimeFrame(timeFrame types.TimeFrame) {
	i.TimeFrame = timeFrame
}
