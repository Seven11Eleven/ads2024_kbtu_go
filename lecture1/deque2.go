package lecture1

type DequeSlice struct {
	items []interface{}
}

func NewDequeSlice() *DequeSlice {
	return &DequeSlice{items: make([]interface{}, 0)}
}

func (d *DequeSlice) PushFront(item interface{}) {
	d.items = append([]interface{}{item}, d.items...)
}

func (d *DequeSlice) PushBack(item interface{}) {
	d.items = append(d.items, item)
}
func (d *DequeSlice) PopFront() interface{} {
	if len(d.items) == 0 {
		return nil
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item
}

func (d *DequeSlice) PopBack() interface{} {
	if len(d.items) == 0 {
		return nil
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item
}
func (d *DequeSlice) Front() interface{} {
	if len(d.items) == 0 {
		return nil
	}
	return d.items[0]
}
func (d *DequeSlice) Back() interface{} {
	if len(d.items) == 0 {
		return nil
	}
	return d.items[len(d.items)-1]
}
