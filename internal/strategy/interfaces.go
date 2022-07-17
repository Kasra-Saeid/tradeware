package strategy

type EntryRules interface {
	CanEnter() bool
}

type TakeProfitRules interface {
	CanTakeProfit() bool
}

type StopLossRules interface {
	IsSlHit() bool
}
