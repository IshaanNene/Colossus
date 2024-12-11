# ARM Emulator (Colossus)

## Overview

Colossus is an ARM emulator designed to simulate the ARM architecture, allowing for the execution of ARM instructions. This project includes components for instruction decoding, memory management, and peripheral handling, making it a comprehensive tool for ARM instruction testing and development.

## Features

- **Instruction Decoding**: Decodes ARMv7 instructions into structured representations.
- **Memory Management**: Provides a memory model that supports reading and writing of words and bytes.
- **Pipeline Architecture**: Implements a basic instruction pipeline for fetching, decoding, and executing instructions.
- **Peripheral Support**: Allows for the integration of peripheral devices, such as timers, into the emulator.
- **Logging**: Includes a logging utility for tracking emulator operations and errors.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/IshaanNene/Colossus
   cd colossus
   ```

2. Build the emulator:
   ```bash
   ./scripts/build.sh
   ```

### Usage

1. Load a program into the CPU:
   ```go
   cpu.LoadProgram([]uint32{0xE3A00001, 0xE3A01002}) // Example instructions
   ```

2. Execute the program:
   ```go
   cpu.ExecuteProgram()
   ```

3. Dump the CPU registers to see the state after execution:
   ```go
   cpu.DumpRegisters()
   ```
## Code Structure

- **cmd/**: Contains the main entry point for the emulator.
- **pkg/**: Contains utility packages such as logging and caching.
- **core/**: Implements the core functionality of the emulator, including CPU, memory, and instruction handling.
- **isa/**: Defines the instruction set architecture, including ARMv7 and Thumb instructions.
- **io/**: Manages peripheral devices and their interactions with the CPU.
- **internal/**: Contains internal utilities and helpers for the emulator.
- **test/**: Contains unit tests for various components of the emulator.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

- Inspired by various ARM architecture resources and emulation projects.
- Thanks to the open-source community for their contributions and support.
