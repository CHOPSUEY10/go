package main

import (
	"fmt"
	"grafers/graph"
	"grafers/node"
)

func main() {
	gr := &graph.Graph{}

	p := &node.Person{ID: 1}
	gr.AddNode(p)

	g := &node.Ghost{ID: 2}
	gr.AddNode(g)

	n := &node.Person{ID: 3}
	gr.AddNode(n)

	gr.LinkNode(p, g)
	gr.LinkNode(g, p)
	gr.LinkNode(p, n)
	gr.LinkNode(n, g)

	for _, v := range gr.Nodes {

		id := v.GetID()
		p := v.DisplayAdj()
		fmt.Println(id, p)
	}

}
