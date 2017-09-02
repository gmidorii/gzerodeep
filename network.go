package main

import (
	"fmt"
	"math"

	"github.com/gonum/matrix/mat64"
)

func network() {
	// 2 x 2 matrix
	matrix := mat64.NewDense(2, 2, nil)
	fmt.Println(matrix)
}

func sigmoid(x mat64.Dense) mat64.Dense {
	return 1 / (1 + math.Exp(-x))
}

func softmax(a float64) {
}
