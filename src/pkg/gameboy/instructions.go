package gameboy

import (
	"fmt"
	"gopherboy/pkg/common"
	"log"
	"os"
	"strings"
)

// OpcodeCycles is the number of cpu cycles for each normal opcode.
var OpcodeCycles = []int{
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

type InstrInfo struct {
	opcode byte
	addr uint16
	instr string
	leftOp string
	rightOp string
	instrType string
}

func NewInstrInfo(opcode byte, addr uint16, instr string) *InstrInfo {
	instrType, args, found := strings.Cut(instr, " ")
	instrInfo := &InstrInfo{
		opcode: opcode,
		addr: addr,
		instr: instr,
		instrType: instrType,
	}
	if found {
		leftOp, rightOp, _ := strings.Cut(args, ",")
		instrInfo.leftOp = leftOp
		instrInfo.rightOp = rightOp
	}
	return instrInfo
}

func (ii *InstrInfo) updateInstrPlaceholder(op string, cpu *CPU) string {
	newInstr := op
	if strings.Contains(op, "a16") || strings.Contains(op, "n16") {
		b1 := uint16(cpu.MMU.ReadAt(cpu.PC))
		b2 := uint16(cpu.MMU.ReadAt(cpu.PC + 1))
		addr := b2 << 8 | b1
		newInstr = strings.ReplaceAll(op, "a16", fmt.Sprintf("a16 {%#4x}", addr))
		newInstr = strings.ReplaceAll(newInstr, "n16", fmt.Sprintf("n16 {%#4x}", addr))
	}
	if strings.Contains(op, "n8") || strings.Contains(op, "a8") {
		addr := cpu.MMU.ReadAt(cpu.PC)
		newInstr = strings.ReplaceAll(op, "n8", fmt.Sprintf("n8 {%#2x}", addr))
		newInstr = strings.ReplaceAll(newInstr, "a8", fmt.Sprintf("a8 {$FF00 + %#2x}", addr))
	}
	if strings.Contains(op, "e8") {
		addr := int8(cpu.MMU.ReadAt(cpu.PC))
		newInstr = strings.ReplaceAll(newInstr, "e8", fmt.Sprintf("e8 {%d}", addr))
	}
	return newInstr
}

func (ii *InstrInfo) DebugInfo(cpuState *CPU) string {
	ii.leftOp = ii.updateInstrPlaceholder(ii.leftOp, cpuState)
	ii.rightOp = ii.updateInstrPlaceholder(ii.rightOp, cpuState)
	instrStr := ii.instrType + " " + ii.leftOp
	if len(ii.rightOp) > 0 {
		instrStr += "," + ii.rightOp
	}
	return fmt.Sprintf("%#4x %#2x\t%s\n", ii.addr, ii.opcode, instrStr)
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

func (cpu *CPU) instrDECr16(r *Register) {
	val := r.Value()
	r.Set(val - 1)
}

func (cpu *CPU) instrINCr16(r *Register) {
	val := r.Value()
	r.Set(val + 1)
}

// Increment register and update flags
func (cpu *CPU) instrINCr8(setFunc func(byte), val byte) {
	incVal := val + 1
	setFunc(incVal)
	cpu.setZ(incVal == 0)
	cpu.setN(false)
	cpu.setH(common.IsHalfCarry(val, 1))
}

// Decrement register and update flags
func (cpu *CPU) instrDECr8(setFunc func(byte), val byte) {
	decVal := val - 1
	setFunc(decVal)
	cpu.setZ(decVal == 0)
	cpu.setN(true)
	cpu.setH(common.IsHalfBorrow(val, 1))
}

func (cpu *CPU) instrLDn8(dest uint16, src byte) {
	cpu.MMU.WriteAt(dest, src)
}

// Common function for LD r8,r8
func (cpu *CPU) instrLDr8(setFunc func(byte), val byte) {
	setFunc(val)
}

// Implement CP A,r8
func (cpu *CPU) instrCPr8(val byte) {
	a := cpu.AF.Hi()
	
	cpu.setZ(a == val)
	cpu.setN(true)
	cpu.setH(common.IsHalfBorrow(a, val))
	cpu.setC(val > a)
}

func (cpu *CPU) instrJR(offset byte) {
	signedOffset := int8(offset)
	currAddr := int16(cpu.PC)
	cpu.PC = uint16(currAddr + int16(signedOffset))
}

// Push an 8-bit value onto the stack
func (cpu *CPU) instrPushSPn8(val byte) {
	cpu.instrDECr16(&cpu.SP)
	cpu.instrLDn8(cpu.SP.Value(), val)
}

// Push a 16-bit register onto the stack
func (cpu *CPU) instrPushSPr16(r *Register) {
	cpu.instrPushSPn8(r.Hi())
	cpu.instrPushSPn8(r.Lo())
}

// Pop a byte from stack
func (cpu *CPU) instrPopSPn8() byte {
	defer cpu.instrINCr16(&cpu.SP)
	return cpu.MMU.ReadAt(cpu.SP.Value())
}

// Push a 16-bit value onto the stack
func (cpu *CPU) instrPushSPn16(val uint16) {
	upperByte := byte(val >> 8)
	lowerByte := byte(val & 0xFF)
	cpu.instrPushSPn8(upperByte)
	cpu.instrPushSPn8(lowerByte)
}

func (cpu *CPU) instrPopSPr16(r *Register) {
	r.SetLo(cpu.instrPopSPn8())
	r.SetHi(cpu.instrPopSPn8())
}

func (cpu *CPU) instrPopSPr16PC() {
	valLo := cpu.instrPopSPn8()
	valHi := cpu.instrPopSPn8()
	cpu.PC = uint16(valHi) << 8 | uint16(valLo)
}

/* opcodes to function mappings */
var instructions = [0x100]func(cpu *CPU) {
	/* LD */
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
	/* JR */
	0x18: func(cpu *CPU) {
		// JR e8
		cpu.instrJR(cpu.popPC8())
	},
	0x20: func(cpu *CPU) {
		// JR NZ, e8
		if !cpu.testZ() {
			cpu.instrJR(cpu.popPC8())
			return
		}
		cpu.popPC8()
	},
	0x30: func(cpu *CPU) {
		// JR NC, e8
		if !cpu.testC() {
			cpu.instrJR(cpu.popPC8())
			return
		}
		cpu.popPC8()
	},
	0x28: func(cpu *CPU) {
		// JR Z, e8
		if cpu.testZ() {
			cpu.instrJR(cpu.popPC8())
			return
		}
		cpu.popPC8()
	},
	0x38: func(cpu *CPU) {
		// JR C, e8
		if cpu.testC() {
			cpu.instrJR(cpu.popPC8())
			return
		}
		cpu.popPC8()
	},
	0x32: func(cpu *CPU) {
		// LD [HLD], A
		cpu.instrLDn8(cpu.HL.Value(), cpu.AF.Hi())
		cpu.instrDECr16(&cpu.HL)
	},
	0x0E: func(cpu *CPU) {
		// LD C, n8
		val := cpu.popPC8()
		cpu.BC.SetLo(val)
	},
	0x1E: func(cpu *CPU) {
		// LD E, n8
		val := cpu.popPC8()
		cpu.DE.SetLo(val)
	},
	0x2E: func(cpu *CPU) {
		// LD L, n8
		val := cpu.popPC8()
		cpu.HL.SetLo(val)
	},
	0x3E: func(cpu *CPU) {
		// LD A, n8
		val := cpu.popPC8()
		cpu.AF.SetHi(val)
	},
	/* LD r8,n8 */
	0x06: func(cpu *CPU) {
		// LD B, n8
		val := cpu.popPC8()
		cpu.BC.SetHi(val)
	},
	/* LD to HRAM */
	0xE0: func(cpu *CPU) {
		// LD [0xFFOO + n8], A
		destAddr := 0xFF00 + uint16(cpu.popPC8())
		cpu.instrLDn8(destAddr, cpu.AF.Hi())
	},
	0xE2: func(cpu *CPU) {
		// LD [0xFFOO + C], A
		destAddr := 0xFF00 + uint16(cpu.BC.Lo())
		cpu.instrLDn8(destAddr, cpu.AF.Hi())
	},
	0xF0: func(cpu *CPU) {
		// LD A, [0xFFOO + n8]
		srcAddr := 0xFF00 + uint16(cpu.popPC8())
		cpu.AF.SetHi(cpu.MMU.ReadAt(srcAddr))
	},
	0xF2: func(cpu *CPU) {
		// LD A, [0xFFOO + C]
		srcAddr := 0xFF00 + uint16(cpu.BC.Lo())
		cpu.AF.SetHi(cpu.MMU.ReadAt(srcAddr))
	},
	/* Interrupt control (Ideally these should be delayed by 1 instruction) */
	0xF3: func(cpu *CPU) {
		// DI
		cpu.gb.ResetIME()
	},
	0xFB: func(cpu *CPU) {
		// EI
		cpu.gb.SetIME()
	},
	0xD9: func(cpu *CPU) {
		// RETI
		cpu.gb.SetIME()
		cpu.instrPopSPr16PC()
	},
	/* CP n8 */
	0xFE: func(cpu *CPU) {
		// CP n8
		val := cpu.popPC8()
		cpu.instrCPr8(val)
	},
	/* LD A, [r16] */
	0x0A: func(cpu *CPU) {
		// LD A, [BC]
		srcAddr := cpu.BC.Value()
		cpu.AF.SetHi(cpu.MMU.ReadAt(srcAddr))
	},
	0x1A: func(cpu *CPU) {
		// LD A, [DE]
		srcAddr := cpu.DE.Value()
		cpu.AF.SetHi(cpu.MMU.ReadAt(srcAddr))
	},
	0x2A: func(cpu *CPU) {
		// LD A, [HL+]
		srcAddr := cpu.HL.Value()
		cpu.AF.SetHi(cpu.MMU.ReadAt(srcAddr))
		cpu.instrINCr16(&cpu.HL)
	},
	0x3A: func(cpu *CPU) {
		// LD A, [HL-]
		srcAddr := cpu.DE.Value()
		cpu.AF.SetHi(cpu.MMU.ReadAt(srcAddr))
		cpu.instrDECr16(&cpu.HL)
	},
	/* Jumps and Sub-routines */
	0xCD: func(cpu *CPU) {
		// CALL a16
		addrJP := cpu.popPC16()
		cpu.instrPushSPn16(cpu.PC)
		cpu.PC = addrJP
	},
	/* Stack operations */
	0xC5: func(cpu *CPU) {
		// PUSH BC
		cpu.instrPushSPr16(&cpu.BC)
	},
	0xC1: func(cpu *CPU) {
		// POP BC
		cpu.instrPopSPr16(&cpu.BC)
	},
	0xD5: func(cpu *CPU) {
		// PUSH DE
		cpu.instrPushSPr16(&cpu.DE)
	},
	0xD1: func(cpu *CPU) {
		// POP DE
		cpu.instrPopSPr16(&cpu.DE)
	},
	0xE5: func(cpu *CPU) {
		// PUSH HL
		cpu.instrPushSPr16(&cpu.HL)
	},
	0xEA: func(cpu *CPU) {
		// LD [a16], A
		cpu.instrLDn8(cpu.popPC16(), cpu.AF.Hi())

	},
	0xFA: func(cpu *CPU) {
		// LD A, [a16]
		srcAddr := cpu.popPC16()
		cpu.AF.SetHi(cpu.MMU.ReadAt(srcAddr))
	},
	0xE1: func(cpu *CPU) {
		// POP HL
		cpu.instrPopSPr16(&cpu.HL)
	},
	// 0xF5: func(cpu *CPU) {
	// 	// PUSH AF
	// },
	// 0xF1: func(cpu *CPU) {
	// 	// POP AF
	// },
	/* INC */
	0x04: func(cpu *CPU) {
		// INC B
		cpu.instrINCr8(cpu.BC.SetHi, cpu.BC.Hi())
	},
	0x14: func(cpu *CPU) {
		// INC D
		cpu.instrINCr8(cpu.DE.SetHi, cpu.DE.Hi())
	},
	0x24: func(cpu *CPU) {
		// INC H
		cpu.instrINCr8(cpu.HL.SetHi, cpu.HL.Hi())
	},
	0x34: func(cpu *CPU) {
		// INC [HL]
		val := cpu.MMU.ReadAt(cpu.HL.Value())
		newVal := val + 1
		cpu.MMU.WriteAt(cpu.HL.Value(), newVal)

		cpu.setZ(newVal == 0)
		cpu.setN(false)
		cpu.setH(common.IsHalfCarry(val, 1))
	},
	0x0C: func(cpu *CPU) {
		// INC C
		cpu.instrINCr8(cpu.BC.SetLo, cpu.BC.Lo())
	},
	0x1C: func(cpu *CPU) {
		// INC E
		cpu.instrINCr8(cpu.DE.SetLo, cpu.DE.Lo())
	},
	0x2C: func(cpu *CPU) {
		// INC L
		cpu.instrINCr8(cpu.HL.SetLo, cpu.HL.Lo())
	},
	0x3C: func(cpu *CPU) {
		// INC A
		cpu.instrINCr8(cpu.AF.SetHi, cpu.AF.Hi())
	},
	0x03: func(cpu *CPU) {
		// INC BC
		cpu.instrINCr16(&cpu.BC)
	},
	0x13: func(cpu *CPU) {
		// INC DE
		cpu.instrINCr16(&cpu.DE)
	},
	0x23: func(cpu *CPU) {
		// INC HL
		cpu.instrINCr16(&cpu.HL)
	},
	0x33: func(cpu *CPU) {
		// INC SP
		cpu.instrINCr16(&cpu.SP)
	},
	/* DEC */
	0x05: func(cpu *CPU) {
		// DEC B
		cpu.instrDECr8(cpu.BC.SetHi, cpu.BC.Hi())
	},
	0x15: func(cpu *CPU) {
		// DEC D
		cpu.instrDECr8(cpu.DE.SetHi, cpu.DE.Hi())
	},
	0x25: func(cpu *CPU) {
		// DEC H
		cpu.instrDECr8(cpu.HL.SetHi, cpu.HL.Hi())
	},
	0x35: func(cpu *CPU) {
		// DEC [HL]
		val := cpu.MMU.ReadAt(cpu.HL.Value())
		newVal := val - 1
		cpu.MMU.WriteAt(cpu.HL.Value(), newVal)

		cpu.setZ(newVal == 0)
		cpu.setN(true)
		cpu.setH(common.IsHalfBorrow(val, 1))
	},
	0x0D: func(cpu *CPU) {
		// DEC C
		cpu.instrDECr8(cpu.BC.SetLo, cpu.BC.Lo())
	},
	0x1D: func(cpu *CPU) {
		// DEC E
		cpu.instrDECr8(cpu.DE.SetLo, cpu.DE.Lo())
	},
	0x2D: func(cpu *CPU) {
		// DEC L
		cpu.instrDECr8(cpu.HL.SetLo, cpu.DE.Lo())
	},
	0x3D: func(cpu *CPU) {
		// DEC A
		cpu.instrDECr8(cpu.AF.SetHi, cpu.AF.Hi())
	},
	0x77: func(cpu *CPU) {
		// LD [HL], A
		cpu.instrLDn8(cpu.HL.Value(), cpu.AF.Hi())
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
	/* BIT shift */
	0x07: func(cpu *CPU) {
		// RLA
		cpu.instrRL(cpu.AF.SetHi, cpu.AF.Hi())
		cpu.setZ(false)
	},
	0x17: func(cpu *CPU) {
		// RLCA
		cpu.instrRLC(cpu.AF.SetHi, cpu.AF.Hi())
		cpu.setZ(false)
	},
	0xC9: func(cpu *CPU) {
		// RET
		cpu.instrPopSPr16PC()
	},
}


func init() {
	// Populate instructions from 0x40 - 0xBF
	for i := 0; i < 8; i++ {
		i := i
		// LD r8,r8
		instructions[0x40 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrLDr8(cpu.BC.SetHi, params.val)
		}
		instructions[0x48 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrLDr8(cpu.BC.SetLo, params.val)
		}
		instructions[0x50 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrLDr8(cpu.DE.SetHi, params.val)
		}
		instructions[0x58 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrLDr8(cpu.DE.SetLo, params.val)
		}
		instructions[0x60 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrLDr8(cpu.HL.SetHi, params.val)
		}
		instructions[0x68 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrLDr8(cpu.HL.SetLo, params.val)
		}
		// Skip HALT instruction
		if i != 6 {
			instructions[0x70 + i] = func(cpu *CPU) {
				params := buildCbInstrParams(cpu, i)
				cpu.instrLDr8(func (v byte) {cpu.MMU.WriteAt(cpu.HL.Value(), v)}, params.val)
			}
		}
		instructions[0x78 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrLDr8(cpu.AF.SetHi, params.val)
		}
		// CP A,r8
		instructions[0xB8 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrCPr8(params.val)
		}
	}
	for k := range instructions {
		if instructions[k] == nil {
			instructions[k] = func(cpu *CPU) {
				log.Printf("Unimplemented opcode: %#2x", k)
				// TODO(abhi): replace this once debugger is implemented.
				os.Exit(1)
			}
		}
	}
}