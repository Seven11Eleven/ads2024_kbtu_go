package lecture5

import "errors"

type Heap struct {
	items []int
}

func swap(x, y *int) {
	*x, *y = *y, *x
}

func (h *Heap) heapify(i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	size := len(h.items) - 1
	if left < size && h.items[left] > h.items[largest] {
		largest = left
	}
	if right < size && h.items[right] > h.items[largest] {
		largest = right
	}
	if largest != i {
		swap(&h.items[i], &h.items[largest])
		h.heapify(largest)
	}
}

func (h *Heap) Insert(key int) {
	h.items = append(h.items, key)
	i := len(h.items) - 1
	for i != 0 && h.items[(i-1)/2] < h.items[i] {
		swap(&h.items[i], &h.items[(i-1)/2])
		i = (i - 1) / 2
	}
}

func (h *Heap) ExtractMax() (int, error) {
	if len(h.items) <= 0 {
		return -1, errors.New("heap underflow")
	}
	if len(h.items) == 1 {
		root := h.items[0]
		h.items = h.items[:0]
		return root, nil
	}
	root := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.heapify(0)
	return root, nil
}

func (h *Heap) DeleteKey(i int) {
	if i >= len(h.items)-1 {
		return
	}
	h.items[i] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.heapify(i)
}

func (h *Heap) IncreaseKey(i, newVal int) {
	if i >= len(h.items)-1 || h.items[i] >= newVal {
		return
	}

	h.items[i] = newVal
	for i != 0 && h.items[(i-1)/2] < h.items[i] {
		swap(&h.items[i], &h.items[(i-1)/2])
		i = (i - 1) / 2
	}
}
