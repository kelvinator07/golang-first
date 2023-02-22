package main

import (
	"fmt"

	"rsc.io/quote"
	
	"geeky.com/intro-demo/about"
)

func main() {
    fmt.Println("Hello, World!")
	fmt.Println(about.Name())
	fmt.Println(quote.Go())
}
