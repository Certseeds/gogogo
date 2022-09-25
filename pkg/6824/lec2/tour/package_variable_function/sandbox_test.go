package package_variable_function

import (
	"fmt"
	"testing"
	"time"
)

func TestSandbox(t *testing.T) {
	sandbox()
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
}
