package dl

import (
	"math"
)

func softmax(X []float64) []float64 {

	Y := make([]float64, len(X))
	sum := 0.0

	// get the exponential
	for i, x := range X {
		Y[i] = math.Exp(x)
		sum += Y[i]
	}

	// normalize the score
	for i, _ := range Y {
		Y[i] = Y[i] / sum
	}

	return Y

}
