package main

// seems only need to use make package to `main`, it will be a application start

import (
	"fmt"
	"gogogo/greetings"
	"log"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Printf("Hello,%s!\n", "world")
	fmt.Println(quote.GoV3()) // update it to v3 is a stupid idea
	world, err := greetings.Hello("World")
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err.Error()) // log.Fatal seems exit(-1)
	} else {
		fmt.Println(world)
	}
}
