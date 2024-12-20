package lecture1

type Stack struct {
	items []interface{}
}

func NewStack() *Stack {
	return &Stack{items: make([]interface{}, 0)}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (item interface{}) {
	if len(s.items) == 0 {
		return nil
	}
	item = s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]
	return item
}

func (s *Stack) Peek() (item interface{}) {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() (flag bool) {
	return len(s.items) == 0
}

func (s *Stack) Size() (size int) {
	return len(s.items)
}
