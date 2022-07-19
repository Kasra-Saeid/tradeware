package trade_vessel

type And struct {
	rules []Rules
}

func NewAnd() *And {
	return &And{
		rules: make([]Rules, 0, 0),
	}
}

func (o *And) AddRule(rule Rules) {
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
