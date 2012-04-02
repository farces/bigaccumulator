package bufbig

import (
	"math"
	"math/big"
)

/* Wrapper for big.Int addition, implementing an intermediary accumulator to improve
   performance where an accumulator with possible size > math.MaxInt64 is required. 
   Flushes t_acc buffer when Value() is requested, or when addition could potentially
   exceed max int64 size. 
*/
type BigAccumulator struct {
	t_acc     int64 //intermediate acc
	req_flush bool  //is flush required?
	val       *big.Int
}

//add an int to a bigint, buffers additions and flushes when overflow detected
func (x *BigAccumulator) AddInt(y int) *BigAccumulator {
	n := int64(y)
	if y > 0 && (x.t_acc > (math.MaxInt64 - n)) {
		x.flush()
	} else if y < 0 && (x.t_acc < (math.MinInt64 - n)) {
		x.flush()
	}
	x.t_acc = x.t_acc + n
	x.req_flush = true
	return x
}

func (x *BigAccumulator) flush() {
	if !x.req_flush {
		return
	}
	x.val.Add(x.val, big.NewInt(x.t_acc))
	x.t_acc = 0
	x.req_flush = false
}

func (x *BigAccumulator) Init() *BigAccumulator {
    if x.val == nil {
        x.val = new(big.Int)
    }
    return x
}
//returns underlying big.Int value, after flushing any buffered value
func (x *BigAccumulator) Value() *big.Int {
	if x.req_flush {
		x.flush()
	}
	return x.val
}
