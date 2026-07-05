package nodes

type Node struct {
	Id, Weight int
	Adj        map[int]*Node
}

func (n *Node) GetAdj() map[int]*Node {
	return n.Adj
}
func (n *Node) GetWeight() int {
	return n.Weight
}
func (n *Node) GetId() int {
	return n.Id
}

func (n *Node) SetAdj(adj map[int]*Node) {
	n.Adj = adj
}
func (n *Node) SetWeight(weight int) {
	n.Weight = weight
}
func (n *Node) SetId(id int) {
	n.Id = id
}

func (n *Node) Link(m *Node) {
	if n == nil || m == nil {
		return
	}

	// Adj dibuat saat pertama kali node punya tetangga.
	// Ini membuat pembuatan node baru tetap sederhana: Adj boleh nil dulu.
	if n.Adj == nil {
		n.Adj = make(map[int]*Node)
	}
	n.Adj[m.GetId()] = m
}
