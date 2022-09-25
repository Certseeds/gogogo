package package_variable_function

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPackages(t *testing.T) {
	packages()
	fmt.Println("My favorite number is", rand.Intn(10))
}
