package trade_middleware

type HigherThanUpperBoundary struct {
	values        []float64
	upperBoundary float64
	indexBar      int
}

func NewHigherThanUpperBoundary(values []float64, upperBoundary float64, indexBar int) *HigherThanUpperBoundary {
	return &HigherThanUpperBoundary{
		values:        values,
		upperBoundary: upperBoundary,
		indexBar:      indexBar,
	}
}
func (l *HigherThanUpperBoundary) Check() bool {
	latestValue := l.values[len(l.values)-l.indexBar]
	if latestValue > l.upperBoundary {
		return true
	}
	return false
}
