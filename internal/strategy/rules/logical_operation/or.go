package logical_operation

import (
	"tradechef_backtest/internal/strategy/rules"
)

type Or struct {
	rules []rules.Rules
}

func (o *Or) AddRule(rule rules.Rules) {
	o.rules = append(o.rules, rule)
}
func (o *Or) Check() bool {
	for _, v := range o.rules {
		if v.Check() {
			return true
		}
	}
	return false
}
