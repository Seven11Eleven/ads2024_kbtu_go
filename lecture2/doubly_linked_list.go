package lecture2

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (dll *DoublyLinkedList) InsertFirst(item interface{}) {

	node := &Node{Item: item}
	if dll.Head == nil {
		dll.Head = node
		dll.Tail = node
	} else {
		node.Next = dll.Head
		dll.Head.Prev = node
		dll.Head = node
	}
}

func (dll *DoublyLinkedList) InsertLast(item interface{}) {
	node := &Node{Item: item}
	if dll.Tail == nil {
		dll.Head = node
		dll.Tail = node
	} else {
		node.Prev = dll.Tail
		dll.Tail.Next = node
		dll.Tail = node
	}
}

func (dll *DoublyLinkedList) DeleteFirst() {
	if dll.Head == nil {
		return
	}
	if dll.Head.Next == nil {
		dll.Head = nil
		dll.Tail = nil
		return
	}
	dll.Head = dll.Head.Next
	dll.Head.Prev = nil
}

func (dll *DoublyLinkedList) DeleteLast() {
	if dll.Tail == nil {
		return
	}
	if dll.Tail.Prev == nil {
		dll.Head = nil
		dll.Tail = nil
		return
	}
	dll.Tail = dll.Tail.Prev
	dll.Tail.Next = nil
}

func (dll *DoublyLinkedList) Delete(item interface{}) {
	if dll.Head == nil {
		return
	}
	curr := dll.Head
	for curr != nil {
		if curr.Item == item {
			if curr.Prev == nil {
				dll.Head = curr.Next
				if dll.Head != nil {
					dll.Head.Prev = nil
				}
			} else {
				curr.Prev.Next = curr.Next
				if curr.Next != nil {
					curr.Next.Prev = curr.Prev
				}
			}

			if curr.Next == nil {
				dll.Tail = curr.Prev
			}

			return
		}
		curr = curr.Next
	}
}

func (dll *DoublyLinkedList) Search(item interface{}) bool {
	if dll.Head == nil {
		return false
	}
	left := dll.Head
	right := dll.Tail
	for left != right && left != nil && right != nil {
		if left.Item == item || right.Item == item {
			return true
		}
		left = left.Next
		right = right.Prev
	}
	if left == right && left.Item == item {
		return true
	}
	return false
}
