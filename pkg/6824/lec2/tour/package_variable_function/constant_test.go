package package_variable_function

import (
	"fmt"
	"testing"
)

func TestConstant(t *testing.T) {
	constant()
}

func TestSmallOrBig(t *testing.T) {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
