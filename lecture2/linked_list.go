package lecture2

type Node struct {
	Item interface{}
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Insert(item interface{}) {
	node := &Node{Item: item}
	node.Next = ll.Head
	ll.Head = node
}

func (ll *LinkedList) DeletionFirst() interface{} {
	if ll.Head == nil {
		return nil
	}
	item := ll.Head.Item
	ll.Head = ll.Head.Next
	return item
}

func (ll *LinkedList) DeletionLast() interface{} {
	if ll.Head == nil {
		return nil
	}

	if ll.Head.Next == nil {
		item := ll.Head.Item
		ll.Head = nil
		return item
	}

	curr := ll.Head
	for curr.Next != nil && curr.Next.Next != nil {
		curr = curr.Next
	}

	item := curr.Next.Item
	curr.Next = nil

	return item
}

func (ll *LinkedList) Search(item interface{}) bool {
	if ll.Head == nil {
		return false
	}
	curr := ll.Head
	for curr != nil {
		if curr.Item == item {
			return true
		}
		curr = curr.Next
	}
	return false
}
