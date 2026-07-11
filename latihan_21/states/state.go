package states

import "latihan_21/nodes"

type State struct {
	being *nodes.Node
}

func (b *State) SetState(node *nodes.Node) {

	b.being = node

}

func (b *State) GetState() *nodes.Node {

	return b.being

}
