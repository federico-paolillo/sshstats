package attackers

type AttackersProvider interface {
	Last10() []*Attacker
}
