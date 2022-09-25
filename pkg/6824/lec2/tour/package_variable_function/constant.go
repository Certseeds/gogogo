package package_variable_function

import "fmt"

const (
	MathPi = 3.141592653589793238462643383279
	Big    = 1 << 100
	Small  = Big >> 99
)

func constant() {
	const World = "nihao"

	fmt.Println("Hello", World)
	fmt.Println("Happy", MathPi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
