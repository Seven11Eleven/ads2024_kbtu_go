package lecture2_test

import (
	"github.com/Seven11Eleven/ads2024_kbtu_go/lecture2"
	"testing"
)

func TestLinkedList_Insert(t *testing.T) {
	ll := &lecture2.LinkedList{}

	// Вставляем элементы
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)

	// Проверка, что голова списка указывает на последний вставленный элемент
	if ll.Head.Item != 30 {
		t.Errorf("Expected Head Item 30, but got %v", ll.Head.Item)
	}

	// Проверка последовательности элементов
	if ll.Head.Next.Item != 20 {
		t.Errorf("Expected second Item 20, but got %v", ll.Head.Next.Item)
	}

	if ll.Head.Next.Next.Item != 10 {
		t.Errorf("Expected third Item 10, but got %v", ll.Head.Next.Next.Item)
	}
}

func TestLinkedList_DeletionFirst(t *testing.T) {
	ll := &lecture2.LinkedList{}

	// Вставляем элементы
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)

	// Удаляем первый элемент
	Item := ll.DeletionFirst()
	if Item != 30 {
		t.Errorf("Expected first deleted Item 30, but got %v", Item)
	}

	// Проверяем, что теперь голова списка указывает на следующий элемент
	if ll.Head.Item != 20 {
		t.Errorf("Expected Head Item 20, but got %v", ll.Head.Item)
	}

	// Удаляем еще один элемент
	Item = ll.DeletionFirst()
	if Item != 20 {
		t.Errorf("Expected second deleted Item 20, but got %v", Item)
	}

	// Проверяем, что голова списка указывает на последний элемент
	if ll.Head.Item != 10 {
		t.Errorf("Expected Head Item 10, but got %v", ll.Head.Item)
	}
}

func TestLinkedList_DeletionLast(t *testing.T) {
	ll := &lecture2.LinkedList{}

	// Вставляем элементы
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)

	// Удаляем последний элемент
	Item := ll.DeletionLast()
	if Item != 10 {
		t.Errorf("Expected last deleted Item 10, but got %v", Item)
	}

	// Проверяем, что теперь последний элемент в списке
	if ll.Head.Item != 30 {
		t.Errorf("Expected Head Item 30, but got %v", ll.Head.Item)
	}

	// Удаляем еще один элемент
	Item = ll.DeletionLast()
	if Item != 20 {
		t.Errorf("Expected second last deleted Item 20, but got %v", Item)
	}
	Item = ll.DeletionLast()

	// Проверяем, что список пуст
	if ll.Head != nil {
		t.Errorf("Expected empty list, but Head is %v", ll.Head)
	}
}

func TestLinkedList_Search(t *testing.T) {
	ll := &lecture2.LinkedList{}

	// Вставляем элементы
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)

	// Проверка наличия элементов
	if !ll.Search(20) {
		t.Errorf("Expected to find Item 20, but did not")
	}

	if ll.Search(40) {
		t.Errorf("Expected not to find Item 40, but it was found")
	}
}
