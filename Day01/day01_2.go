// Second Question
package main

import (
	"fmt"
)

type Node struct {
	s  string
	left *Node
	right *Node
}

func (n Node) preorder(res *string) {
	*res += n.s
	if n.left != nil {
		n.left.preorder(res)
	}
	if n.right != nil {
		n.right.preorder(res)
	}

}

func (n Node) postorder(res *string) {
	if n.left != nil {
		n.left.postorder(res)
	}
	if n.right != nil {
		n.right.postorder(res)
	}
	*res += n.s
}

func main() {

	nodea := Node{"a", nil, nil}
	nodeb := Node{"b", nil, nil}
	nodec := Node{"c", nil, nil}
	nodem := Node{"-", &nodeb, &nodec}
	nodep := Node{"+", &nodea, &nodem}

	var preres, postres string

	nodep.preorder(&preres)
	fmt.Println(preres)

	nodep.postorder(&postres)
	fmt.Println(postres)

}