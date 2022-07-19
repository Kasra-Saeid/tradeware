package trade_middleware

type Indicator struct {
	Name      IndicatorName
	Settings  IndicatorSettings
	TimeFrame TimeFrame
	Value     []float64
}

func (i *Indicator) GetValue() []float64 {
	return i.Value
}

func (i *Indicator) GetName() IndicatorName {
	return i.Name
}

func (i *Indicator) GetSettings() IndicatorSettings {
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

func (i *Indicator) GetTimeFrame() TimeFrame {
	return i.TimeFrame
}

func (i *Indicator) SetValue(value []float64) {
	i.Value = value
}

func (i *Indicator) SetName(name IndicatorName) {
	i.Name = name
}

func (i *Indicator) SetSettings(source *Source, attrs ...IndicatorSettingsAttr) {
	switch i.Settings.(type) {
	case *AdxSettings:
		for _, v := range attrs {
			if v.Attr == AdxLength {
				i.Settings.(*AdxSettings).AdxLength = (int)(v.Value)
			}
		}
	case *EmaSettings:
		for _, v := range attrs {
			if v.Attr == EmaLength {
				i.Settings.(*EmaSettings).EmaLength = (int)(v.Value)
			}
		}
	case *AtrSettings:
		for _, v := range attrs {
			if v.Attr == AtrLength {
				i.Settings.(*AtrSettings).AtrLength = (int)(v.Value)
			}
		}
	case *KeltnerWidthSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == EmaLength {
				i.Settings.(*KeltnerWidthSettings).EmaLength = (int)(v.Value)
			} else if v.Attr == AtrLength {
				i.Settings.(*KeltnerWidthSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == Multiplier {
				i.Settings.(*KeltnerWidthSettings).Multiplier = (int)(v.Value)
			}
		}
		i.Settings.(*KeltnerWidthSettings).Source = *source

	case *SuperTrendSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == AtrLength {
				i.Settings.(*SuperTrendSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == Multiplier {
				i.Settings.(*SuperTrendSettings).Multiplier = v.Value
			}
		}
		i.Settings.(*SuperTrendSettings).Source = *source

	}
}

func (i *Indicator) SetTimeFrame(timeFrame TimeFrame) {
	i.TimeFrame = timeFrame
}
