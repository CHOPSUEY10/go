package graph

type Graph struct {
	Nodes  map[int]Node
	Weight int
}

func (gr *Graph) AddNode(n Node) {
	if gr.Nodes == nil {
		gr.Nodes = make(map[int]Node)
	}

	gr.Nodes[n.GetID()] = n

}

func (gr *Graph) RemoveNode(n Node) {
	delete(gr.Nodes, n.GetID())
}

func (gr *Graph) LinkNode(n Node, m Node) {
	n.Link(m)
}
