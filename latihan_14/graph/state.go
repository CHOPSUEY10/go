package graph

import (
	"networks/nodes"
)

type State struct {
	Current *nodes.Node
	Next    map[int]*nodes.Node
}

type Adj struct {
	Vertices *nodes.Node
	Cost     int
}

func (s *State) Traverse(n *nodes.Node) *nodes.Node {
	if n == nil {
		return nil
	}

	var next *nodes.Node
	var leastCost int
	firstNode := true

	for _, v := range n.GetAdj() {
		if v == nil {
			continue
		}

		// Cost di sini mengikuti rumus awalmu:
		// bobot node saat ini - bobot node tetangga.
		cost := n.GetWeight() - v.GetWeight()
		if firstNode || cost < leastCost {
			leastCost = cost
			next = v
			firstNode = false
		}
	}

	return next

}
