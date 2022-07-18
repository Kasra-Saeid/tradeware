package rules

type CrossDown struct {
	FirstValues  []float64
	SecondValues []float64
	IndexBar     int
}

func (t *CrossDown) Check() bool {
	firstValues := t.FirstValues
	secondValues := t.SecondValues
	if firstValues[len(firstValues)-t.IndexBar] < secondValues[len(secondValues)-t.IndexBar] &&
		firstValues[len(firstValues)-(t.IndexBar-1)] > secondValues[len(secondValues)-(t.IndexBar-1)] {
		return true
	}
	return false
}
