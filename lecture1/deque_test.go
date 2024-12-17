package lecture1_test

import (
	"github.com/Seven11Eleven/ads2024_kbtu_go/lecture1"
	"testing"
)

func TestDeque_PushFront(t *testing.T) {
	var d lecture1.Deque

	// Тестируем DequeLinked
	d = lecture1.NewDeque()
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)

	if d.Front() != 3 {
		t.Errorf("Expected front to be 3, but got %v", d.Front())
	}

	if d.Back() != 1 {
		t.Errorf("Expected back to be 1, but got %v", d.Back())
	}

	// Тестируем DequeSlice
	d = lecture1.NewDequeSlice()
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)

	if d.Front() != 3 {
		t.Errorf("Expected front to be 3, but got %v", d.Front())
	}

	if d.Back() != 1 {
		t.Errorf("Expected back to be 1, but got %v", d.Back())
	}
}

func TestDeque_PushBack(t *testing.T) {
	var d lecture1.Deque

	// Тестируем DequeLinked
	d = lecture1.NewDeque() // Или можно использовать: d := &lecture1.DequeLinked{}
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	if d.Front() != 1 {
		t.Errorf("Expected front to be 1, but got %v", d.Front())
	}

	if d.Back() != 3 {
		t.Errorf("Expected back to be 3, but got %v", d.Back())
	}

	// Тестируем DequeSlice
	d = lecture1.NewDequeSlice()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	if d.Front() != 1 {
		t.Errorf("Expected front to be 1, but got %v", d.Front())
	}

	if d.Back() != 3 {
		t.Errorf("Expected back to be 3, but got %v", d.Back())
	}
}

func TestDeque_PopFront(t *testing.T) {
	var d lecture1.Deque

	// Тестируем DequeLinked
	d = lecture1.NewDeque()
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)

	item := d.PopFront()
	if item != 3 {
		t.Errorf("Expected popped item to be 3, but got %v", item)
	}

	if d.Front() != 2 {
		t.Errorf("Expected front to be 2, but got %v", d.Front())
	}

	// Тестируем DequeSlice
	d = lecture1.NewDequeSlice()
	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)

	item = d.PopFront()
	if item != 3 {
		t.Errorf("Expected popped item to be 3, but got %v", item)
	}

	if d.Front() != 2 {
		t.Errorf("Expected front to be 2, but got %v", d.Front())
	}
}

func TestDeque_PopBack(t *testing.T) {
	var d lecture1.Deque

	// Тестируем DequeLinked
	d = lecture1.NewDeque()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	item := d.PopBack()
	if item != 3 {
		t.Errorf("Expected popped item to be 3, but got %v", item)
	}

	if d.Back() != 2 {
		t.Errorf("Expected back to be 2, but got %v", d.Back())
	}

	// Тестируем DequeSlice
	d = lecture1.NewDequeSlice()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	item = d.PopBack()
	if item != 3 {
		t.Errorf("Expected popped item to be 3, but got %v", item)
	}

	if d.Back() != 2 {
		t.Errorf("Expected back to be 2, but got %v", d.Back())
	}
}

func TestDeque_EmptyDeque(t *testing.T) {
	var d lecture1.Deque

	// Тестируем DequeLinked
	d = lecture1.NewDeque()

	if d.Front() != nil {
		t.Errorf("Expected front to be nil, but got %v", d.Front())
	}

	if d.Back() != nil {
		t.Errorf("Expected back to be nil, but got %v", d.Back())
	}

	if item := d.PopFront(); item != nil {
		t.Errorf("Expected nil when popping from front, but got %v", item)
	}

	if item := d.PopBack(); item != nil {
		t.Errorf("Expected nil when popping from back, but got %v", item)
	}

	// Тестируем DequeSlice
	d = lecture1.NewDequeSlice()

	if d.Front() != nil {
		t.Errorf("Expected front to be nil, but got %v", d.Front())
	}

	if d.Back() != nil {
		t.Errorf("Expected back to be nil, but got %v", d.Back())
	}

	if item := d.PopFront(); item != nil {
		t.Errorf("Expected nil when popping from front, but got %v", item)
	}

	if item := d.PopBack(); item != nil {
		t.Errorf("Expected nil when popping from back, but got %v", item)
	}
}
