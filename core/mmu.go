// mmu.go
package core

// MMU represents the Memory Management Unit
type MMU struct {
	memory *Memory
}

// NewMMU creates a new MMU instance
func NewMMU() *MMU {
	return &MMU{
		memory: NewMemory(1024 * 1024 * 128), // 128 MB of emulated memory
	}
}

// ReadWord translates a virtual address to a physical address and reads a 32-bit word from memory
func (mmu *MMU) ReadWord(virtAddr uint32) uint32 {
	// For now, assume direct mapping between virtual and physical addresses
	return mmu.memory.ReadWord(virtAddr)
}

// WriteWord translates a virtual address to a physical address and writes a 32-bit word to memory
func (mmu *MMU) WriteWord(virtAddr uint32, value uint32) {
	// For now, assume direct mapping between virtual and physical addresses
	mmu.memory.WriteWord(virtAddr, value)
}
