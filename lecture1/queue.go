package lecture1

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{items: make([]interface{}, 0)}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (item interface{}) {
	if len(q.items) == 0 {
		return nil
	}
	item = q.items[0]
	q.items = q.items[1:]
	return item
}
func (q *Queue) Front() (item interface{}) {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

func (q *Queue) IsEmpty() (flag bool) {
	return len(q.items) == 0
}

func (q *Queue) Size() (size interface{}) {
	return len(q.items)
}
