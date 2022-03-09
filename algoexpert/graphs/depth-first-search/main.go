package main

//Node Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.

type Node struct {
	Name     string
	Children []*Node
}

func main() {

}
func (n *Node) depthFirstSearch(array []string) []string {
	// Write your code here.

	// Pseudocode
	//search(Node root) {
	//	if root == nil {
	//		return
	//	}
	//	visit(root)
	//	root.visited = true
	//	for each Node n in root.children {
	//		if n.visited == false {
	//			search(n)
	//		}
	//	}
	//}
	if n == nil {
		return array
	}
	array = append(array, n.Name)

	for _, value := range n.Children {
		array = value.depthFirstSearch(array)
	}
	return array
}
