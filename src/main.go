package main

import (
    "fmt"
    "./perceptron"
    "math/rand"
)

var (
    x1,x2,x3,x4 float64 = 1,0.3,0.2,0.5
    w1,w2,w3,w4 float64 = 1.5,0.2,1.1,1.05
    threshold float64 = 1
    bias float64 = 0.3

    sumInputsWeights float64
    willOpen bool
)

func main(){
    fmt.Println("---Stupid example---")
    fmt.Print("x1,x2,x3,x4 = ")
    fmt.Println(x1,x2,x3,x4)
    fmt.Print("w1,w2,w3,w4 = ")
    fmt.Println(w1,w2,w3,w4)
    fmt.Print("threshold, bias = ")
    fmt.Println(threshold, bias)

    sumInputsWeights = x1 * w1 + x2 * w2 + x3 * w3 + x4 * w4
    fmt.Print("sumInputsWeights = ")
    fmt.Println(sumInputsWeights)

    willOpen = activate(sumInputsWeights + bias) > threshold
    fmt.Print("willOpen = ")
    fmt.Println(willOpen)
    fmt.Println("---End stupid example---")

    w := make([]float64, 5, 5)

    for i := 0; i < 5; i++ {
        w[i] = rand.Float64() * 5
    }

    fmt.Print("w = ")
    fmt.Println(w)

    var p *perceptron.Perceptron = perceptron.ConsPerceptron(5)
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
