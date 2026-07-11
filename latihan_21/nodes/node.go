package nodes

type Node struct {
	Name        string
	MsgSent     string
	MsgReceived map[string]string
	Peers       map[string]*Node
}

func (n *Node) SetName(name string) {
	n.Name = name
}

func (n *Node) GetName() string {
	return n.Name
}

func (n *Node) Link(m *Node) {
	if n == nil || m == nil {
		return
	}

	// Peers dibuat saat pertama kali node punya tetangga.
	// Ini membuat pembuatan node baru tetap sederhana: Peers boleh nil dulu.
	if n.Peers == nil {
		n.Peers = make(map[string]*Node)
	}
	n.Peers[m.Name] = m
}
