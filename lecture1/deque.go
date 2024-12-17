package lecture1

type Deque interface {
	PushFront(item interface{})
	PushBack(item interface{})
	PopFront() interface{}
	PopBack() interface{}
	Front() interface{}
	Back() interface{}
	//IsEmpty() bool
	//Size() int
}

type Node struct {
	item interface{}
	next *Node
	prev *Node
}

type DequeLinked struct {
	front *Node
	back  *Node
	size  int
}

func NewDeque() *DequeLinked {
	return &DequeLinked{}
}

func (d *DequeLinked) PushFront(item interface{}) {
	node := &Node{item: item}
	if d.front == nil {
		d.front = node
		d.back = node
	} else {
		node.next = d.front
		d.front.prev = node
		d.front = node
	}
	d.size++
}

func (d *DequeLinked) PushBack(item interface{}) {
	node := &Node{item: item}
	if d.back == nil {
		d.front = node
		d.back = node
	} else {
		node.prev = d.back
		d.back.next = node
		d.back = node
	}
	d.size++
}

func (d *DequeLinked) PopFront() (item interface{}) {
	if d.front == nil {
		return nil
	}
	item = d.front.item
	d.front = d.front.next
	if d.front != nil {
		d.front.prev = nil
	} else {
		d.back = nil
	}
	d.size--
	return item
}

func (d *DequeLinked) PopBack() (item interface{}) {
	if d.back == nil {
		return nil
	}
	item = d.back.item
	d.back = d.back.prev
	if d.back != nil {
		d.back.next = nil
	} else {
		d.front = nil
	}
	d.size--
	return item
}

func (d *DequeLinked) Front() (item interface{}) {
	if d.front == nil {
		return nil
	}
	return d.front.item
}

func (d *DequeLinked) Back() (item interface{}) {
	if d.back == nil {
		return nil
	}
	return d.back.item
}
