package main

import "github.com/farces/dumb/bufbig"
import (
	"fmt"
	"math/big"
)

func main() {
	x := bufbig.BigAccumulator.Init() //init BigAccumulator (Val init required)
	for y := 0; y < 250000; y++ {
		x.AddInt(1)
	}
	fmt.Println(x.Value())
}
