package trade_vessel

type Kline struct {
	open   float64
	high   float64
	low    float64
	close  float64
	volume float64
}

func NewKline(open, high, low, close, volume float64) *Kline {
	return &Kline{
		open:   open,
		high:   high,
		low:    low,
		close:  close,
		volume: volume,
	}
}
