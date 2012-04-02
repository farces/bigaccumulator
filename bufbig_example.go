package main

import "github.com/farces/dumb/bufbig"
import (
    "fmt"
    "math/big"
)


func main() {
    x := bufbig.BufferedInt{Val:new(big.Int)}
    for y := 0; y<250000; y++ {
        x.AddInt(1)
    }
    fmt.Println(x.Value())
}
