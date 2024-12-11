package core

import (
	"colossus/isa"
	"fmt"
)

type CPU struct {
	Registers [16]uint32
	CPSR      uint32
	MMU       *MMU
	Pipeline  *Pipeline
}

func NewCPU() *CPU {
	return &CPU{
		MMU:      NewMMU(),
		Pipeline: NewPipeline(),
	}
}

func (cpu *CPU) Step() {
	instruction := cpu.Pipeline.Fetch()
	decoded := isa.DecodeInstruction(instruction)
	isa.ExecuteInstruction(decoded)
	cpu.UpdateCPSR()
}

func (cpu *CPU) UpdateCPSR() {
	cpu.CPSR = 0
	if cpu.Registers[0] == 0 {
		cpu.CPSR |= 0x1 // Set zero flag
	}
	if cpu.Registers[1] > 0 {
		cpu.CPSR |= 0x2 // Set positive flag
	}
	if cpu.Registers[2] < 0 {
		cpu.CPSR |= 0x4 // Set negative flag
	}
	fmt.Println("CPSR updated:", cpu.CPSR)
}

func (cpu *CPU) Reset() {
	cpu.Registers = [16]uint32{}
	cpu.CPSR = 0
	cpu.MMU.Reset()
	cpu.Pipeline.Reset()
	fmt.Println("CPU has been reset.")
}

func (cpu *CPU) LoadProgram(program []uint32) {
	for i, instruction := range program {
		if i < len(cpu.Registers) {
			cpu.Registers[i] = instruction
		}
	}
	fmt.Println("Program loaded into CPU registers.")
}

func (cpu *CPU) DumpRegisters() {
	for i, reg := range cpu.Registers {
		fmt.Printf("R%d: %d\n", i, reg)
	}
}

func (cpu *CPU) ExecuteProgram() {
	for {
		cpu.Step()
		if cpu.CPSR&0x1 != 0 { // Check for zero flag
			break
		}
	}
	fmt.Println("Program execution completed.")
}
