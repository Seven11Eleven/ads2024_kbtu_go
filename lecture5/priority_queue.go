package lecture5

type Node struct {
	Data     int
	Priority int
	Next     *Node
}

func NewNode(d, p int) *Node {
	temp := &Node{}
	temp.Data = d
	temp.Priority = p
	temp.Next = nil

	return temp
}

func Peek(head **Node) int { return (*head).Data }

func Pop(head **Node, d, p int) {
	//temp := *head
	*head = (*head).Next

}

func Push(head **Node, d, p int) {
	start := *head
	temp := NewNode(d, p)
	if (*head).Priority < p {
		temp.Next = *head
		*head = temp
	} else {
		for start.Next != nil && start.Next.Priority > p {
			start = start.Next
		}

		temp.Next = start.Next
		start.Next = temp
	}
}

func IsEmpty(head **Node) bool { return (*head) == nil }
