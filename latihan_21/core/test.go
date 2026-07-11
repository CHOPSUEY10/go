package main

import (
	"fmt"
	"latihan_21/connections"
	"latihan_21/nodes"
	"latihan_21/states"
)

func main() {
	conn := &connections.Connection{}
	state := &states.State{}

	nama := []string{"fadli", "putri", "siti", "haaland", "gibran"}
	conn.AllNodes = make(map[string]*nodes.Node, len(nama))

	for _, name := range nama {
		node := &nodes.Node{}
		node.SetName(name)
		conn.AllNodes[name] = node
	}

	conn.ConnectAll(state)

	// for _, name := range nama {
	// 	person := conn.AllNodes[name]
	// 	fmt.Printf("Temannya %s : ", person.GetName())
	// 	for _, v := range person.Peers {
	// 		fmt.Printf("%s ", v.GetName())
	// 	}
	// 	fmt.Println()
	// }
	msgChan := make(chan string, 1)
	msgChan <- "Hidup Pria Oslo"
	sender := "fadli"
	go conn.SendMessageToAll(msgChan, sender)

	for _, name := range nama {

		if sender == name {
			continue
		}
		person := conn.AllNodes[name]
		fmt.Printf("Pesan yang diterima %s :", person.GetName())
		for _, msg := range person.MsgReceived {
			fmt.Printf("%s\n", msg)

		}
	}
}
