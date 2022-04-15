package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// Find Closest Value in BST (Binary Search Tree)
func main() {
	tree := BST {
				//"nodes": [
				//	{"id": "10", "left": "5", "right": "15", "value": 10},
				//	{"id": "15", "left": "13", "right": "22", "value": 15},
				//	{"id": "22", "left": null, "right": null, "value": 22},
				//	{"id": "13", "left": null, "right": "14", "value": 13},
				//	{"id": "14", "left": null, "right": null, "value": 14},
				//	{"id": "5", "left": "2", "right": "5-2", "value": 5},
				//	{"id": "5-2", "left": null, "right": null, "value": 5},
				//	{"id": "2", "left": "1", "right": null, "value": 2},
				//	{"id": "1", "left": null, "right": null, "value": 1}
				//],
			//"root": "10"
			}
	target := 12
	fmt.Println("Closest value in bst:", tree.FindClosestValue(target))
}

func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	currentNode := tree
	closestValue := 0
	positiveTarget := target + 1
	for currentNode != nil {
		if target < currentNode.Value {
			if positiveTarget == currentNode.Value {
				closestValue = positiveTarget
				break
			}else {
				currentNode = currentNode.Left
			}
		}else if target > currentNode.Value {
			if positiveTarget == currentNode.Value {
				closestValue = positiveTarget
				break
			}else {
				currentNode = currentNode.Right
			}
		} else {
			if positiveTarget == currentNode.Value {
				closestValue = positiveTarget
				break
			} else {
				closestValue = -1
				break
			}
		}
	}
	return closestValue
}

