package graph

import "sort"

type State struct {
	Now     *Node
	Next    *Node
	Visited map[int]bool
}

func (s *State) Traverse(n *Node) *Node {
	if s == nil {
		return nil
	}

	if s.Visited == nil {
		s.Visited = make(map[int]bool)
	}

	current := n
	if s.Now != nil && *s.Now != nil {
		current = s.Now
	}

	if current == nil || *current == nil {
		return nil
	}

	currentID := (*current).GetID()
	if currentID != 0 {
		s.Visited[currentID] = true
	}

	adj := (*current).DisplayAdj()
	if len(adj) == 0 {
		s.Next = nil
		return nil
	}

	ids := make([]int, 0, len(adj))
	for id := range adj {
		ids = append(ids, id)
	}
	sort.Ints(ids)

	for _, id := range ids {
		if s.Visited[id] {
			continue
		}

		nextNode := adj[id]
		s.Next = &nextNode
		s.Now = s.Next
		return s.Next
	}

	s.Next = nil
	return nil
}
