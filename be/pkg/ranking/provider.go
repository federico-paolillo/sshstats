package ranking

type RankProvider interface {
	Top15(nodeName string) []*Placement
}
