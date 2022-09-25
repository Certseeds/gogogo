package package_variable_function

import (
	"fmt"
	"testing"
)

func TestVariableWithInit(t *testing.T) {
	variableWithInit()
	var c, python, java = true, false, "no!"
	fmt.Println(a, b, c, python, java)
}
