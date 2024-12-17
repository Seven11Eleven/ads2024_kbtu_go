package lecture1_test

import (
	"github.com/Seven11Eleven/ads2024_kbtu_go/lecture1"
	"testing"
)

func TestQueue_Enqueue_Dequeue(t *testing.T) {
	q := lecture1.NewQueue()

	// Тестируем Enqueue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Проверка, что размер очереди правильный
	if q.Size() != 3 {
		t.Errorf("Expected queue size 3, but got %d", q.Size())
	}

	// Тестируем Dequeue
	item := q.Dequeue()
	if item != 1 {
		t.Errorf("Expected 1, but got %v", item)
	}

	// После одного Dequeue размер очереди должен быть 2
	if q.Size() != 2 {
		t.Errorf("Expected queue size 2, but got %d", q.Size())
	}

	// Тестируем, что следующее удаление будет 2
	item = q.Dequeue()
	if item != 2 {
		t.Errorf("Expected 2, but got %v", item)
	}

	// Проверяем размер очереди после двух удалений
	if q.Size() != 1 {
		t.Errorf("Expected queue size 1, but got %d", q.Size())
	}
}

func TestQueue_Front(t *testing.T) {
	q := lecture1.NewQueue()

	// Тестируем Front при пустой очереди
	frontItem := q.Front()
	if frontItem != nil {
		t.Errorf("Expected nil, but got %v", frontItem)
	}

	// Добавляем элементы и проверяем Front
	q.Enqueue(1)
	q.Enqueue(2)

	frontItem = q.Front()
	if frontItem != 1 {
		t.Errorf("Expected front item 1, but got %v", frontItem)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	q := lecture1.NewQueue()

	// Проверка, что очередь пуста
	if !q.IsEmpty() {
		t.Errorf("Expected empty queue, but it's not empty")
	}

	// Добавляем элементы в очередь
	q.Enqueue(1)

	// Проверка, что очередь не пуста
	if q.IsEmpty() {
		t.Errorf("Expected non-empty queue, but it's empty")
	}
}

func TestQueue_Dequeue_EmptyQueue(t *testing.T) {
	q := lecture1.NewQueue()

	// Проверка Dequeue на пустой очереди
	item := q.Dequeue()
	if item != nil {
		t.Errorf("Expected nil, but got %v", item)
	}
}
