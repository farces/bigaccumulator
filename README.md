# Various Golang Packages
A bunch of dumb packages for doing fairly regular things in a more regular way than the standard library is designed for.

## BigAccumulator
BigAccumulator is a wrapper for big.Int to be used for accumulators, with intermediate values buffered and flushed only when necessary.

`x := bufbig.NewBigAccumulator()`<br>
`x.AddInt(12345)`<br>
`fmt.Println(x.Value())`<br>

<b>NewBigAccumulator()</b> creates a new accumulator with members set to their zero values (rather than nil). Use this rather than `new(bufbig.BigAccumulator)` where possible.<br>
<b>Value()</b> returns the underlying big.Int, which can be passed to functions requiring a big.Int explicitly.<br>
<b>AddInt(y int)</b> adds the value y to an internal accumulator, which is flushed when .Value() is requested or when an overflow/underflow would occur.<br>
<b>SetValue(string,base int) bool</b> sets the value of the underlying big.Int to the string given in the base given. Resets the intermediate accumulator so SetValue() can be assumed to return the BigAccumulator to a clean state. Returns a boolean as to whether the call was successful; if unsuccessful, no change is made.<br>
<b>Reset()</b> returns the BigAccumulator to a state where both the big.Int and intermediate accumulator are effectively 0.<br>

## Uses
This package allows big.Int accumulators without the overhead of having to cast from int->int64->big.Int each time a small value is added. Replaces:<br>
`x := big.NewInt(0)`<br>
`x.Add(x,big.NewInt(int64(y)))`<br>
`fmt.Println(x)`<br>
with<br>
`x := bufbig.NewBigAccumulator`<br>
`x.Add(y)`<br>
`fmt.Println(x.Value())`<br>

See bufbig_example.go
