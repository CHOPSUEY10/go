package node

import "grafers/graph"

type Person struct {
	ID  int
	Adj map[int]graph.Node
}

func (p *Person) DisplayAdj() map[int]graph.Node {

	return p.Adj

}

func (p *Person) GetID() int {

	return p.ID

}

func (p *Person) Link(n graph.Node) {

	if p.Adj == nil {
		p.Adj = make(map[int]graph.Node)
	}

	p.Adj[n.GetID()] = n
}
