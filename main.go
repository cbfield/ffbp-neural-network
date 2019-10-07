package main

import (
	"fmt";
	"math/rand";
)

type node struct {
	weights []float64
}

func (n node) evaluate(inputs []float64) {
	for _,weight := range n.weights{
		fmt.Println(weight)
	}
}

func gimmeSomeInputs(length int) []float64{
	thing := make([]float64,length)
	for i := range thing{
		r := rand.Float64()
		thing[i] = r
	}
	return thing
}

func main() {
	inputLayerSize := 8
	// hiddenLayerSize := 8
	// outputLayerSize := 2
	// n := node{[]float64{1.4,2.4,7.5}}


	inputLayer := gimmeSomeInputs(inputLayerSize)

	fmt.Println(inputLayer)
}