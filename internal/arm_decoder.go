// arm_decoder.go
package internal

import (
	"colossus/isa"
)

// DecodeInstruction decodes a 32-bit ARM instruction into a structured representation
func DecodeInstruction(instruction uint32) isa.ARMv7Instruction {
	// For simplicity, we'll assume a fixed instruction format
	op := (instruction >> 21) & 0xF   // Extract opcode
	cond := (instruction >> 28) & 0xF // Extract condition code
	rn := (instruction >> 16) & 0xF   // Extract first operand register
	rd := (instruction >> 12) & 0xF   // Extract destination register
	rs := (instruction >> 8) & 0xF    // Extract second operand register
	immed := instruction & 0xFF       // Extract immediate value

	return isa.ARMv7Instruction{
		Opcode:    isa.Opcode(op),
		Cond:      isa.Condition(cond),
		Rn:        rn,
		Rd:        rd,
		Rs:        rs,
		Immediate: immed,
	}
}
