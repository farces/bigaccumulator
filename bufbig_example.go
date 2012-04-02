package main

import "github.com/farces/dumb/bufbig"
import (
	"fmt"
)

func main() {
	x := new(bufbig.BigAccumulator).Init()
	for y := 0; y < 250000; y++ {
		x.AddInt(1)
	}
	fmt.Println(x.Value())
}
