package main

import "math"

func network() {

}

func sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func softmax(a float64) {
}
