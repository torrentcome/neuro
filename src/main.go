package main

import (
	"fmt"
	"./perceptron"
	"math/rand"
)

var (
	x1, x2, x3, x4   float64 = 1.0, 0.3, 0.2, 0.5
	w1, w2, w3, w4   float64 = 1.5, 0.2, 1.1, 1.05
	threshold        float64 = 1.0
	bias             float64 = 0.3
	sumInputsWeights float64
	willOpen         bool
)

func main() {
	fmt.Println("---Stupid example---")
	sumInputsWeights = x1*w1 + x2*w2 + x3*w3 + x4*w4
	fmt.Print("sumInputsWeights = ")
	fmt.Println(sumInputsWeights)

	willOpen = activate(sumInputsWeights+bias) > threshold
	fmt.Print("willOpen = ")
	fmt.Println(willOpen)
	fmt.Println("---End stupid example---")

	const nb = 5
	w := make([]float64, nb, nb)

	for i := 0; i < nb; i++ {
		w[i] = rand.Float64() * nb
	}

	fmt.Print("w = ")
	fmt.Println(w)

	var p *perceptron.Perceptron = perceptron.ConsPerceptron(nb)
	fmt.Println(p)
	fmt.Println(p.Process(w))

}

func activate(sum float64) float64 {
	if sum > 0 {
		return 1
	} else {
		return -1
	}
}
