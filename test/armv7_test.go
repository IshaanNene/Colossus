// armv7_test.go
package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInstructionDecoding tests the instruction decoding functionality for ARMv7
func TestInstructionDecoding(t *testing.T) {
	// Example instruction for testing
	instruction := uint32(0b11000011000000000000000000000000) // Example opcode
	expectedOpcode := AND                                     // Assume AND is defined in your opcode implementation

	// Decode the instruction (replace with actual decode function)
	decodedOpcode := DecodeInstruction(instruction)

	assert.Equal(t, expectedOpcode, decodedOpcode, "Expected AND opcode")
}

// TestMemoryAccess tests memory read and write operations
func TestMemoryAccess(t *testing.T) {
	mem := NewMemory() // Assume NewMemory initializes memory

	// Write a value to memory
	address := uint32(0x1000)
	value := uint32(42)
	mem.Write(address, value)

	// Read the value back
	readValue := mem.Read(address)
	assert.Equal(t, value, readValue, "Expected value to be 42")
}

// Additional tests can be added here
