package lecture4

import (
	"testing"
)

// Helper для сбора значений из in-order обхода
func collectInOrder(root *Node) []int {
	var result []int
	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node != nil {
			inorder(node.Left)
			result = append(result, node.Item)
			inorder(node.Right)
		}
	}
	inorder(root)
	return result
}

func TestInsertNode(t *testing.T) {
	root := &Node{}
	root = root.InsertNode(nil, 50)
	root.InsertNode(root, 30)
	root.InsertNode(root, 20)
	root.InsertNode(root, 40)
	root.InsertNode(root, 70)
	root.InsertNode(root, 60)
	root.InsertNode(root, 80)

	expected := []int{20, 30, 40, 50, 60, 70, 80} // in-order traversal
	result := collectInOrder(root)

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d, but got %d", v, result[i])
		}
	}
}

func TestSearchNode(t *testing.T) {
	root := &Node{}
	root = root.InsertNode(nil, 50)
	root.InsertNode(root, 30)
	root.InsertNode(root, 70)

	// Позитивный тест: поиск существующего узла
	node := root.SearchNode(root, 70)
	if node == nil || node.Item != 70 {
		t.Errorf("Expected to find node with value 70, but got nil or incorrect value")
	}

	// Негативный тест: поиск несуществующего узла
	node = root.SearchNode(root, 100)
	if node != nil {
		t.Errorf("Expected to not find node with value 100, but got %v", node.Item)
	}
}

func TestDeleteNode(t *testing.T) {
	root := &Node{}
	root = root.InsertNode(nil, 50)
	root.InsertNode(root, 30)
	root.InsertNode(root, 20)
	root.InsertNode(root, 40)
	root.InsertNode(root, 70)
	root.InsertNode(root, 60)
	root.InsertNode(root, 80)

	// Удаляем лист
	root.DeleteNode(root, 20)
	expectedAfterDelete20 := []int{30, 40, 50, 60, 70, 80}
	if result := collectInOrder(root); !equal(result, expectedAfterDelete20) {
		t.Errorf("Expected %v after deleting 20, but got %v", expectedAfterDelete20, result)
	}

	// Удаляем узел с одним ребенком
	root.DeleteNode(root, 30)
	expectedAfterDelete30 := []int{40, 50, 60, 70, 80}
	if result := collectInOrder(root); !equal(result, expectedAfterDelete30) {
		t.Errorf("Expected %v after deleting 30, but got %v", expectedAfterDelete30, result)
	}

	// Удаляем узел с двумя детьми
	root.DeleteNode(root, 50)
	expectedAfterDelete50 := []int{40, 60, 70, 80}
	if result := collectInOrder(root); !equal(result, expectedAfterDelete50) {
		t.Errorf("Expected %v after deleting 50, but got %v", expectedAfterDelete50, result)
	}
}

// Helper для сравнения двух срезов
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
