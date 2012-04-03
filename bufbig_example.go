package main

import "github.com/farces/dumb/bufbig"
import (
	"fmt"
)

func main() {
	//x := new(bufbig.BigAccumulator)
	x := bufbig.NewBigAccumulator()
	for y := 0; y < 1000000; y++ {
		x.AddInt(1)
	}
	fmt.Println(x.Value())
}
