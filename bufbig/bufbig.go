package bufbig

import (
	"math"
	"math/big"
)

/* basic big.Int wrapper with an intermediary int64 accumulator */
type BigAccumulator struct {
	t     int64
	dirty bool
	Val   *big.Int
}

//add an int to a bigint, buffers additions and flushes when overflow detected
func (z *BigAccumulator) AddInt(y int) *BigAccumulator {
	if y > 0 && (z.t > (math.MaxInt64 - int64(y))) {
		z.flush()
	} else if y < 0 && (z.t < (math.MinInt64 - int64(y))) {
		z.flush()
	}
	z.t = z.t + int64(y)
	z.dirty = true
	return z
}

func (z *BigAccumulator) flush() {
	if !z.dirty {
		return
	}
	z.Val.Add(z.Val, big.NewInt(z.t))
	z.t = 0
	z.dirty = false
}

//returns underlying big.Int value, after flushing any buffered value
func (z *BigAccumulator) Value() *big.Int {
	if z.dirty {
		z.flush()
	}
	return z.Val
}

//
