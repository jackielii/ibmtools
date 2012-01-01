package ibmtools

import "math"

const p24 = 2 << 23
// Endianness big endian:
// http://upload.wikimedia.org/wikipedia/commons/d/d5/Endianessmap.svg
// IBM floating point architecture:
// http://en.wikipedia.org/wiki/IBM_Floating_Point_Architecture
// IEEE 32 bit floating point
// http://en.wikipedia.org/wiki/Single_precision_floating-point_format
/*
  1     7               24              width in bits
 +-+-------+------------------------+
 |S|  Exp  |  Fraction              |
 |1|1000010|011101101010000000000000|
 +-+-------+------------------------+
 31 30   24 23                     0    bit index (0 on right)
    bias +64
*/
func Ibm32frombits(b [4]byte) float64 {
	ui := uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	sign := 1 - 2*float64(ui>>31&0x01)
	exponent := math.Pow(16.0, float64(ui>>24&0x7f)-64.0)
	mantissa := float64(ui&0x00ffffff) / float64(p24)
	// println(sign, " ", exponent, " ", mantissa)
	return math.Copysign(mantissa*exponent, sign)
}
