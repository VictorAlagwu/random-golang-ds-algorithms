package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func main() {

}


func (tree *BST) Insert(value int) *BST {
	// Write your code here.
	currentNode := tree
	for true {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = &BST{Value: value}
				break
			} else {
				currentNode = currentNode.Left
			}

		} else {
			if currentNode.Right == nil {
				currentNode.Right = &BST{Value: value}
				break
			} else {
				currentNode = currentNode.Right
			}

		}
	}
	// Do not edit the return statement of this method.
	return tree
}

func (tree *BST) Contains(value int) bool {
	// Write your code here.
	currentNode := tree
	for currentNode != nil {
		if value < currentNode.Value {
			currentNode = currentNode.Left
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			return true
		}
	}

	return false
}

func (tree *BST) Remove(value int) *BST {
	// Write your code here.
	tree.remove(value, nil)
	// Do not edit the return statement of this method.
	return tree
}

func (tree *BST) remove(value int, parentNode *BST) {
	currentNode := tree
	for currentNode != nil {
		if value < currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Left
		}else if  value > currentNode.Value {
			parentNode = currentNode
			currentNode = currentNode.Right
		}else {
			if currentNode.Left != nil && currentNode.Right != nil {
				currentNode.Value = currentNode.Right.getMinValue()
				currentNode.Right.remove(currentNode.Value, currentNode)
			}else if parentNode == nil {
				if currentNode.Left != nil {
					currentNode.Value = currentNode.Left.Value
					currentNode.Right = currentNode.Left.Right
					currentNode.Left = currentNode.Left.Left
				}else if currentNode.Right != nil {
					currentNode.Value = currentNode.Right.Value
					currentNode.Left = currentNode.Right.Left
					currentNode.Right = currentNode.Right.Right
				} else {

				}
			}else if parentNode.Left == currentNode {
				if currentNode.Left != nil {
					parentNode.Left = currentNode.Left
				}else {
					parentNode.Left = currentNode.Right
				}
			}else if parentNode.Right == currentNode {
				if currentNode.Left != nil {
					parentNode.Right = currentNode.Left
				}else {
					parentNode.Right = currentNode.Right
				}
			}
			break
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}

	return tree.Left.getMinValue()
}