package main

import (
	"tradechef_backtest/internal/entities"
)

func main() {
	adx := entities.NewAdx()
	var settings = adx.GetSettings().(entities.AdxSettings)
	print(settings.AdxLength)
}
