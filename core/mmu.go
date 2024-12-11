package core

type MMU struct {
	memory *Memory
}

func NewMMU() *MMU {
	return &MMU{
		memory: NewMemory(1024 * 1024 * 128),
	}
}

func (mmu *MMU) ReadWord(virtAddr uint32) (uint32, error) {
	if virtAddr >= 0 && virtAddr < 1024*1024*128 {
		return mmu.memory.ReadWord(virtAddr), nil
	}
	return 0, fmt.Errorf("invalid virtual address: %d", virtAddr)
}

func (mmu *MMU) WriteWord(virtAddr uint32, value uint32) error {
	if virtAddr >= 0 && virtAddr < 1024*1024*128 {
		mmu.memory.WriteWord(virtAddr, value)
		return nil
	}
	return fmt.Errorf("invalid virtual address: %d", virtAddr)
}
