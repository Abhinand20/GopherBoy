package cpu

import (
	"gopherboy/common"
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

func (cpu *CPU) instrDECr16(r *Register) {
	val := r.Value()
	r.Set(val - 1)
}

func (cpu *CPU) instrINCr16(r *Register) {
	val := r.Value()
	r.Set(val + 1)
}

func (cpu *CPU) instrLDr8(setFunc func(byte), dest, src byte) {
	
}

func (cpu *CPU) instrLDr16(setFunc func(uint16), dest, src uint16) {
	
}

func (cpu *CPU) instrLDn8(dest uint16, src byte) {
	cpu.MMU.WriteAt(dest, src)
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
	/* LD 8-bit */
	0x02: func(cpu *CPU) {
		// LD [BC], A
		cpu.instrLDn8(cpu.BC.Value(), cpu.AF.Hi())
	},
	0x12: func(cpu *CPU) {
		// LD [DE], A
		cpu.instrLDn8(cpu.DE.Value(), cpu.AF.Hi())
	},
	0x22: func(cpu *CPU) {
		// LD [HLI], A
		cpu.instrLDn8(cpu.HL.Value(), cpu.AF.Hi())
		cpu.instrINCr16(&cpu.HL)
	},
	0x20: nil,
	// 0x20: func(cpu *CPU) {
	// 	// JR NZ, e8
	// 	if cpu.testZ() {
	// 		log.Printf("Took jump\n")
	// 		offset := int8(cpu.popPC8())
	// 		currAddr := int16(cpu.PC)
	// 		cpu.PC = uint16(currAddr + int16(offset))
	// 		log.Printf("Now at %x", cpu.PC)
	// 	}
	// },
	0x32: func(cpu *CPU) {
		// LD [HLD], A
		cpu.instrLDn8(cpu.HL.Value(), cpu.AF.Hi())
		cpu.instrDECr16(&cpu.HL)
	},
	/* XOR */
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
		// XOR A
		cpu.instrXOR(cpu.AF.Hi())
	},
	0xEE: func(cpu *CPU) {
		// XOR n8
		n8 := cpu.popPC8()
		cpu.instrXOR(n8)
	},
}

