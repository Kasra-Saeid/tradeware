package indicator

import (
	"github.com/kasrasaeed/trade_vessel/constants"
	indicators2 "github.com/kasrasaeed/trade_vessel/indicator/indicators"
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
	case *indicators2.AdxSettings:
		return i.Settings.(*indicators2.AdxSettings)
	case *indicators2.EmaSettings:
		return i.Settings.(*indicators2.EmaSettings)
	case *indicators2.AtrSettings:
		return i.Settings.(*indicators2.AtrSettings)
	case *indicators2.KeltnerWidthSettings:
		return i.Settings.(*indicators2.KeltnerWidthSettings)
	case *indicators2.SuperTrendSettings:
		return i.Settings.(*indicators2.SuperTrendSettings)
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
	case *indicators2.AdxSettings:
		for _, v := range attrs {
			if v.Attr == constants.AdxLength {
				i.Settings.(*indicators2.AdxSettings).AdxLength = (int)(v.Value)
			}
		}
	case *indicators2.EmaSettings:
		for _, v := range attrs {
			if v.Attr == constants.EmaLength {
				i.Settings.(*indicators2.EmaSettings).EmaLength = (int)(v.Value)
			}
		}
	case *indicators2.AtrSettings:
		for _, v := range attrs {
			if v.Attr == constants.AtrLength {
				i.Settings.(*indicators2.AtrSettings).AtrLength = (int)(v.Value)
			}
		}
	case *indicators2.KeltnerWidthSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == constants.EmaLength {
				i.Settings.(*indicators2.KeltnerWidthSettings).EmaLength = (int)(v.Value)
			} else if v.Attr == constants.AtrLength {
				i.Settings.(*indicators2.KeltnerWidthSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == constants.Multiplier {
				i.Settings.(*indicators2.KeltnerWidthSettings).Multiplier = (int)(v.Value)
			}
		}
		i.Settings.(*indicators2.KeltnerWidthSettings).Source = *source

	case *indicators2.SuperTrendSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == constants.AtrLength {
				i.Settings.(*indicators2.SuperTrendSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == constants.Multiplier {
				i.Settings.(*indicators2.SuperTrendSettings).Multiplier = v.Value
			}
		}
		i.Settings.(*indicators2.SuperTrendSettings).Source = *source

	}
}

func (i *Indicator) SetTimeFrame(timeFrame types.TimeFrame) {
	i.TimeFrame = timeFrame
}
