package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	loop2(5)
}

func loop2(amountOfLoops int) {
	for i := 0; i < amountOfLoops; i++ {
		fmt.Println(quote.Go(), &i)
	}
}
