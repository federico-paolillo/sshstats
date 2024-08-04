package stats

type Attempts = []*LoginAttempt
type Attackers = []*Attacker

type Statistician interface {
	Top15LoginAttemps(nodeName string) (Attempts, error)
	Last10Attackers() (Attackers, error)
}
