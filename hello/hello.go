package main

// seems only need to use make package to `main`, it will be a application start

import (
	"fmt"
	"gogogo/greetings"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Printf("Hello,%s!\n", "world")
	fmt.Println(quote.GoV3()) // update it to v3 is a stupid idea
	world := greetings.Hello("world")
	fmt.Printf(world)
}
