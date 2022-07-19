package strategy

import "tradechef_backtest/internal/strategy/rules"

type Strategy struct {
	entryRules      []rules.Rules
	takeProfitRules []rules.Rules
	stopLossRules   []rules.Rules
}

func NewStrategy(entryRules, takeProfitRules, stopLossRules []rules.Rules) *Strategy {
	return &Strategy{
		entryRules:      entryRules,
		takeProfitRules: takeProfitRules,
		stopLossRules:   stopLossRules,
	}
}

func (s *Strategy) AddEntryRule(er rules.Rules) {
	s.entryRules = append(s.entryRules, er)
}

func (s *Strategy) AddTakeProfitRule(tpr rules.Rules) {
	s.takeProfitRules = append(s.takeProfitRules, tpr)
}

func (s *Strategy) AddStopLossRule(slr rules.Rules) {
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
