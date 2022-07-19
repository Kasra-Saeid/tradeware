package trade_vessel

type IIndicator interface {
	GetName() IndicatorName
	GetSettings() IndicatorSettings
	GetTimeFrame() TimeFrame
	GetValue() []float64
	SetName(name IndicatorName)
	SetValue(value []float64)
	SetSettings(source *Source, values ...IndicatorSettingsAttr)
	SetTimeFrame(timeFrame TimeFrame)
}

type Rules interface {
	Check() bool
}
