package bits

import "golang.org/x/exp/constraints"

// BitCount32
// https://stackoverflow.com/questions/109023/count-the-number-of-set-bits-in-a-32-bit-integer
// https://en.wikipedia.org/wiki/Hamming_weight
func BitCount32[T uint32 | int32](i T) T {
	i = i - ((i >> 1) & 0x55555555)                // add pairs of bits
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333) // quads
	i = (i + (i >> 4)) & 0x0F0F0F0F                // groups of 8
	return T((i * 0x01010101) >> 24)               // horizontal sum of bytes
}

func Xor[T constraints.Integer](left, right T) T {
	return (left | right) & ^(left & right)
}

func CheckBit[T constraints.Integer](value, index T) bool {
	return (value & (1 << index)) == (1 << index)
}

func SetBit[T constraints.Integer](value T, index T) T {
	return value | (1 << index)
}

func ClearBit[T constraints.Integer](value T, index T) T {
	return value & ^(1 << index)
}

func ToggleBit[T constraints.Integer](value T, index T) T {
	return Xor(value, 1<<index)
}
