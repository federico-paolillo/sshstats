package stats

type Statistician interface {
	Top15LoginAttemps(nodeName string) []*LoginAttempt
	Last10Attackers() []*Attacker
}
