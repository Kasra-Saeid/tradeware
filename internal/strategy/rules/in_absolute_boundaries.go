package rules

type InAbsoluteBoundaries struct {
	UpperBoundary float64
	LowerBoundary float64
	Values        []float64
}

func (i *InAbsoluteBoundaries) Check() bool {
	var startIndex int

	for index := len(i.Values) - 1; index >= 0; index-- {
		if i.Values[index] < i.UpperBoundary && i.Values[index] >= i.LowerBoundary && i.Values[index-1] < i.LowerBoundary {
			startIndex = index
			break
		}
	}
	if startIndex == 0 {
		return false
	}
	for index := len(i.Values) - 1; index >= startIndex; index-- {
		if !(i.Values[index] < i.UpperBoundary && i.Values[index] > i.LowerBoundary) {
			return false
		}
	}
	return true
}
