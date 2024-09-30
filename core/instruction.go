// instruction.go
package core

import "colossus/isa"

// InstructionHandler represents a decoded instruction that the CPU can execute
type InstructionHandler interface {
	Execute(cpu *CPU)
}

// ARMv7InstructionHandler handles ARMv7 instructions
type ARMv7InstructionHandler struct {
	Instruction isa.ARMv7Instruction
}

// Execute executes an ARMv7 instruction on the CPU
func (handler *ARMv7InstructionHandler) Execute(cpu *CPU) {
	isa.ExecuteInstruction(handler.Instruction)
}

// ThumbInstructionHandler handles Thumb instructions
type ThumbInstructionHandler struct {
	Instruction isa.ThumbInstruction
}

// Execute executes a Thumb instruction on the CPU
func (handler *ThumbInstructionHandler) Execute(cpu *CPU) {
	isa.ExecuteThumbInstruction(handler.Instruction)
}
