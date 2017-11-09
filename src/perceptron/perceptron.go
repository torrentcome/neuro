package perceptron

import (
	"fmt"
	"math/rand"
)

type Perceptron struct {
	Weights []float64
	Bias    float64
}

func ConsPerceptron(n int) *Perceptron {
	weights := make([]float64, n, n)
	bias := rand.Float64()*2 - 1

	for i := 0; i < n; i++ {
		weights[i] = rand.Float64()*2 - 1
	}

	return &Perceptron{
		Weights: weights,
		Bias:    bias,
	}
}

func (p *Perceptron) Process(inputs []float64) float64 {
	fmt.Println(inputs)
	sum := p.Bias
	for i, input := range inputs {
		fmt.Println(p.Weights[i])
		sum += float64(input) * p.Weights[i]
	}
	return active(sum)
}

func active(sum float64) float64 {
	if sum > 0 {
		return 1
	} else {
		return -1
	}
}
