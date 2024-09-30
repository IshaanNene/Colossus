// memory.go
package core

// Memory represents the memory for the CPU
type Memory struct {
	data []byte
}

// NewMemory creates a new Memory instance with a specified size
func NewMemory(size int) *Memory {
	return &Memory{
		data: make([]byte, size),
	}
}

// ReadWord reads a 32-bit word from memory at the given address
func (mem *Memory) ReadWord(addr uint32) uint32 {
	return uint32(mem.data[addr]) |
		uint32(mem.data[addr+1])<<8 |
		uint32(mem.data[addr+2])<<16 |
		uint32(mem.data[addr+3])<<24
}

// WriteWord writes a 32-bit word to memory at the given address
func (mem *Memory) WriteWord(addr uint32, value uint32) {
	mem.data[addr] = byte(value & 0xFF)
	mem.data[addr+1] = byte((value >> 8) & 0xFF)
	mem.data[addr+2] = byte((value >> 16) & 0xFF)
	mem.data[addr+3] = byte((value >> 24) & 0xFF)
}
