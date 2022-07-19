package indicator

import (
	"github.com/kasrasaeed/trade_vessel/constants"
	"github.com/kasrasaeed/trade_vessel/types"
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
	case *AtrSettings:
		return i.Settings.(*AtrSettings)
	case *KeltnerWidthSettings:
		return i.Settings.(*KeltnerWidthSettings)
	case *SuperTrendSettings:
		return i.Settings.(*SuperTrendSettings)
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
	case *AtrSettings:
		for _, v := range attrs {
			if v.Attr == constants.AtrLength {
				i.Settings.(*AtrSettings).AtrLength = (int)(v.Value)
			}
		}
	case *KeltnerWidthSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == constants.EmaLength {
				i.Settings.(*KeltnerWidthSettings).EmaLength = (int)(v.Value)
			} else if v.Attr == constants.AtrLength {
				i.Settings.(*KeltnerWidthSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == constants.Multiplier {
				i.Settings.(*KeltnerWidthSettings).Multiplier = (int)(v.Value)
			}
		}
		i.Settings.(*KeltnerWidthSettings).Source = *source

	case *SuperTrendSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == constants.AtrLength {
				i.Settings.(*SuperTrendSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == constants.Multiplier {
				i.Settings.(*SuperTrendSettings).Multiplier = v.Value
			}
		}
		i.Settings.(*SuperTrendSettings).Source = *source

	}
}

func (i *Indicator) SetTimeFrame(timeFrame types.TimeFrame) {
	i.TimeFrame = timeFrame
}
