package strategy

type CrossUp struct {
	FirstValues  []float64
	SecondValues []float64
}

func (t *CrossUp) CanEnter() bool {
	firstValues := t.FirstValues
	secondValues := t.SecondValues
	if firstValues[len(firstValues)-1] > secondValues[len(secondValues)-1] &&
		firstValues[len(firstValues)-2] < secondValues[len(secondValues)-2] {
		return true
	}
	return false
}
