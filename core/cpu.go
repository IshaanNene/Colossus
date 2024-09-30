// cpu.go
package core

import (
	"colossus/isa"
)

// CPU represents the ARM CPU with registers and control flags
type CPU struct {
	Registers [16]uint32 // General-purpose registers R0-R15 (R15 is the program counter)
	CPSR      uint32     // Current Program Status Register (flags, mode, etc.)
	MMU       *MMU       // Memory Management Unit
	Pipeline  *Pipeline  // Instruction pipeline
}

// NewCPU creates a new CPU instance
func NewCPU() *CPU {
	return &CPU{
		MMU:      NewMMU(),
		Pipeline: NewPipeline(),
	}
}

// Step executes one instruction from the pipeline
func (cpu *CPU) Step() {
	instruction := cpu.Pipeline.Fetch()
	decoded := isa.DecodeInstruction(instruction)
	isa.ExecuteInstruction(decoded)
}
