package lecture2

import "testing"

func TestInsertFirst(t *testing.T) {
	dll := &DoublyLinkedList{}

	dll.InsertFirst(10)
	if dll.Head == nil || dll.Head.Item != 10 {
		t.Errorf("Expected Head item 10, but got %v", dll.Head.Item)
	}
	if dll.Tail == nil || dll.Tail.Item != 10 {
		t.Errorf("Expected Tail item 10, but got %v", dll.Tail.Item)
	}

	// Вставляем второй элемент
	dll.InsertFirst(20)
	if dll.Head == nil || dll.Head.Item != 20 {
		t.Errorf("Expected Head item 20, but got %v", dll.Head.Item)
	}
	if dll.Tail == nil || dll.Tail.Item != 10 {
		t.Errorf("Expected Tail item 10, but got %v", dll.Tail.Item)
	}
	if dll.Head.Next == nil || dll.Head.Next.Item != 10 {
		t.Errorf("Expected second item to be 10, but got %v", dll.Head.Next.Item)
	}
}

func TestInsertLast(t *testing.T) {
	dll := &DoublyLinkedList{}

	// Вставляем первый элемент
	dll.InsertLast(10)
	if dll.Head == nil || dll.Head.Item != 10 {
		t.Errorf("Expected Head item 10, but got %v", dll.Head.Item)
	}
	if dll.Tail == nil || dll.Tail.Item != 10 {
		t.Errorf("Expected Tail item 10, but got %v", dll.Tail.Item)
	}

	// Вставляем второй элемент
	dll.InsertLast(20)
	if dll.Tail == nil || dll.Tail.Item != 20 {
		t.Errorf("Expected Tail item 20, but got %v", dll.Tail.Item)
	}
	if dll.Head.Next == nil || dll.Head.Next.Item != 20 {
		t.Errorf("Expected second item to be 20, but got %v", dll.Head.Next.Item)
	}
}

func TestDeleteFirst(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.InsertFirst(10)
	dll.InsertFirst(20)

	// Удаляем первый элемент
	dll.DeleteFirst()
	if dll.Head == nil || dll.Head.Item != 10 {
		t.Errorf("Expected Head item 10 after deleting first, but got %v", dll.Head.Item)
	}
	if dll.Tail == nil || dll.Tail.Item != 10 {
		t.Errorf("Expected Tail item 10, but got %v", dll.Tail.Item)
	}

	// Удаляем последний элемент
	dll.DeleteFirst()
	if dll.Head != nil {
		t.Errorf("Expected empty list, but Head is %v", dll.Head)
	}
	if dll.Tail != nil {
		t.Errorf("Expected empty list, but Tail is %v", dll.Tail)
	}
}

func TestDeleteLast(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.InsertFirst(20)
	dll.InsertFirst(10)

	dll.DeleteLast()
	if dll.Tail == nil || dll.Tail.Item != 10 {
		t.Errorf("Expected Tail item 10, but got %v", dll.Tail.Item)
	}

	dll.DeleteLast()
	if dll.Head != nil {
		t.Errorf("Expected empty list, but Head is %v", dll.Head)
	}
	if dll.Tail != nil {
		t.Errorf("Expected empty list, but Tail is %v", dll.Tail)
	}
}

func TestDelete(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.InsertFirst(10)
	dll.InsertFirst(20)
	dll.InsertFirst(30)

	// Удаляем элемент в середине
	dll.Delete(20)
	if dll.Head.Next.Item != 10 {
		t.Errorf("Expected second item to be 10, but got %v", dll.Head.Next.Item)
	}

	// Удаляем последний элемент
	dll.Delete(10)
	if dll.Tail.Item != 30 {
		t.Errorf("Expected Tail item 30, but got %v", dll.Tail.Item)
	}

	// Удаляем последний оставшийся элемент
	dll.Delete(30)
	if dll.Head != nil || dll.Tail != nil {
		t.Errorf("Expected empty list, but Head: %v, Tail: %v", dll.Head, dll.Tail)
	}
}

func TestSearch(t *testing.T) {
	dll := &DoublyLinkedList{}
	dll.InsertFirst(10)
	dll.InsertFirst(20)
	dll.InsertFirst(30)

	// Поиск существующего элемента
	if !dll.Search(20) {
		t.Errorf("Expected to find 20 in the list")
	}

	// Поиск несуществующего элемента
	if dll.Search(100) {
		t.Errorf("Expected not to find 100 in the list")
	}

	// Поиск первого элемента
	if !dll.Search(30) {
		t.Errorf("Expected to find 30 in the list")
	}

	// Поиск последнего элемента
	if !dll.Search(10) {
		t.Errorf("Expected to find 10 in the list")
	}
}
