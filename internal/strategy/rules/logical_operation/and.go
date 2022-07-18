package logical_operation

import (
	"tradechef_backtest/internal/strategy/rules"
)

type And struct {
	rules []rules.Rules
}

func (o *And) AddRule(rule rules.Rules) {
	o.rules = append(o.rules, rule)
}
func (o *And) Check() bool {
	for _, v := range o.rules {
		if !v.Check() {
			return false
		}
	}
	return true
}
