package stats

type Attempts = []*LoginAttempt
type Attackers = []*Attacker

type Provider interface {
	Top15LoginAttempts(nodeName string) (Attempts, error)
	Last10Attackers() (Attackers, error)
}
