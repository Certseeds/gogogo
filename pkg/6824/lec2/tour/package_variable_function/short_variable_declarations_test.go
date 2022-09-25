package package_variable_function

import (
	"fmt"
	"testing"
)

func TestShortVariableDeclarations(t *testing.T) {
	shortVariableDeclarations()
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
