package main

import (
	"fmt"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Printf("Hello,%s!\n", "world")
	fmt.Printf(quote.GoV3()) // update it to v3 is a stupid idea
}
