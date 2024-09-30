// core_utils.go
package internal

// Bitwise operations for utility functions

// SignExtend extends a value to the full width of a register
func SignExtend(value uint32, bits int) uint32 {
	if value&(1<<(bits-1)) != 0 {
		value |= ^uint32(0) << bits // Extend sign bit
	}
	return value
}

// RotateRight rotates a 32-bit value right by the given amount
func RotateRight(value uint32, amount int) uint32 {
	return (value >> amount) | (value << (32 - amount))
}
