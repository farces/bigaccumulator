package bufbig

import (
	"math"
	"math/big"
)

/* Wrapper for big.Int addition, implementing an intermediate accumulator to improve
   performance where an accumulator with possible size > math.MaxInt64 is required. 
   Flushes t_acc buffer when Value() is requested, or when addition could potentially
   exceed max int64 size. 
*/
type BigAccumulator struct {
	t_acc int64 //intermediate acc
	val   *big.Int
}

//add an int to a bigint, buffers additions and flushes when overflow or underflow detected
func (x *BigAccumulator) AddInt(y int) *BigAccumulator {
	if x.val == nil {
		x.val = new(big.Int)
	}

	n := int64(y)
	if y > 0 && (x.t_acc > (math.MaxInt64 - n)) {
		x.flush()
	} else if y < 0 && (x.t_acc < (math.MinInt64 - n)) {
		x.flush()
	}
	x.t_acc = x.t_acc + n
	return x
}

func (x *BigAccumulator) flush() {
	if x.t_acc == 0 {
		return
	}
	x.val.Add(x.val, big.NewInt(x.t_acc))
	x.t_acc = 0
}

//returns underlying big.Int value, after flushing any buffered value
func (x *BigAccumulator) Value() *big.Int {
	if x.t_acc != 0 {
		x.flush()
	}
	return x.val
}
