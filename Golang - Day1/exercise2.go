package main

import "fmt"

type TreeNode struct {
	val string
	left *TreeNode
	right *TreeNode
}

// postorder traversal: using method
func (node TreeNode) preorder(result *string) {
	*result += node.val
	if node.left != nil {
		node.left.preorder(result)
	}
	if node.right != nil {
		node.right.preorder(result)
	}

}

// postorder traversal: passing node as argument
func postorder(node *TreeNode, result *string) {
	if node == nil {
		return
	}
	postorder(node.left, result)
	postorder(node.right, result)
	*result += node.val
}


func main() {
	// building tree
	nodeA := TreeNode{"a", nil, nil}
	nodeB := TreeNode{"b", nil, nil}
	nodeC := TreeNode{"c", nil, nil}
	nodeMinus := TreeNode{"-", &nodeB, &nodeC}
	nodePlus := TreeNode{"+", &nodeA, &nodeMinus}

	var preResult, postResult string
	nodePlus.preorder(&preResult)
	fmt.Println(preResult)

	postorder(&nodePlus, &postResult)
	fmt.Println(postResult)

}
