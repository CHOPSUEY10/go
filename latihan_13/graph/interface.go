package graph

type Node interface {
	GetID() int
	DisplayAdj() map[int]Node
	Link(Node)
}
