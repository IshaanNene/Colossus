// thumb.go
package isa

// Thumb instruction set opcodes (16-bit encoding)
type ThumbOpcode int

const (
	// Thumb Data Processing Instructions
	THUMB_AND ThumbOpcode = iota // 0000 - Bitwise AND
	THUMB_EOR                    // 0001 - Bitwise Exclusive OR
	THUMB_LSL                    // 0010 - Logical Shift Left
	THUMB_LSR                    // 0011 - Logical Shift Right
	THUMB_ASR                    // 0100 - Arithmetic Shift Right
	THUMB_ADD                    // 0101 - Add
	THUMB_SUB                    // 0110 - Subtract
	THUMB_CMP                    // 0111 - Compare
	THUMB_MOV                    // 1000 - Move
	THUMB_B                      // 1001 - Branch
)

// ThumbInstruction represents a decoded Thumb instruction
type ThumbInstruction struct {
	Opcode    ThumbOpcode // Thumb instruction opcode
	Rd        int         // Destination register
	Rn        int         // Source register (if applicable)
	Immediate int         // Immediate value (if applicable)
}

// DecodeThumbInstruction decodes a 16-bit Thumb instruction
func DecodeThumbInstruction(instruction uint16) ThumbInstruction {
	// Extract opcode (bits 15-11)
	opcode := ThumbOpcode((instruction >> 11) & 0x1F)

	// Extract Rd (destination register) (bits 2-0)
	rd := int(instruction & 0x7)

	// Extract Rn (source register) or immediate value depending on instruction
	rn := int((instruction >> 3) & 0x7)
	immediate := int(instruction & 0xFF)

	return ThumbInstruction{
		Opcode:    opcode,
		Rd:        rd,
		Rn:        rn,
		Immediate: immediate,
	}
}

func ExecuteThumbInstruction(instr ThumbInstruction) {
	switch instr.Opcode {
	case THUMB_AND:
	case THUMB_EOR:
	case THUMB_LSL:
	case THUMB_LSR:
	case THUMB_ASR:
	case THUMB_ADD:
	case THUMB_SUB:
	case THUMB_CMP:
	case THUMB_MOV:
	case THUMB_B:
	default:
		panic("Unknown Thumb opcode")
	}
}
