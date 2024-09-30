// peripherals.go
package io

import (
	"fmt"
)

// Peripheral represents a generic peripheral device in the emulator
type Peripheral interface {
	Read(address uint32) uint32
	Write(address uint32, value uint32)
}

// Timer is an example peripheral that can be controlled by the ARM processor
type Timer struct {
	counter uint32
}

// NewTimer creates a new Timer peripheral
func NewTimer() *Timer {
	return &Timer{
		counter: 0,
	}
}

// Read reads the timer value
func (t *Timer) Read(address uint32) uint32 {
	// For simplicity, we can just return the counter value
	fmt.Printf("Reading from Timer at address %x\n", address)
	return t.counter
}

// Write sets the timer value
func (t *Timer) Write(address uint32, value uint32) {
	fmt.Printf("Writing to Timer at address %x, value: %d\n", address, value)
	t.counter = value
}

// PeripheralBus manages the collection of peripherals
type PeripheralBus struct {
	peripherals map[uint32]Peripheral // Maps addresses to peripherals
}

// NewPeripheralBus creates a new PeripheralBus
func NewPeripheralBus() *PeripheralBus {
	return &PeripheralBus{
		peripherals: make(map[uint32]Peripheral),
	}
}

// AddPeripheral adds a peripheral to the bus
func (pb *PeripheralBus) AddPeripheral(address uint32, peripheral Peripheral) {
	pb.peripherals[address] = peripheral
}

// Read reads from a peripheral at the specified address
func (pb *PeripheralBus) Read(address uint32) uint32 {
	if peripheral, ok := pb.peripherals[address]; ok {
		return peripheral.Read(address)
	}
	fmt.Printf("No peripheral found at address %x\n", address)
	return 0
}

// Write writes to a peripheral at the specified address
func (pb *PeripheralBus) Write(address uint32, value uint32) {
	if peripheral, ok := pb.peripherals[address]; ok {
		peripheral.Write(address, value)
	} else {
		fmt.Printf("No peripheral found at address %x\n", address)
	}
}
