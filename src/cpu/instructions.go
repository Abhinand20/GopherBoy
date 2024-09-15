package cpu

import (
	"log"
	"os"
)

// OpcodeCycles is the number of cpu cycles for each normal opcode.
var OpcodeCycles = []Cycles{
	1, 3, 2, 2, 1, 1, 2, 1, 5, 2, 2, 2, 1, 1, 2, 1, // 0
	0, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1, // 1
	2, 3, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1, // 2
	2, 3, 2, 2, 3, 3, 3, 1, 2, 2, 2, 2, 1, 1, 2, 1, // 3
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 4
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 5
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 6
	2, 2, 2, 2, 2, 2, 0, 2, 1, 1, 1, 1, 1, 1, 2, 1, // 7
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 8
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 9
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // a
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // b
	2, 3, 3, 4, 3, 4, 2, 4, 2, 4, 3, 0, 3, 6, 2, 4, // c
	2, 3, 3, 0, 3, 4, 2, 4, 2, 4, 3, 0, 3, 0, 2, 4, // d
	3, 3, 2, 0, 0, 4, 2, 4, 4, 1, 4, 0, 0, 0, 2, 4, // e
	3, 3, 2, 1, 0, 4, 2, 4, 3, 2, 4, 1, 0, 0, 2, 4, // f
} //0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f

// CBOpcodeCycles is the number of cpu cycles for each CB opcode.
var CBOpcodeCycles = []Cycles{
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 0
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 1
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 2
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 3
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 4
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 5
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 6
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 7
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 8
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 9
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // A
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // B
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // C
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // D
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // E
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // F
} //0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f


// TODO(abhi): Implement
func (cpu *CPU) setZeroFlag() {
}

// TODO(abhi): Implement
func (cpu *CPU) resetZeroFlag() {
}

// Implements common XOR operations and sets register value using the provided function.
func (cpu *CPU) instrXOR(val byte) {
	a := cpu.AF.Hi()
	a ^= val
	cpu.AF.SetHi(a)
	cpu.setZ(a == 0)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setC(false)
}


/* opcodes to function mappings */
var instructions = [0x100]func(cpu *CPU) {
	0x01: func(cpu *CPU) {
		// LD BC,n16
		val := cpu.popPC16()
		cpu.BC.Set(val)
	},
	0x11: func(cpu *CPU) {
		// LD DE,n16
		val := cpu.popPC16()
		cpu.DE.Set(val)
	},
	0x21: func(cpu *CPU) {
		// LD HL,n16
		val := cpu.popPC16()
		cpu.HL.Set(val)
	},
	0x31: func(cpu *CPU) {
		// LD SP,n16
		val := cpu.popPC16()
		cpu.SP.Set(val)
	},
	0xA8: func(cpu *CPU) {
		// XOR B
		cpu.instrXOR(cpu.BC.Hi())
	},
	0xA9: func(cpu *CPU) {
		// XOR C
		cpu.instrXOR(cpu.BC.Lo())
	},
	0xAA: func(cpu *CPU) {
		// XOR D 
		cpu.instrXOR(cpu.DE.Hi())
	},
	0xAB: func(cpu *CPU) {
		// XOR E
		cpu.instrXOR(cpu.DE.Lo())
	},
	0xAC: func(cpu *CPU) {
		// XOR H
		cpu.instrXOR(cpu.HL.Hi())
	},
	0xAD: func(cpu *CPU) {
		// XOR L
		cpu.instrXOR(cpu.HL.Lo())
	},
	0xAE: func(cpu *CPU) {
		// XOR [HL] 
		val := cpu.MMU.ReadAt(cpu.HL.Value())
		cpu.instrXOR(val)
	},
	0xAF: func(cpu *CPU) {
		// XOR A,A
		cpu.instrXOR(cpu.AF.Hi())
	},
}


func init() {
	for k, v := range instructions {
		if v == nil {
			instructions[k] = func(cpu *CPU) {
				log.Printf("Unimplemented opcode: %#2x", k)
				// TODO(abhi): replace this once debugger is implemented.
				os.Exit(1)
			}
		}
	}
}