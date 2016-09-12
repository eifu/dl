package dl

import (
	"fmt"
	"testing"
)

func TestSoftmax(t *testing.T) {

	test := []float64{1.0, 2.0, 3.0}

	result := softmax(test)
	fmt.Println(result)	

}
