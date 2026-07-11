package connections

import (
	"errors"
	"latihan_21/nodes"
	"latihan_21/states"
	"sync"
)

type Connection struct {
	AllNodes map[string]*nodes.Node
}

func (person *Connection) SendMessageToOne(msg chan string, s, r string) (string, error) {

	sender := person.AllNodes[s]
	receiver := person.AllNodes[r]

	if receiver.Peers[sender.GetName()] == sender {

		receiver.MsgReceived[sender.GetName()] = <-msg
		close(msg)
		return "Succesfully sent", nil
	}
	return "", errors.New("Cannot find any peers you want to send")
}

func (person *Connection) SendMessageToAll(msg chan string, s string) {

	sender := person.AllNodes[s]
	sender.MsgSent = <-msg

	for _, v := range person.AllNodes {
		receiver := person.AllNodes[v.GetName()]
		if receiver.MsgReceived == nil {
			receiver.MsgReceived = make(map[string]string)
		}
		if v.GetName() == sender.GetName() {
			continue
		}
		if receiver.Peers[sender.GetName()] == sender {
			receiver.MsgReceived[v.GetName()] = sender.MsgSent
		}
	}
	close(msg)

}

func (person *Connection) ConnectAll(st *states.State) {
	if len(person.AllNodes) == 0 {
		return
	}
	nodesList := make([]*nodes.Node, 0, len(person.AllNodes))
	for _, v := range person.AllNodes {
		nodesList = append(nodesList, v)
	}

	var wg sync.WaitGroup
	for _, current := range nodesList {
		current := current
		wg.Add(1)
		go func() {
			defer wg.Done()
			st.SetState(current)
			for _, other := range nodesList {
				if current == other {
					continue
				}
				current.Link(other)
			}
		}()
	}
	wg.Wait()
}
