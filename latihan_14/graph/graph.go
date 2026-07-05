package graph

import (
	"networks/nodes"
)

type Graph struct {
	Nodes map[int]*nodes.Node
}

func (g *Graph) AddNode(n *nodes.Node) (s *State) {
	if n == nil {
		return &State{}
	}

	// Map harus dibuat sebelum diisi. Zero value map adalah nil dan akan panic
	// jika langsung dipakai untuk assignment seperti g.Nodes[id] = node.
	if g.Nodes == nil {
		g.Nodes = make(map[int]*nodes.Node)
	}

	g.Nodes[n.GetId()] = n
	return &State{Current: n, Next: n.GetAdj()}
}

func (g *Graph) RemoveNode(n *nodes.Node, s *State) *State {
	if s == nil {
		s = &State{}
	}
	if n == nil {
		return s
	}

	// Jika node yang dihapus adalah posisi saat ini, pindahkan state dulu.
	// Kalau tidak ada tetangga, Current akan menjadi nil.
	if s.Current == n {
		s.Current = s.Traverse(n)
		if s.Current != nil {
			s.Next = s.Current.GetAdj()
		} else {
			s.Next = nil
		}
	}

	delete(g.Nodes, n.GetId())

	// Hapus juga edge yang mengarah ke node ini agar adjacency tidak menyimpan
	// referensi ke node yang sudah tidak ada di graph.
	for _, node := range g.Nodes {
		delete(node.GetAdj(), n.GetId())
	}

	return s
}

func (g *Graph) Traverse(s *State) *State {
	if s == nil {
		s = &State{}
	}
	if s.Current == nil {
		s.Next = nil
		return s
	}

	s.Current = s.Traverse(s.Current)
	if s.Current != nil {
		s.Next = s.Current.GetAdj()
	} else {
		s.Next = nil
	}

	return s

}

func (g *Graph) LinkNode(n, m *nodes.Node) {
	if n == nil || m == nil {
		return
	}

	n.Link(m)
}

func (g *Graph) ShortestPath(f *State, t int) (res map[int]*nodes.Node) {

	state := true
	res = make(map[int]*nodes.Node)
	var curr *State
	for state {
		curr = g.Traverse(f)
		res[curr.Current.GetId()] = curr.Current
		if curr.Current.GetId() == t {
			break
		}

	}
	return res

}
