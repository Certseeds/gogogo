package flow_control

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func forCycle() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forCycleWithStateEmpty() {
	sum := 1
	x := 0
	for sum < 1024 {
		sum += sum
		x++
	}
	fmt.Println(sum)
}

func forCycleWithoutThreePartOrCalledWhile() {
	sum := 1
	for sum < 1024 {
		sum += sum
	}
	fmt.Println(sum)
}

func forEver() (sum int) {
	for {
		sum = 30
		return
	}
}

func sqrt(num float64) string {
	if num < 0 {
		return "-" + sqrt(-num)
	}
	return fmt.Sprint(math.Sqrt(num))
}

func powWithLimit(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		// C++17才有的 if-init语法
		return v
	}
	return lim
}

func powWithLimitButPrint(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println(v)
	}
	return lim
}

func SqrtReal(input float64) float64 {
	for x := float64(1); ; x -= (x*x - input) / (2 * x) {
		if math.Abs(x*x-input) < 0.0000001 {
			return x
		}
	}
}

func switchRunningOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switchWithEvaluationCase() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchWithoutInput() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferPrint() {
	// ps, it is function level, not scope
	defer fmt.Println("world")
	fmt.Println("hello")
}

func deferPrintStack() {
	defer fmt.Println("world")
	for count := 0; count < 10; count++ {
		defer fmt.Println(count)
	}
	fmt.Println("hello")
}
