package tradeware

type FirstAboveSecond struct {
	firstValues  []float64
	secondValues []float64
	indexBar     int
}

func NewCloseAbove(firstValues, SecondValues []float64, indexBar int) *FirstAboveSecond {
	return &FirstAboveSecond{
		firstValues:  firstValues,
		secondValues: SecondValues,
		indexBar:     indexBar,
	}
}

func (c *FirstAboveSecond) Check() bool {
	firstValues := c.firstValues
	secondValues := c.secondValues
	if firstValues[len(firstValues)-c.indexBar] > secondValues[len(secondValues)-c.indexBar] {
		return true
	}
	return false
}
