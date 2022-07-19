package trade_middleware

type Order struct {
	entryPrice        float64
	stopLossPrice     float64
	takeProfitTrigger bool
	takeProfitUnit    float64
	qty               float64
	side              Side
	symbol            string
	leverage          float64
	status            Status
}

func NewOrder(ep, sl float64, tpUnit, qty float64, side Side, symbol string, leverage float64) *Order {
	return &Order{
		entryPrice:        0,
		stopLossPrice:     0,
		takeProfitTrigger: false,
		takeProfitUnit:    0,
		qty:               0,
		side:              "",
		symbol:            "",
		leverage:          0,
		status:            Available,
	}
}

func (o *Order) OpenOrder(ep, sl float64, tpUnit, qty float64, side Side, symbol string, leverage float64) {
	o.entryPrice = ep
	o.stopLossPrice = sl
	o.takeProfitTrigger = false
	o.takeProfitUnit = tpUnit
	o.qty = qty
	o.side = side
	o.symbol = symbol
	o.leverage = leverage
	o.status = InProgress
}

func (o *Order) CloseOrder() {
	o.entryPrice = 0
	o.stopLossPrice = 0
	o.takeProfitTrigger = false
	o.takeProfitUnit = 0
	o.qty = 0
	o.side = ""
	o.symbol = ""
	o.leverage = 0
	o.status = Available
}

func (o *Order) OrderTpTrigger() {
	o.takeProfitTrigger = true
}
