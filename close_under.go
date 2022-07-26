package tradeware

type CloseUnder struct {
	firstValues  []float64
	secondValues []float64
	indexBar     int
}

func NewCloseUnder(firstValues, SecondValues []float64, indexBar int) *CloseUnder {
	return &CloseUnder{
		firstValues:  firstValues,
		secondValues: SecondValues,
		indexBar:     indexBar,
	}
}

func (c *CloseUnder) Check() bool {
	firstValues := c.firstValues
	secondValues := c.secondValues
	if firstValues[len(firstValues)-c.indexBar] < secondValues[len(secondValues)-c.indexBar] {
		return true
	}
	return false
}
