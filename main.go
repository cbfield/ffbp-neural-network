package main

import (
	"fmt"
)

type node struct {
	m float64
	b float64
}

type layer struct {
	nodes []node
}

func main() {
	node1 := node{m:2,b:0}
	node2 := node{m:3,b:0}
	node3 := node{m:6,b:0}

	nodes := make([]node, 0)
	nodes = append(nodes,node1)
	nodes = append(nodes,node2)
	nodes = append(nodes,node3)
	
	layer1 := layer{nodes:nodes}

	fmt.Println(layer1.nodes)

}