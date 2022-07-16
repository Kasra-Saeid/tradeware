package entities

import (
	"tradechef_backtest/constants"
	"tradechef_backtest/types"
)

type Indicator struct {
	Name     types.IndicatorName
	Settings types.IndicatorSettings
	Value    []float64
}

func (i *Indicator) GetValue() []float64 {
	return i.Value
}

func (i *Indicator) GetName() types.IndicatorName {
	return i.Name
}

func (i *Indicator) GetSettings() types.IndicatorSettings {
	return i.Settings
}

func (i *Indicator) SetValue(value []float64) {
	i.Value = value
}

func (i *Indicator) SetName(name types.IndicatorName) {
	i.Name = name
}

func (i *Indicator) SetSettings(values ...types.IndicatorSettingsAttr) {
	switch i.Settings.(type) {
	case AdxSettings:
		for _, v := range values {
			if v.Name == constants.AdxLength {
				i.Settings = &AdxSettings{
					AdxLength: v.Value,
				}
				break
			}
		}

	}

}
