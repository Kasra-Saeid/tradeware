package order

import "tradechef_backtest/types"

type Order struct {
	EntryPrice        float64
	StopLossPrice     float64
	TakeProfitTrigger bool
	TakeProfitUnit    float64
	Qty               float64
	Side              types.Side
	Symbol            string
	Leverage          float64
}
