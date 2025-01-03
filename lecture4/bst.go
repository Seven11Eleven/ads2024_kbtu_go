package lecture4

import "fmt"

type Node struct {
	Item  int
	Left  *Node
	Right *Node
}

func (noda *Node) CreateNode(item int) *Node {
	newNoda := &Node{}
	newNoda.Item = item
	newNoda.Left, newNoda.Right = nil, nil
	return newNoda
}

func (noda *Node) InsertNode(root *Node, data int) *Node {
	if root == nil {
		return noda.CreateNode(data)
	}

	if data < root.Item {
		root.Left = noda.InsertNode(root.Left, data)
	} else if data > root.Item {
		root.Right = noda.InsertNode(root.Right, data)
	}
	return root
}

func (noda *Node) InorderTraversal(root *Node) {
	if root != nil {
		noda.InorderTraversal(root.Left)
		fmt.Println(root.Item)
		noda.InorderTraversal(root.Right)
	}
}

func (noda *Node) SearchNode(root *Node, key int) *Node {
	if root == nil || root.Item == key {
		return root
	}

	if root.Item < key {
		return noda.SearchNode(root.Right, key)
	}

	return noda.SearchNode(root.Left, key)
}

func (noda *Node) MinValueNode(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (noda *Node) DeleteNode(root *Node, data int) *Node {
	if root == nil {
		return root
	}
	if data < root.Item {
		root.Left = noda.DeleteNode(root.Left, data)
	} else if data > root.Item {
		root.Right = noda.DeleteNode(root.Right, data)
	} else {
		if root.Left == nil {
			temp := root.Right
			root = nil
			return temp
		} else if root.Right == nil {
			temp := root.Left
			root = nil
			return temp
		}
		temp := noda.MinValueNode(root.Right)
		root.Item = temp.Item
		root.Right = noda.DeleteNode(root.Right, temp.Item)
	}
	return root
}
