package core

import "colossus/isa"

type Pipeline struct {
	fetchedInstruction uint32
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (pipeline *Pipeline) Fetch(cpu *CPU) {
	pipeline.fetchedInstruction = cpu.MMU.ReadWord(cpu.Registers[15])
	cpu.Registers[15] += 4
}

func (pipeline *Pipeline) Decode() isa.Instruction {
	return isa.DecodeInstruction(pipeline.fetchedInstruction)
}

func (pipeline *Pipeline) Execute(cpu *CPU) {
	instruction := pipeline.Decode()
	var handler InstructionHandler

	switch instruction.(type) {
	case isa.ARMv7Instruction:
		handler = &ARMv7InstructionHandler{Instruction: instruction.(isa.ARMv7Instruction)}
	case isa.ThumbInstruction:
		handler = &ThumbInstructionHandler{Instruction: instruction.(isa.ThumbInstruction)}
	default:
		return
	}

	handler.Execute(cpu)
}
