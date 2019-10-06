package main

import (
	"fmt"
)

type node struct {
	m float64
	b float64
}

func main() {
	fmt.Println("Hello World!")

	var x float64 = 1.5
	var node1 = node{m:2,b:0}
	var y float64
	
	y = node1.m*x + node1.b

	fmt.Println(y)

}