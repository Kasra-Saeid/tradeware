package trade_middleware

type CrossUp struct {
	firstValues  []float64
	secondValues []float64
	indexBar     int
}

func NewCrossUp(firstValues, SecondValues []float64, indexBar int) *CrossUp {
	return &CrossUp{
		firstValues:  firstValues,
		secondValues: SecondValues,
		indexBar:     indexBar,
	}
}

func (t *CrossUp) Check() bool {
	firstValues := t.firstValues
	secondValues := t.secondValues
	if firstValues[len(firstValues)-t.indexBar] > secondValues[len(secondValues)-t.indexBar] &&
		firstValues[len(firstValues)-(t.indexBar-1)] < secondValues[len(secondValues)-(t.indexBar-1)] {
		return true
	}
	return false
}
