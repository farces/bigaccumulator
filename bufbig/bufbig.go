package bufbig

import (
	"math/big"
)

/* basic big.Int wrapper with an intermediary int64 accumulator */
type BufferedInt struct {
    t int64
    dirty bool
	Val *big.Int
}

//add an int to a bigint, buffers additions and flushes when overflow detected
func (z *BufferedInt) AddInt(y int) *BufferedInt {
    if (z.t >= (9223372036854775807 - int64(y)))  { z.flush() }
    z.t = z.t+int64(y)
    z.dirty = true
    return z
}

//flush value from z.t to val (big.Int)
func (z *BufferedInt) flush() {
    if !z.dirty { return }
    z.Val.Add(z.Val, big.NewInt(z.t))
    z.t = 0
    z.dirty = false
}

//returns underlying big.Int value, after flushing any buffered value
func (z *BufferedInt) Value() *big.Int {
    if z.dirty { z.flush() }
    return z.Val
}
//

