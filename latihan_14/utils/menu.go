package utils

import (
	"bufio"
	"fmt"
	"networks/graph"
	"networks/nodes"
	"os"
	"strconv"
	"strings"
)

func readText(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt(reader *bufio.Reader, prompt string) int {
	for {
		text := readText(reader, prompt)
		value, err := strconv.Atoi(text)
		if err == nil {
			return value
		}

		fmt.Println("Input harus berupa angka.")
	}
}

func Menu() {
	g := &graph.Graph{}
	s := &graph.State{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\n Implementasi Koding Graf  \n\n")
		fmt.Printf("\n _________________________  \n\n")

		fmt.Printf("\nMenu : \n\n")
		fmt.Printf("\ta. Tambah Node \n")
		fmt.Printf("\tb. Hubungkan Node\n")
		fmt.Printf("\tc. Hapus Node\n")
		fmt.Printf("\td. Traverse\n")
		fmt.Printf("\ts. Cari Jalur Terpendek\n")
		fmt.Printf("\te. Tampilkan Graph\n")
		fmt.Printf("\tq. Keluar\n\n")

		choice := strings.ToLower(readText(reader, "Pilihan : "))

		switch choice {
		case "a":
			nodeCount := readInt(reader, "\nMasukkan berapa node yang ingin dibuat: ")

			for i := 1; i <= nodeCount; i++ {
				id := nextNodeID(g)
				nodeWeight := readInt(reader, fmt.Sprintf("Masukkan bobot node id-%d: ", id))
				*s = *addNode(id, nodeWeight, nil, g)
			}

		case "b":
			fromID := readInt(reader, "\nMasukkan id node asal: ")
			toID := readInt(reader, "Masukkan id node tujuan: ")

			fromNode := g.Nodes[fromID]
			toNode := g.Nodes[toID]
			if fromNode == nil || toNode == nil {
				fmt.Println("Node asal atau tujuan tidak ditemukan.")
				continue
			}

			g.LinkNode(fromNode, toNode)
			s.Current = fromNode
			s.Next = fromNode.GetAdj()
			fmt.Printf("Node %d berhasil dihubungkan ke node %d.\n", fromID, toID)

		case "c":
			id := readInt(reader, "\nMasukkan id node yang ingin dihapus: ")
			*s = *removeNode(id, g, s)

		case "d":
			*s = *g.Traverse(s)
			if s.Current == nil {
				fmt.Println("Tidak ada node berikutnya untuk dikunjungi.")
				continue
			}
			fmt.Printf("Sekarang berada di node %d.\n", s.Current.GetId())

		case "s":
			id := readInt(reader, "\nMasukkan id node yang ingin dikunjungi : ")
			res := g.ShortestPath(s, id)
			for k := range res {
				fmt.Printf("Node path\n-> %d ", k)
			}
			fmt.Printf("\n")

		case "e":
			printGraph(g)

		case "q":
			return

		default:
			fmt.Println("Pilihan tidak dikenal.")
		}
	}
}

func addNode(id, weight int, adj map[int]*nodes.Node, g *graph.Graph) *graph.State {
	add := &nodes.Node{

		Id:     id,
		Weight: weight,
		Adj:    adj,
	}

	return g.AddNode(add)
}

func removeNode(id int, g *graph.Graph, s *graph.State) *graph.State {
	if g.Nodes == nil {
		return s
	}

	return g.RemoveNode(g.Nodes[id], s)

}

func nextNodeID(g *graph.Graph) int {
	nextID := 1
	for id := range g.Nodes {
		if id >= nextID {
			nextID = id + 1
		}
	}
	return nextID
}

func printGraph(g *graph.Graph) {
	if len(g.Nodes) == 0 {
		fmt.Println("Graph masih kosong.")
		return
	}

	for id, node := range g.Nodes {
		fmt.Printf("Node %d, bobot %d", id, node.GetWeight())
		if len(node.GetAdj()) == 0 {
			fmt.Println(", tetangga: -")
			continue
		}

		fmt.Print(", tetangga:")
		for adjID := range node.GetAdj() {
			fmt.Printf(" %d", adjID)
		}
		fmt.Println()
	}
}
