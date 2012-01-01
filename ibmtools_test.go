package ibmtools

import (
	"testing"
	//	"strconv"
)

var (
	bit = "11000010011101101010000000000000"
	by  = [4]byte{194, 118, 160, 0}
)

// see example in :
// http://en.wikipedia.org/wiki/IBM_Floating_Point_Architecture#An_Example
// 11000010011101101010000000000000 == -118.625
func TestIbm2ieee(t *testing.T) {
	// bit := "11000010011101101010000000000000"
	// var b [4]byte
	// for i := 0; i < 4; i++ {
	// 	n, _ := strconv.ParseUint(bit[i*8:i*8+8], 2, 8)
	// 	b[i] = byte(n)
	// 	print(b[i], ",")
	// }
	if Ibm32frombits(by) != -118.625 {
		t.Error("result is not -118.625, is: ", Ibm32frombits(by))
	}
}

func BenchmarkIbm2ieee(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ibm32frombits(by)
	}
}
