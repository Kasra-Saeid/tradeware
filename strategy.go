package tradeware

type Strategy struct {
	entryRules      []Rules
	takeProfitRules []Rules
	stopLossRules   []Rules
}

func NewStrategy(entryRules, takeProfitRules, stopLossRules []Rules) *Strategy {
	return &Strategy{
		entryRules:      entryRules,
		takeProfitRules: takeProfitRules,
		stopLossRules:   stopLossRules,
	}
}

func (s *Strategy) AddEntryRule(er Rules) {
	s.entryRules = append(s.entryRules, er)
}

func (s *Strategy) AddTakeProfitRule(tpr Rules) {
	s.takeProfitRules = append(s.takeProfitRules, tpr)
}

func (s *Strategy) AddStopLossRule(slr Rules) {
	s.stopLossRules = append(s.stopLossRules, slr)
}

func (s *Strategy) CheckEntry() bool {
	for _, v := range s.entryRules {
		if !v.Check() {
			return false
		}
	}
	return true
}

func (s *Strategy) CheckTakeProfit() bool {
	for _, v := range s.takeProfitRules {
		if !v.Check() {
			return false
		}
	}
	return true
}

func (s *Strategy) CheckSlHit() bool {
	for _, v := range s.stopLossRules {
		if !v.Check() {
			return false
		}
	}
	return true
}