var cbInstructions = [0x100]func(cpu *CPU) {
	0x00: nil,
	/*
	0x00: func(cpu *CPU) {
		// RLC B
	},
	*/
	0x01: nil,
	/*
	0x01: func(cpu *CPU) {
		// RLC C
	},
	*/
	0x02: nil,
	/*
	0x02: func(cpu *CPU) {
		// RLC D
	},
	*/
	0x03: nil,
	/*
	0x03: func(cpu *CPU) {
		// RLC E
	},
	*/
	0x04: nil,
	/*
	0x04: func(cpu *CPU) {
		// RLC H
	},
	*/
	0x05: nil,
	/*
	0x05: func(cpu *CPU) {
		// RLC L
	},
	*/
	0x06: nil,
	/*
	0x06: func(cpu *CPU) {
		// RLC [HL]
	},
	*/
	0x07: nil,
	/*
	0x07: func(cpu *CPU) {
		// RLC A
	},
	*/
	0x08: nil,
	/*
	0x08: func(cpu *CPU) {
		// RRC B
	},
	*/
	0x09: nil,
	/*
	0x09: func(cpu *CPU) {
		// RRC C
	},
	*/
	0x0A: nil,
	/*
	0x0A: func(cpu *CPU) {
		// RRC D
	},
	*/
	0x0B: nil,
	/*
	0x0B: func(cpu *CPU) {
		// RRC E
	},
	*/
	0x0C: nil,
	/*
	0x0C: func(cpu *CPU) {
		// RRC H
	},
	*/
	0x0D: nil,
	/*
	0x0D: func(cpu *CPU) {
		// RRC L
	},
	*/
	0x0E: nil,
	/*
	0x0E: func(cpu *CPU) {
		// RRC [HL]
	},
	*/
	0x0F: nil,
	/*
	0x0F: func(cpu *CPU) {
		// RRC A
	},
	*/
	0x10: nil,
	/*
	0x10: func(cpu *CPU) {
		// RL B
	},
	*/
	0x11: nil,
	/*
	0x11: func(cpu *CPU) {
		// RL C
	},
	*/
	0x12: nil,
	/*
	0x12: func(cpu *CPU) {
		// RL D
	},
	*/
	0x13: nil,
	/*
	0x13: func(cpu *CPU) {
		// RL E
	},
	*/
	0x14: nil,
	/*
	0x14: func(cpu *CPU) {
		// RL H
	},
	*/
	0x15: nil,
	/*
	0x15: func(cpu *CPU) {
		// RL L
	},
	*/
	0x16: nil,
	/*
	0x16: func(cpu *CPU) {
		// RL [HL]
	},
	*/
	0x17: nil,
	/*
	0x17: func(cpu *CPU) {
		// RL A
	},
	*/
	0x18: nil,
	/*
	0x18: func(cpu *CPU) {
		// RR B
	},
	*/
	0x19: nil,
	/*
	0x19: func(cpu *CPU) {
		// RR C
	},
	*/
	0x1A: nil,
	/*
	0x1A: func(cpu *CPU) {
		// RR D
	},
	*/
	0x1B: nil,
	/*
	0x1B: func(cpu *CPU) {
		// RR E
	},
	*/
	0x1C: nil,
	/*
	0x1C: func(cpu *CPU) {
		// RR H
	},
	*/
	0x1D: nil,
	/*
	0x1D: func(cpu *CPU) {
		// RR L
	},
	*/
	0x1E: nil,
	/*
	0x1E: func(cpu *CPU) {
		// RR [HL]
	},
	*/
	0x1F: nil,
	/*
	0x1F: func(cpu *CPU) {
		// RR A
	},
	*/
	0x20: nil,
	/*
	0x20: func(cpu *CPU) {
		// SLA B
	},
	*/
	0x21: nil,
	/*
	0x21: func(cpu *CPU) {
		// SLA C
	},
	*/
	0x22: nil,
	/*
	0x22: func(cpu *CPU) {
		// SLA D
	},
	*/
	0x23: nil,
	/*
	0x23: func(cpu *CPU) {
		// SLA E
	},
	*/
	0x24: nil,
	/*
	0x24: func(cpu *CPU) {
		// SLA H
	},
	*/
	0x25: nil,
	/*
	0x25: func(cpu *CPU) {
		// SLA L
	},
	*/
	0x26: nil,
	/*
	0x26: func(cpu *CPU) {
		// SLA [HL]
	},
	*/
	0x27: nil,
	/*
	0x27: func(cpu *CPU) {
		// SLA A
	},
	*/
	0x28: nil,
	/*
	0x28: func(cpu *CPU) {
		// SRA B
	},
	*/
	0x29: nil,
	/*
	0x29: func(cpu *CPU) {
		// SRA C
	},
	*/
	0x2A: nil,
	/*
	0x2A: func(cpu *CPU) {
		// SRA D
	},
	*/
	0x2B: nil,
	/*
	0x2B: func(cpu *CPU) {
		// SRA E
	},
	*/
	0x2C: nil,
	/*
	0x2C: func(cpu *CPU) {
		// SRA H
	},
	*/
	0x2D: nil,
	/*
	0x2D: func(cpu *CPU) {
		// SRA L
	},
	*/
	0x2E: nil,
	/*
	0x2E: func(cpu *CPU) {
		// SRA [HL]
	},
	*/
	0x2F: nil,
	/*
	0x2F: func(cpu *CPU) {
		// SRA A
	},
	*/
	0x30: nil,
	/*
	0x30: func(cpu *CPU) {
		// SWAP B
	},
	*/
	0x31: nil,
	/*
	0x31: func(cpu *CPU) {
		// SWAP C
	},
	*/
	0x32: nil,
	/*
	0x32: func(cpu *CPU) {
		// SWAP D
	},
	*/
	0x33: nil,
	/*
	0x33: func(cpu *CPU) {
		// SWAP E
	},
	*/
	0x34: nil,
	/*
	0x34: func(cpu *CPU) {
		// SWAP H
	},
	*/
	0x35: nil,
	/*
	0x35: func(cpu *CPU) {
		// SWAP L
	},
	*/
	0x36: nil,
	/*
	0x36: func(cpu *CPU) {
		// SWAP [HL]
	},
	*/
	0x37: nil,
	/*
	0x37: func(cpu *CPU) {
		// SWAP A
	},
	*/
	0x38: nil,
	/*
	0x38: func(cpu *CPU) {
		// SRL B
	},
	*/
	0x39: nil,
	/*
	0x39: func(cpu *CPU) {
		// SRL C
	},
	*/
	0x3A: nil,
	/*
	0x3A: func(cpu *CPU) {
		// SRL D
	},
	*/
	0x3B: nil,
	/*
	0x3B: func(cpu *CPU) {
		// SRL E
	},
	*/
	0x3C: nil,
	/*
	0x3C: func(cpu *CPU) {
		// SRL H
	},
	*/
	0x3D: nil,
	/*
	0x3D: func(cpu *CPU) {
		// SRL L
	},
	*/
	0x3E: nil,
	/*
	0x3E: func(cpu *CPU) {
		// SRL [HL]
	},
	*/
	0x3F: nil,
	/*
	0x3F: func(cpu *CPU) {
		// SRL A
	},
	*/
	0x40: nil,
	/*
	0x40: func(cpu *CPU) {
		// BIT 0,B
	},
	*/
	0x41: nil,
	/*
	0x41: func(cpu *CPU) {
		// BIT 0,C
	},
	*/
	0x42: nil,
	/*
	0x42: func(cpu *CPU) {
		// BIT 0,D
	},
	*/
	0x43: nil,
	/*
	0x43: func(cpu *CPU) {
		// BIT 0,E
	},
	*/
	0x44: nil,
	/*
	0x44: func(cpu *CPU) {
		// BIT 0,H
	},
	*/
	0x45: nil,
	/*
	0x45: func(cpu *CPU) {
		// BIT 0,L
	},
	*/
	0x46: nil,
	/*
	0x46: func(cpu *CPU) {
		// BIT 0,[HL]
	},
	*/
	0x47: nil,
	/*
	0x47: func(cpu *CPU) {
		// BIT 0,A
	},
	*/
	0x48: nil,
	/*
	0x48: func(cpu *CPU) {
		// BIT 1,B
	},
	*/
	0x49: nil,
	/*
	0x49: func(cpu *CPU) {
		// BIT 1,C
	},
	*/
	0x4A: nil,
	/*
	0x4A: func(cpu *CPU) {
		// BIT 1,D
	},
	*/
	0x4B: nil,
	/*
	0x4B: func(cpu *CPU) {
		// BIT 1,E
	},
	*/
	0x4C: nil,
	/*
	0x4C: func(cpu *CPU) {
		// BIT 1,H
	},
	*/
	0x4D: nil,
	/*
	0x4D: func(cpu *CPU) {
		// BIT 1,L
	},
	*/
	0x4E: nil,
	/*
	0x4E: func(cpu *CPU) {
		// BIT 1,[HL]
	},
	*/
	0x4F: nil,
	/*
	0x4F: func(cpu *CPU) {
		// BIT 1,A
	},
	*/
	0x50: nil,
	/*
	0x50: func(cpu *CPU) {
		// BIT 2,B
	},
	*/
	0x51: nil,
	/*
	0x51: func(cpu *CPU) {
		// BIT 2,C
	},
	*/
	0x52: nil,
	/*
	0x52: func(cpu *CPU) {
		// BIT 2,D
	},
	*/
	0x53: nil,
	/*
	0x53: func(cpu *CPU) {
		// BIT 2,E
	},
	*/
	0x54: nil,
	/*
	0x54: func(cpu *CPU) {
		// BIT 2,H
	},
	*/
	0x55: nil,
	/*
	0x55: func(cpu *CPU) {
		// BIT 2,L
	},
	*/
	0x56: nil,
	/*
	0x56: func(cpu *CPU) {
		// BIT 2,[HL]
	},
	*/
	0x57: nil,
	/*
	0x57: func(cpu *CPU) {
		// BIT 2,A
	},
	*/
	0x58: nil,
	/*
	0x58: func(cpu *CPU) {
		// BIT 3,B
	},
	*/
	0x59: nil,
	/*
	0x59: func(cpu *CPU) {
		// BIT 3,C
	},
	*/
	0x5A: nil,
	/*
	0x5A: func(cpu *CPU) {
		// BIT 3,D
	},
	*/
	0x5B: nil,
	/*
	0x5B: func(cpu *CPU) {
		// BIT 3,E
	},
	*/
	0x5C: nil,
	/*
	0x5C: func(cpu *CPU) {
		// BIT 3,H
	},
	*/
	0x5D: nil,
	/*
	0x5D: func(cpu *CPU) {
		// BIT 3,L
	},
	*/
	0x5E: nil,
	/*
	0x5E: func(cpu *CPU) {
		// BIT 3,[HL]
	},
	*/
	0x5F: nil,
	/*
	0x5F: func(cpu *CPU) {
		// BIT 3,A
	},
	*/
	0x60: nil,
	/*
	0x60: func(cpu *CPU) {
		// BIT 4,B
	},
	*/
	0x61: nil,
	/*
	0x61: func(cpu *CPU) {
		// BIT 4,C
	},
	*/
	0x62: nil,
	/*
	0x62: func(cpu *CPU) {
		// BIT 4,D
	},
	*/
	0x63: nil,
	/*
	0x63: func(cpu *CPU) {
		// BIT 4,E
	},
	*/
	0x64: nil,
	/*
	0x64: func(cpu *CPU) {
		// BIT 4,H
	},
	*/
	0x65: nil,
	/*
	0x65: func(cpu *CPU) {
		// BIT 4,L
	},
	*/
	0x66: nil,
	/*
	0x66: func(cpu *CPU) {
		// BIT 4,[HL]
	},
	*/
	0x67: nil,
	/*
	0x67: func(cpu *CPU) {
		// BIT 4,A
	},
	*/
	0x68: nil,
	/*
	0x68: func(cpu *CPU) {
		// BIT 5,B
	},
	*/
	0x69: nil,
	/*
	0x69: func(cpu *CPU) {
		// BIT 5,C
	},
	*/
	0x6A: nil,
	/*
	0x6A: func(cpu *CPU) {
		// BIT 5,D
	},
	*/
	0x6B: nil,
	/*
	0x6B: func(cpu *CPU) {
		// BIT 5,E
	},
	*/
	0x6C: nil,
	/*
	0x6C: func(cpu *CPU) {
		// BIT 5,H
	},
	*/
	0x6D: nil,
	/*
	0x6D: func(cpu *CPU) {
		// BIT 5,L
	},
	*/
	0x6E: nil,
	/*
	0x6E: func(cpu *CPU) {
		// BIT 5,[HL]
	},
	*/
	0x6F: nil,
	/*
	0x6F: func(cpu *CPU) {
		// BIT 5,A
	},
	*/
	0x70: nil,
	/*
	0x70: func(cpu *CPU) {
		// BIT 6,B
	},
	*/
	0x71: nil,
	/*
	0x71: func(cpu *CPU) {
		// BIT 6,C
	},
	*/
	0x72: nil,
	/*
	0x72: func(cpu *CPU) {
		// BIT 6,D
	},
	*/
	0x73: nil,
	/*
	0x73: func(cpu *CPU) {
		// BIT 6,E
	},
	*/
	0x74: nil,
	/*
	0x74: func(cpu *CPU) {
		// BIT 6,H
	},
	*/
	0x75: nil,
	/*
	0x75: func(cpu *CPU) {
		// BIT 6,L
	},
	*/
	0x76: nil,
	/*
	0x76: func(cpu *CPU) {
		// BIT 6,[HL]
	},
	*/
	0x77: nil,
	/*
	0x77: func(cpu *CPU) {
		// BIT 6,A
	},
	*/
	0x78: nil,
	/*
	0x78: func(cpu *CPU) {
		// BIT 7,B
	},
	*/
	0x79: nil,
	/*
	0x79: func(cpu *CPU) {
		// BIT 7,C
	},
	*/
	0x7A: nil,
	/*
	0x7A: func(cpu *CPU) {
		// BIT 7,D
	},
	*/
	0x7B: nil,
	/*
	0x7B: func(cpu *CPU) {
		// BIT 7,E
	},
	*/
	0x7C: func(cpu *CPU) {
		// BIT 7,H
		on := common.TestBitAtIndex(cpu.HL.Hi(), 7)
		cpu.setZ(!on)
		cpu.setN(false)
		cpu.setH(true)
	},
	0x7D: nil,
	/*
	0x7D: func(cpu *CPU) {
		// BIT 7,L
	},
	*/
	0x7E: nil,
	/*
	0x7E: func(cpu *CPU) {
		// BIT 7,[HL]
	},
	*/
	0x7F: nil,
	/*
	0x7F: func(cpu *CPU) {
		// BIT 7,A
	},
	*/
	0x80: nil,
	/*
	0x80: func(cpu *CPU) {
		// RES 0,B
	},
	*/
	0x81: nil,
	/*
	0x81: func(cpu *CPU) {
		// RES 0,C
	},
	*/
	0x82: nil,
	/*
	0x82: func(cpu *CPU) {
		// RES 0,D
	},
	*/
	0x83: nil,
	/*
	0x83: func(cpu *CPU) {
		// RES 0,E
	},
	*/
	0x84: nil,
	/*
	0x84: func(cpu *CPU) {
		// RES 0,H
	},
	*/
	0x85: nil,
	/*
	0x85: func(cpu *CPU) {
		// RES 0,L
	},
	*/
	0x86: nil,
	/*
	0x86: func(cpu *CPU) {
		// RES 0,[HL]
	},
	*/
	0x87: nil,
	/*
	0x87: func(cpu *CPU) {
		// RES 0,A
	},
	*/
	0x88: nil,
	/*
	0x88: func(cpu *CPU) {
		// RES 1,B
	},
	*/
	0x89: nil,
	/*
	0x89: func(cpu *CPU) {
		// RES 1,C
	},
	*/
	0x8A: nil,
	/*
	0x8A: func(cpu *CPU) {
		// RES 1,D
	},
	*/
	0x8B: nil,
	/*
	0x8B: func(cpu *CPU) {
		// RES 1,E
	},
	*/
	0x8C: nil,
	/*
	0x8C: func(cpu *CPU) {
		// RES 1,H
	},
	*/
	0x8D: nil,
	/*
	0x8D: func(cpu *CPU) {
		// RES 1,L
	},
	*/
	0x8E: nil,
	/*
	0x8E: func(cpu *CPU) {
		// RES 1,[HL]
	},
	*/
	0x8F: nil,
	/*
	0x8F: func(cpu *CPU) {
		// RES 1,A
	},
	*/
	0x90: nil,
	/*
	0x90: func(cpu *CPU) {
		// RES 2,B
	},
	*/
	0x91: nil,
	/*
	0x91: func(cpu *CPU) {
		// RES 2,C
	},
	*/
	0x92: nil,
	/*
	0x92: func(cpu *CPU) {
		// RES 2,D
	},
	*/
	0x93: nil,
	/*
	0x93: func(cpu *CPU) {
		// RES 2,E
	},
	*/
	0x94: nil,
	/*
	0x94: func(cpu *CPU) {
		// RES 2,H
	},
	*/
	0x95: nil,
	/*
	0x95: func(cpu *CPU) {
		// RES 2,L
	},
	*/
	0x96: nil,
	/*
	0x96: func(cpu *CPU) {
		// RES 2,[HL]
	},
	*/
	0x97: nil,
	/*
	0x97: func(cpu *CPU) {
		// RES 2,A
	},
	*/
	0x98: nil,
	/*
	0x98: func(cpu *CPU) {
		// RES 3,B
	},
	*/
	0x99: nil,
	/*
	0x99: func(cpu *CPU) {
		// RES 3,C
	},
	*/
	0x9A: nil,
	/*
	0x9A: func(cpu *CPU) {
		// RES 3,D
	},
	*/
	0x9B: nil,
	/*
	0x9B: func(cpu *CPU) {
		// RES 3,E
	},
	*/
	0x9C: nil,
	/*
	0x9C: func(cpu *CPU) {
		// RES 3,H
	},
	*/
	0x9D: nil,
	/*
	0x9D: func(cpu *CPU) {
		// RES 3,L
	},
	*/
	0x9E: nil,
	/*
	0x9E: func(cpu *CPU) {
		// RES 3,[HL]
	},
	*/
	0x9F: nil,
	/*
	0x9F: func(cpu *CPU) {
		// RES 3,A
	},
	*/
	0xA0: nil,
	/*
	0xA0: func(cpu *CPU) {
		// RES 4,B
	},
	*/
	0xA1: nil,
	/*
	0xA1: func(cpu *CPU) {
		// RES 4,C
	},
	*/
	0xA2: nil,
	/*
	0xA2: func(cpu *CPU) {
		// RES 4,D
	},
	*/
	0xA3: nil,
	/*
	0xA3: func(cpu *CPU) {
		// RES 4,E
	},
	*/
	0xA4: nil,
	/*
	0xA4: func(cpu *CPU) {
		// RES 4,H
	},
	*/
	0xA5: nil,
	/*
	0xA5: func(cpu *CPU) {
		// RES 4,L
	},
	*/
	0xA6: nil,
	/*
	0xA6: func(cpu *CPU) {
		// RES 4,[HL]
	},
	*/
	0xA7: nil,
	/*
	0xA7: func(cpu *CPU) {
		// RES 4,A
	},
	*/
	0xA8: nil,
	/*
	0xA8: func(cpu *CPU) {
		// RES 5,B
	},
	*/
	0xA9: nil,
	/*
	0xA9: func(cpu *CPU) {
		// RES 5,C
	},
	*/
	0xAA: nil,
	/*
	0xAA: func(cpu *CPU) {
		// RES 5,D
	},
	*/
	0xAB: nil,
	/*
	0xAB: func(cpu *CPU) {
		// RES 5,E
	},
	*/
	0xAC: nil,
	/*
	0xAC: func(cpu *CPU) {
		// RES 5,H
	},
	*/
	0xAD: nil,
	/*
	0xAD: func(cpu *CPU) {
		// RES 5,L
	},
	*/
	0xAE: nil,
	/*
	0xAE: func(cpu *CPU) {
		// RES 5,[HL]
	},
	*/
	0xAF: nil,
	/*
	0xAF: func(cpu *CPU) {
		// RES 5,A
	},
	*/
	0xB0: nil,
	/*
	0xB0: func(cpu *CPU) {
		// RES 6,B
	},
	*/
	0xB1: nil,
	/*
	0xB1: func(cpu *CPU) {
		// RES 6,C
	},
	*/
	0xB2: nil,
	/*
	0xB2: func(cpu *CPU) {
		// RES 6,D
	},
	*/
	0xB3: nil,
	/*
	0xB3: func(cpu *CPU) {
		// RES 6,E
	},
	*/
	0xB4: nil,
	/*
	0xB4: func(cpu *CPU) {
		// RES 6,H
	},
	*/
	0xB5: nil,
	/*
	0xB5: func(cpu *CPU) {
		// RES 6,L
	},
	*/
	0xB6: nil,
	/*
	0xB6: func(cpu *CPU) {
		// RES 6,[HL]
	},
	*/
	0xB7: nil,
	/*
	0xB7: func(cpu *CPU) {
		// RES 6,A
	},
	*/
	0xB8: nil,
	/*
	0xB8: func(cpu *CPU) {
		// RES 7,B
	},
	*/
	0xB9: nil,
	/*
	0xB9: func(cpu *CPU) {
		// RES 7,C
	},
	*/
	0xBA: nil,
	/*
	0xBA: func(cpu *CPU) {
		// RES 7,D
	},
	*/
	0xBB: nil,
	/*
	0xBB: func(cpu *CPU) {
		// RES 7,E
	},
	*/
	0xBC: nil,
	/*
	0xBC: func(cpu *CPU) {
		// RES 7,H
	},
	*/
	0xBD: nil,
	/*
	0xBD: func(cpu *CPU) {
		// RES 7,L
	},
	*/
	0xBE: nil,
	/*
	0xBE: func(cpu *CPU) {
		// RES 7,[HL]
	},
	*/
	0xBF: nil,
	/*
	0xBF: func(cpu *CPU) {
		// RES 7,A
	},
	*/
	0xC0: nil,
	/*
	0xC0: func(cpu *CPU) {
		// SET 0,B
	},
	*/
	0xC1: nil,
	/*
	0xC1: func(cpu *CPU) {
		// SET 0,C
	},
	*/
	0xC2: nil,
	/*
	0xC2: func(cpu *CPU) {
		// SET 0,D
	},
	*/
	0xC3: nil,
	/*
	0xC3: func(cpu *CPU) {
		// SET 0,E
	},
	*/
	0xC4: nil,
	/*
	0xC4: func(cpu *CPU) {
		// SET 0,H
	},
	*/
	0xC5: nil,
	/*
	0xC5: func(cpu *CPU) {
		// SET 0,L
	},
	*/
	0xC6: nil,
	/*
	0xC6: func(cpu *CPU) {
		// SET 0,[HL]
	},
	*/
	0xC7: nil,
	/*
	0xC7: func(cpu *CPU) {
		// SET 0,A
	},
	*/
	0xC8: nil,
	/*
	0xC8: func(cpu *CPU) {
		// SET 1,B
	},
	*/
	0xC9: nil,
	/*
	0xC9: func(cpu *CPU) {
		// SET 1,C
	},
	*/
	0xCA: nil,
	/*
	0xCA: func(cpu *CPU) {
		// SET 1,D
	},
	*/
	0xCB: nil,
	/*
	0xCB: func(cpu *CPU) {
		// SET 1,E
	},
	*/
	0xCC: nil,
	/*
	0xCC: func(cpu *CPU) {
		// SET 1,H
	},
	*/
	0xCD: nil,
	/*
	0xCD: func(cpu *CPU) {
		// SET 1,L
	},
	*/
	0xCE: nil,
	/*
	0xCE: func(cpu *CPU) {
		// SET 1,[HL]
	},
	*/
	0xCF: nil,
	/*
	0xCF: func(cpu *CPU) {
		// SET 1,A
	},
	*/
	0xD0: nil,
	/*
	0xD0: func(cpu *CPU) {
		// SET 2,B
	},
	*/
	0xD1: nil,
	/*
	0xD1: func(cpu *CPU) {
		// SET 2,C
	},
	*/
	0xD2: nil,
	/*
	0xD2: func(cpu *CPU) {
		// SET 2,D
	},
	*/
	0xD3: nil,
	/*
	0xD3: func(cpu *CPU) {
		// SET 2,E
	},
	*/
	0xD4: nil,
	/*
	0xD4: func(cpu *CPU) {
		// SET 2,H
	},
	*/
	0xD5: nil,
	/*
	0xD5: func(cpu *CPU) {
		// SET 2,L
	},
	*/
	0xD6: nil,
	/*
	0xD6: func(cpu *CPU) {
		// SET 2,[HL]
	},
	*/
	0xD7: nil,
	/*
	0xD7: func(cpu *CPU) {
		// SET 2,A
	},
	*/
	0xD8: nil,
	/*
	0xD8: func(cpu *CPU) {
		// SET 3,B
	},
	*/
	0xD9: nil,
	/*
	0xD9: func(cpu *CPU) {
		// SET 3,C
	},
	*/
	0xDA: nil,
	/*
	0xDA: func(cpu *CPU) {
		// SET 3,D
	},
	*/
	0xDB: nil,
	/*
	0xDB: func(cpu *CPU) {
		// SET 3,E
	},
	*/
	0xDC: nil,
	/*
	0xDC: func(cpu *CPU) {
		// SET 3,H
	},
	*/
	0xDD: nil,
	/*
	0xDD: func(cpu *CPU) {
		// SET 3,L
	},
	*/
	0xDE: nil,
	/*
	0xDE: func(cpu *CPU) {
		// SET 3,[HL]
	},
	*/
	0xDF: nil,
	/*
	0xDF: func(cpu *CPU) {
		// SET 3,A
	},
	*/
	0xE0: nil,
	/*
	0xE0: func(cpu *CPU) {
		// SET 4,B
	},
	*/
	0xE1: nil,
	/*
	0xE1: func(cpu *CPU) {
		// SET 4,C
	},
	*/
	0xE2: nil,
	/*
	0xE2: func(cpu *CPU) {
		// SET 4,D
	},
	*/
	0xE3: nil,
	/*
	0xE3: func(cpu *CPU) {
		// SET 4,E
	},
	*/
	0xE4: nil,
	/*
	0xE4: func(cpu *CPU) {
		// SET 4,H
	},
	*/
	0xE5: nil,
	/*
	0xE5: func(cpu *CPU) {
		// SET 4,L
	},
	*/
	0xE6: nil,
	/*
	0xE6: func(cpu *CPU) {
		// SET 4,[HL]
	},
	*/
	0xE7: nil,
	/*
	0xE7: func(cpu *CPU) {
		// SET 4,A
	},
	*/
	0xE8: nil,
	/*
	0xE8: func(cpu *CPU) {
		// SET 5,B
	},
	*/
	0xE9: nil,
	/*
	0xE9: func(cpu *CPU) {
		// SET 5,C
	},
	*/
	0xEA: nil,
	/*
	0xEA: func(cpu *CPU) {
		// SET 5,D
	},
	*/
	0xEB: nil,
	/*
	0xEB: func(cpu *CPU) {
		// SET 5,E
	},
	*/
	0xEC: nil,
	/*
	0xEC: func(cpu *CPU) {
		// SET 5,H
	},
	*/
	0xED: nil,
	/*
	0xED: func(cpu *CPU) {
		// SET 5,L
	},
	*/
	0xEE: nil,
	/*
	0xEE: func(cpu *CPU) {
		// SET 5,[HL]
	},
	*/
	0xEF: nil,
	/*
	0xEF: func(cpu *CPU) {
		// SET 5,A
	},
	*/
	0xF0: nil,
	/*
	0xF0: func(cpu *CPU) {
		// SET 6,B
	},
	*/
	0xF1: nil,
	/*
	0xF1: func(cpu *CPU) {
		// SET 6,C
	},
	*/
	0xF2: nil,
	/*
	0xF2: func(cpu *CPU) {
		// SET 6,D
	},
	*/
	0xF3: nil,
	/*
	0xF3: func(cpu *CPU) {
		// SET 6,E
	},
	*/
	0xF4: nil,
	/*
	0xF4: func(cpu *CPU) {
		// SET 6,H
	},
	*/
	0xF5: nil,
	/*
	0xF5: func(cpu *CPU) {
		// SET 6,L
	},
	*/
	0xF6: nil,
	/*
	0xF6: func(cpu *CPU) {
		// SET 6,[HL]
	},
	*/
	0xF7: nil,
	/*
	0xF7: func(cpu *CPU) {
		// SET 6,A
	},
	*/
	0xF8: nil,
	/*
	0xF8: func(cpu *CPU) {
		// SET 7,B
	},
	*/
	0xF9: nil,
	/*
	0xF9: func(cpu *CPU) {
		// SET 7,C
	},
	*/
	0xFA: nil,
	/*
	0xFA: func(cpu *CPU) {
		// SET 7,D
	},
	*/
	0xFB: nil,
	/*
	0xFB: func(cpu *CPU) {
		// SET 7,E
	},
	*/
	0xFC: nil,
	/*
	0xFC: func(cpu *CPU) {
		// SET 7,H
	},
	*/
	0xFD: nil,
	/*
	0xFD: func(cpu *CPU) {
		// SET 7,L
	},
	*/
	0xFE: nil,
	/*
	0xFE: func(cpu *CPU) {
		// SET 7,[HL]
	},
	*/
	0xFF: nil,
	/*
	0xFF: func(cpu *CPU) {
		// SET 7,A
	},
	*/

}

func init() {
	for k := range instructions {
		if instructions[k] == nil {
			instructions[k] = func(cpu *CPU) {
				log.Printf("Unimplemented opcode: %#2x", k)
				// TODO(abhi): replace this once debugger is implemented.
				os.Exit(1)
			}
		}
		if cbInstructions[k] == nil {
			cbInstructions[k] = func(cpu *CPU) {
				log.Printf("[CB Prefix] Unimplemented opcode: %#2x", k)
				// TODO(abhi): replace this once debugger is implemented.
				os.Exit(1)
			}
		}
	}
}