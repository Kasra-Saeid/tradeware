package tradeware

type FirstUnderSecond struct {
	firstValues  []float64
	secondValues []float64
	indexBar     int
}

func NewFirstUnderSecond(firstValues, SecondValues []float64, indexBar int) *FirstUnderSecond {
	return &FirstUnderSecond{
		firstValues:  firstValues,
		secondValues: SecondValues,
		indexBar:     indexBar,
	}
}

func (c *FirstUnderSecond) Check() bool {
	firstValues := c.firstValues
	secondValues := c.secondValues
	if firstValues[len(firstValues)-c.indexBar] < secondValues[len(secondValues)-c.indexBar] {
		return true
	}
	return false
}
