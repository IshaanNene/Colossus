package core

type Memory struct {
	data []byte
}

func NewMemory(size int) *Memory {
	return &Memory{
		data: make([]byte, size),
	}
}

func (mem *Memory) ReadWord(addr uint32) (uint32, error) {
	if addr+3 >= uint32(len(mem.data)) {
		return 0, fmt.Errorf("address out of bounds: %d", addr)
	}
	return uint32(mem.data[addr]) |
		uint32(mem.data[addr+1])<<8 |
		uint32(mem.data[addr+2])<<16 |
		uint32(mem.data[addr+3])<<24, nil
}

func (mem *Memory) WriteWord(addr uint32, value uint32) error {
	if addr+3 >= uint32(len(mem.data)) {
		return fmt.Errorf("address out of bounds: %d", addr)
	}
	mem.data[addr] = byte(value & 0xFF)
	mem.data[addr+1] = byte((value >> 8) & 0xFF)
	mem.data[addr+2] = byte((value >> 16) & 0xFF)
	mem.data[addr+3] = byte((value >> 24) & 0xFF)
	return nil
}

func (mem *Memory) ReadBytes(addr uint32, length int) ([]byte, error) {
	if addr+uint32(length) > uint32(len(mem.data)) {
		return nil, fmt.Errorf("address out of bounds: %d", addr)
	}
	return mem.data[addr : addr+uint32(length)], nil
}

func (mem *Memory) WriteBytes(addr uint32, bytes []byte) error {
	if addr+uint32(len(bytes)) > uint32(len(mem.data)) {
		return fmt.Errorf("address out of bounds: %d", addr)
	}
	copy(mem.data[addr:], bytes)
	return nil
}

func (mem *Memory) Clear() {
	for i := range mem.data {
		mem.data[i] = 0
	}
}

func (mem *Memory) Size() int {
	return len(mem.data)
}
