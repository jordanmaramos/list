package list

import "fmt"

type DoublyLinkedList struct {
	head      *NodeDoubly
	tail      *NodeDoubly
	lastIndex int
}

type NodeDoubly struct {
	previous *NodeDoubly
	next     *NodeDoubly
	value    int
}

func (doublyLinkedList *DoublyLinkedList) Init() {
	doublyLinkedList.head = nil
	doublyLinkedList.tail = nil
	doublyLinkedList.lastIndex = -1
}

func (doublyLinkedList *DoublyLinkedList) Add(value int) {
	newNode := NodeDoubly{value: value, next: nil, previous: nil}

	if doublyLinkedList.head == nil {
		doublyLinkedList.head = &newNode //Head
		doublyLinkedList.tail = &newNode //Tail = Head
	} else {
		lastNode := doublyLinkedList.tail

		if lastNode.previous == nil {
			//If tail.previous is nil, lastNode is the head
			doublyLinkedList.head.next = &newNode    //First tail
			newNode.previous = doublyLinkedList.tail //Previous points to head/tail
			doublyLinkedList.tail = &newNode         //Tail is now the newNode
		} else {
			//If not, there is more than 1 element
			newNode.previous = doublyLinkedList.tail //newNode.previous points to the tai
			doublyLinkedList.tail.next = &newNode    //tail.next points to newNode
			doublyLinkedList.tail = &newNode         //Replace tail with newNode, the new tail
		}
	}

	//Increment last index
	doublyLinkedList.lastIndex++
}

func (doublyLinkedList *DoublyLinkedList) AddOnIndex(value, index int) error {
	newNode := NodeDoubly{value: value, next: nil, previous: nil}
	if index < 0 || index > doublyLinkedList.lastIndex {
		return fmt.Errorf("Index inválido.")
	}

	if index == 0 {
		//For replacing first element
		newNode.next = doublyLinkedList.head
		doublyLinkedList.head.previous = &newNode
		doublyLinkedList.head = &newNode
	} else if index == doublyLinkedList.lastIndex {
		//For replacing last element
		newNode.previous = doublyLinkedList.tail.previous
		doublyLinkedList.tail.previous.next = &newNode
		newNode.next = doublyLinkedList.tail
		doublyLinkedList.tail.previous = &newNode
	} else {
		//Find the closest way
		distanceFromEnd := doublyLinkedList.lastIndex - index
		if distanceFromEnd < index {
			//Search node from the end
			auxNode := doublyLinkedList.tail

			//Find node at index
			for i := doublyLinkedList.lastIndex; i != index; i-- {
				auxNode = auxNode.previous
			}

			//Add newNode at index
			auxNode.previous.next = &newNode
			newNode.previous = auxNode.previous
			auxNode.previous = &newNode
			newNode.next = auxNode
		} else {
			//Search node from the start
			auxNode := doublyLinkedList.head

			//Find node at index
			for i := 0; i != index; i++ {
				auxNode = auxNode.next
			}

			//Add newNode at index
			auxNode.previous.next = &newNode
			newNode.previous = auxNode.previous
			auxNode.previous = &newNode
			newNode.next = auxNode
		}
	}

	//Increment last position
	doublyLinkedList.lastIndex++

	return nil
}

func (doublyLinkedList *DoublyLinkedList) Remove() error {
	if doublyLinkedList.lastIndex < 0 {
		return fmt.Errorf("Lista vazia.")
	}

	doublyLinkedList.tail.previous.next = nil
	doublyLinkedList.tail = doublyLinkedList.tail.previous

	//Update lastIndex
	doublyLinkedList.lastIndex--

	return nil
}

func (doublyLinkedList *DoublyLinkedList) RemoveOnIndex(index int) error {
	if doublyLinkedList.lastIndex < 0 {
		return fmt.Errorf("Lista vazia.")
	}

	if index < 0 || index > doublyLinkedList.lastIndex {
		return fmt.Errorf("Index inválido.")
	}

	if index == 0 {
		doublyLinkedList.head.next.previous = nil
		doublyLinkedList.head = doublyLinkedList.head.next
	} else if index == doublyLinkedList.lastIndex {
		doublyLinkedList.tail.previous.next = nil
		doublyLinkedList.tail = doublyLinkedList.tail.previous
	} else {
		distanceFromEnd := doublyLinkedList.lastIndex - index

		if distanceFromEnd < index {
			auxNode := doublyLinkedList.tail
			for i := doublyLinkedList.lastIndex; i != index; i-- {
				auxNode = auxNode.previous
			}

			auxNode.previous.next = auxNode.next
			auxNode.next.previous = auxNode.previous

		} else {
			auxNode := doublyLinkedList.head
			for i := 0; i < doublyLinkedList.lastIndex; i++ {
				auxNode = auxNode.next
			}

			auxNode.previous.next = auxNode.next
			auxNode.next.previous = auxNode.previous
		}
	}

	//Update last index
	doublyLinkedList.lastIndex--

	return nil
}

func (doublyLinkedList *DoublyLinkedList) GetNode(index int) (*NodeDoubly, error) {
	if doublyLinkedList.lastIndex < 0 {
		return nil, fmt.Errorf("Lista vazia.")
	}

	if index < 0 || index > doublyLinkedList.lastIndex {
		return nil, fmt.Errorf("Index inválido")
	}

	if index == 0 {
		return doublyLinkedList.head, nil
	} else if index == doublyLinkedList.lastIndex {
		return doublyLinkedList.tail, nil
	} else {
		distanceFromEnd := doublyLinkedList.lastIndex - index

		if distanceFromEnd < index {
			//Search from the tail
			auxNode := doublyLinkedList.tail
			for i := doublyLinkedList.lastIndex; i != index; i-- {
				auxNode = auxNode.previous
			}
			return auxNode, nil
		} else {
			//Search from the head
			auxNode := doublyLinkedList.head
			for i := 0; i != index; i++ {
				auxNode = auxNode.next
			}
			return auxNode, nil
		}
	}

	return nil, nil
}

func (doublyLinkedList *DoublyLinkedList) GetValue(index int) (int, error) {
	if doublyLinkedList.lastIndex < 0 {
		return 0, fmt.Errorf("Lista Vazia")
	}

	if index < 0 || index > doublyLinkedList.lastIndex {
		return 0, fmt.Errorf("Index inválido")
	}

	node, erro := doublyLinkedList.GetNode(index)

	return node.value, erro
}

func (doublyLinkedList *DoublyLinkedList) SetValue(value, index int) error {
	if doublyLinkedList.lastIndex < 0 {
		return fmt.Errorf("Lista vazia.")
	}

	if index < 0 || index > doublyLinkedList.lastIndex {
		return fmt.Errorf("Index inválido.")
	}

	if index == 0 {
		doublyLinkedList.head.value = value
	} else if index == doublyLinkedList.lastIndex {
		doublyLinkedList.tail.value = value
	} else {
		distanceFromEnd := doublyLinkedList.lastIndex - index

		if distanceFromEnd < index {
			//Search from the tail
			auxNode := doublyLinkedList.tail
			for i := doublyLinkedList.lastIndex; i != index; i-- {
				auxNode = auxNode.previous
			}
			auxNode.value = value
		} else {
			//Search from the head
			auxNode := doublyLinkedList.head
			for i := 0; i != index; i++ {
				auxNode = auxNode.next
			}
			auxNode.value = value
		}
	}
	return nil
}

func (doublyLinkedList *DoublyLinkedList) Length() int {
	return doublyLinkedList.lastIndex + 1
}
