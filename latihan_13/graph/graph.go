package graph

import "fmt"

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

func (gr *Graph) ShowGraph() {
	for id, n := range gr.Nodes {
		adj := n.DisplayAdj()
		x := len(adj)
		fmt.Println("Node:", id, "jumlah link:", x)
		for _, a := range adj {
			fmt.Printf("\n\tTetangga %d\n\n", a.GetID())
		}
	}
}

func (gr *Graph) FirstNode() Node {
	return gr.Nodes[1]
}
