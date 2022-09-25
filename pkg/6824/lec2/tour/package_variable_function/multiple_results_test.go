package package_variable_function

import (
	"fmt"
	"testing"
)

func Test_multiplyResultsSwap(t *testing.T) {
	a, b := multiplyResultsSwap("hello", "world")
	fmt.Println(a, b)
}
