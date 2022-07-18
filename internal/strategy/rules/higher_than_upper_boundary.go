package rules

type HigherThanUpperBoundary struct {
	Values        []float64
	UpperBoundary float64
	IndexBar      int
}

func (l *HigherThanUpperBoundary) Check() bool {
	latestValue := l.Values[len(l.Values)-l.IndexBar]
	if latestValue > l.UpperBoundary {
		return true
	}
	return false
}
