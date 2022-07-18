package strategy

import "tradechef_backtest/internal/strategy/rules"

type Strategy struct {
	EntryRules      []rules.Rules
	TakeProfitRules []rules.Rules
	StopLossRules   []rules.Rules
}

func (s *Strategy) AddEntryRule(er rules.Rules) {
	s.EntryRules = append(s.EntryRules, er)
}

func (s *Strategy) AddTakeProfitRule(tpr rules.Rules) {
	s.TakeProfitRules = append(s.TakeProfitRules, tpr)
}

func (s *Strategy) AddStopLossRule(slr rules.Rules) {
	s.StopLossRules = append(s.StopLossRules, slr)
}

func (s *Strategy) CheckEntry() bool {
	for _, v := range s.EntryRules {
		if !v.Check() {
			return false
		}
	}
	return true
}

func (s *Strategy) CheckTakeProfit() bool {
	for _, v := range s.TakeProfitRules {
		if !v.Check() {
			return false
		}
	}
	return true
}

func (s *Strategy) CheckSlHit() bool {
	for _, v := range s.StopLossRules {
		if !v.Check() {
			return false
		}
	}
	return true
}
