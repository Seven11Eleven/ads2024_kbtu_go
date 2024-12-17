package lecture1_test

import (
	"github.com/Seven11Eleven/ads2024_kbtu_go/lecture1"
	"testing"
)

func TestStack_Push_Pop(t *testing.T) {
	stack := lecture1.NewStack()

	// Тестируем Push
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Проверка, что размер стека правильный
	if stack.Size() != 3 {
		t.Errorf("Expected stack size 3, but got %d", stack.Size())
	}

	// Тестируем Pop
	item := stack.Pop()
	if item != 3 {
		t.Errorf("Expected 3, but got %v", item)
	}

	// После одного Pop размер стека должен быть 2
	if stack.Size() != 2 {
		t.Errorf("Expected stack size 2, but got %d", stack.Size())
	}

	// Тестируем, что следующее удаление будет 2
	item = stack.Pop()
	if item != 2 {
		t.Errorf("Expected 2, but got %v", item)
	}

	// Проверяем размер стека после двух удалений
	if stack.Size() != 1 {
		t.Errorf("Expected stack size 1, but got %d", stack.Size())
	}
}

func TestStack_Peek(t *testing.T) {
	stack := lecture1.NewStack()

	// Тестируем Peek при пустом стеке
	peekItem := stack.Peek()
	if peekItem != nil {
		t.Errorf("Expected nil, but got %v", peekItem)
	}

	// Добавляем элементы и проверяем Peek
	stack.Push(1)
	stack.Push(2)

	peekItem = stack.Peek()
	if peekItem != 2 {
		t.Errorf("Expected top item 2, but got %v", peekItem)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := lecture1.NewStack()

	// Проверка, что стек пуст
	if !stack.IsEmpty() {
		t.Errorf("Expected empty stack, but it's not empty")
	}

	// Добавляем элементы в стек
	stack.Push(1)

	// Проверка, что стек не пуст
	if stack.IsEmpty() {
		t.Errorf("Expected non-empty stack, but it's empty")
	}
}

func TestStack_Pop_EmptyStack(t *testing.T) {
	stack := lecture1.NewStack()

	// Проверка Pop на пустом стеке
	item := stack.Pop()
	if item != nil {
		t.Errorf("Expected nil, but got %v", item)
	}
}
