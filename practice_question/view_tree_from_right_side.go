package practice_question

import "fmt"

func main() {
	// node0 := Node{name: "0"}
	// node1 := Node{name: "1"}
	node2 := Node{name: "2"}
	node3 := Node{name: "3"}
	node4 := Node{name: "4"}
	node5 := Node{name: "5"}
	node6 := Node{name: "6"}
	node7 := Node{name: "7"}
	node8 := Node{name: "8"}
	node9 := Node{name: "9"}

	// node2.rightChild = &node0
	// node2.leftChild = &node1
	node4.rightChild = &node2
	node4.leftChild = &node3
	node5.rightChild = &node4 // node 5 is the root
	node5.leftChild = &node6
	node6.leftChild = &node7
	node7.leftChild = &node8
	node7.rightChild = &node9

	// Shape designed illustrated below
	//     /\
	//    / /\
	//   /\  /\

	result := searchTree(node5, 0, nil)
	for k, v := range result {
		fmt.Println(k, v.name)
	}
}

type Node struct {
	name       string
	leftChild  *Node
	rightChild *Node
}

func searchTree(node Node, level int, firstNodeFoundAtLevel map[int]Node) map[int]Node {
	if firstNodeFoundAtLevel == nil {
		firstNodeFoundAtLevel = make(map[int]Node)
	}

	_, ok := firstNodeFoundAtLevel[level]
	if !ok {
		firstNodeFoundAtLevel[level] = node
	}

	if node.rightChild != nil {
		firstNodeFoundAtLevel = searchTree(*node.rightChild, level+1, firstNodeFoundAtLevel)
	}

	if node.leftChild != nil {
		firstNodeFoundAtLevel = searchTree(*node.leftChild, level+1, firstNodeFoundAtLevel)
	}

	return firstNodeFoundAtLevel
}

func countNodes(node Node) int {
	if node.leftChild == nil && node.rightChild == nil {
		return 1
	}

	if node.leftChild != nil && node.rightChild == nil {
		return 1 + countNodes(*node.leftChild)
	}

	if node.leftChild != nil && node.rightChild != nil {
		return 1 + countNodes(*node.leftChild) + countNodes(*node.rightChild)
	}

	return 1 + countNodes(*node.rightChild)
}
