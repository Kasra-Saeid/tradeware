package strategy

type Strategy struct {
	EntryRules      []EntryRules
	TakeProfitRules []TakeProfitRules
	StopLossRules   []StopLossRules
}

func (s *Strategy) AddEntryRule(er EntryRules) {
	s.EntryRules = append(s.EntryRules, er)
}

func (s *Strategy) AddTakeProfitRule(tpr TakeProfitRules) {
	s.TakeProfitRules = append(s.TakeProfitRules, tpr)
}

func (s *Strategy) AddStopLossRule(slr StopLossRules) {
	s.StopLossRules = append(s.StopLossRules, slr)
}

func (s *Strategy) CanEnter() bool {
	for _, v := range s.EntryRules {
		if !v.CanEnter() {
			return false
		}
	}
	return true
}
