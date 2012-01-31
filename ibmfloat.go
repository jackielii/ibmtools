package ibmtools

const e24 = 16777216.0 // 2^24
const one uint64 = 1

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
	var sign float64
	if b[0]>>7 == 0 {
		sign = 1.0
	} else {
		sign = -1.0
	}
	return sign * float64(one<<(4*(b[0]&0x7f-64))) *
		float64(uint32(b[3])|uint32(b[2])<<8|uint32(b[1])<<16) / e24
	//exponent := 1 << (4 * (b[0]&0x7f - 64))
	//mantissa := float64(uint32(b[3])|uint32(b[2])<<8|uint32(b[1])<<16) / e24
	//return sign*mantissa*float64(exponent)
}
