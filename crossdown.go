package trade_vessel

type CrossDown struct {
	firstValues  []float64
	secondValues []float64
	indexBar     int
}

func NewCrossDown(firstValues, SecondValues []float64, indexBar int) *CrossDown {
	return &CrossDown{
		firstValues:  firstValues,
		secondValues: SecondValues,
		indexBar:     indexBar,
	}
}

func (t *CrossDown) Check() bool {
	firstValues := t.firstValues
	secondValues := t.secondValues
	if firstValues[len(firstValues)-t.indexBar] < secondValues[len(secondValues)-t.indexBar] &&
		firstValues[len(firstValues)-(t.indexBar-1)] > secondValues[len(secondValues)-(t.indexBar-1)] {
		return true
	}
	return false
}
