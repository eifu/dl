package dl

import (
	"fmt"
	"testing"
)

func TestClassifier(t *testing.T) {

	n := Node{
		weight: []float64{0.3, 0.5, 0.5},
		bias:   0.1,
	}

	X := []float64{4.0, 3.1, 2.0}

	Y := n.classifier(X)
	fmt.Println(Y)

}

func TestSoftmax(t *testing.T) {

	test := []float64{1.0, 2.0, 3.0}

	result := softmax(test)
	fmt.Println(result)

}
