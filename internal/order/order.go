package order

type Order struct {
	EntryPrice        float64
	StopLossPrice     float64
	TakeProfitTrigger bool
	TakeProfitUnit    float64
	Qty               float64
	Side              string
	Symbol            string
	Leverage          float64
}
