package node

import "grafers/graph"

type Ghost struct {
	ID  int
	Adj map[int]graph.Node
}

func (g *Ghost) DisplayAdj() map[int]graph.Node {
	if g == nil {
		return nil
	}

	return g.Adj
}

func (g *Ghost) GetID() int {
	if g == nil {
		return 0
	}

	return g.ID
}

func (g *Ghost) Link(n graph.Node) {
	if g == nil || n == nil {
		return
	}

	if g.Adj == nil {
		g.Adj = make(map[int]graph.Node)
	}

	g.Adj[n.GetID()] = n
}
