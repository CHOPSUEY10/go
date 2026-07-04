package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"grafers/graph"
	"grafers/node"
)

func readInt(reader *bufio.Reader) int {
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	value, _ := strconv.Atoi(text)
	return value
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jumlah node: ")
	nodeCount := readInt(reader)

	gr := &graph.Graph{}
	for i := 1; i <= nodeCount; i++ {
		gr.AddNode(&node.Person{ID: i})
	}

	fmt.Println("Masukkan adjacency setiap node (pisahkan dengan koma, contoh: 2,3).")
	for i := 1; i <= nodeCount; i++ {
		fmt.Printf("Adjacency node %d: ", i)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}

			adjID, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Input tidak valid: %s\n", part)
				continue
			}

			fromNode, okFrom := gr.Nodes[i]
			toNode, okTo := gr.Nodes[adjID]
			if okFrom && okTo {
				gr.LinkNode(fromNode, toNode)
			}
		}
	}

	fmt.Println("\nGraph adjacency:")
	gr.ShowGraph()

	fmt.Print("Masukkan node awal: ")
	startID := readInt(reader)
	fmt.Print("Masukkan jumlah langkah: ")
	steps := readInt(reader)

	startNode, ok := gr.Nodes[startID]
	if !ok {
		fmt.Printf("Node %d tidak ditemukan\n", startID)
		return
	}

	state := graph.State{}
	current := startNode
	fmt.Printf("\nMulai dari node: %d\n", current.GetID())

	for i := 0; i < steps; i++ {
		nextNode := state.Traverse(&current)
		if nextNode == nil {
			fmt.Println("Tidak ada jalur lagi.")
			break
		}

		fmt.Printf("Langkah %d -> node: %d\n", i+1, (*nextNode).GetID())
		current = *nextNode
	}
}
