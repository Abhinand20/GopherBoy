package cpu

import "gopherboy/common"

// CBOpcodeCycles is the number of cpu cycles for each CB opcode.
var CBOpcodeCycles = []int{
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

type cbInstrParams struct {
	setter func(byte)
	val byte
}

var cbInstructions [0x100]func(cpu *CPU)

func buildCbInstrParams(cpu *CPU, i int) cbInstrParams {
	index := [8]cbInstrParams{
		// B
		0: {setter: cpu.BC.SetHi, val: cpu.BC.Hi()},
		// C
		1: {setter: cpu.BC.SetLo, val: cpu.BC.Lo()},
		// D
		2: {setter: cpu.DE.SetHi, val: cpu.DE.Hi()},
		// E
		3: {setter: cpu.DE.SetLo, val: cpu.DE.Lo()},
		// H
		4: {setter: cpu.HL.SetHi, val: cpu.HL.Hi()},
		// L
		5: {setter: cpu.HL.SetLo, val: cpu.HL.Lo()},
		// [HL]
		6: {setter: func(v byte) {cpu.MMU.WriteAt(cpu.HL.Value(), v)}, val: cpu.MMU.ReadAt(cpu.HL.Value())},
		// A
		7: {setter: cpu.AF.SetHi, val: cpu.AF.Hi()},
	}
	return index[i]
}

func (cpu *CPU) instrRLC(setFunc func(byte), val byte) {
	c := (val >> 7) & 0x1
	newVal := ((val << 1) & 0xFF) | c
	setFunc(newVal)

	cpu.setZ(newVal == 0)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setC(c == 1)
}

func (cpu *CPU) instrRL(setFunc func(byte), val byte) {
	oldC := common.BoolToByte(cpu.testC())
	newC := (val >> 7) & 0x1
	newVal := ((val << 1) & 0xFF) | oldC
	setFunc(newVal)

	cpu.setZ(newVal == 0)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setC(newC == 1)
}

func (cpu *CPU) instrSLA(setFunc func(byte), val byte) {
	c := (val >> 7) & 0x1
	newVal := (val << 1) & 0xFF 
	setFunc(newVal)

	cpu.setZ(newVal == 0)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setC(c == 1)
}

func (cpu *CPU) instrRRC(setFunc func(byte), val byte) {
	c := val & 0x1
	newVal := ((val >> 1) & 0xFF) |	(c << 7)
	setFunc(newVal)

	cpu.setZ(newVal == 0)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setC(c == 1)
}

func (cpu *CPU) instrRR(setFunc func(byte), val byte) {
	oldC := common.BoolToByte(cpu.testC())
	newC := val & 0x1
	newVal := ((val >> 1) & 0xFF) | (oldC << 7)
	setFunc(newVal)

	cpu.setZ(newVal == 0)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setC(newC == 1)
}

func (cpu *CPU) instrSRA(setFunc func(byte), val byte) {
	c := val & 0x1
	newVal := (val & (1 << 7)) | (val >> 1)
	setFunc(newVal)

	cpu.setZ(newVal == 0)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setC(c == 1)
}

func (cpu *CPU) instrSRL(setFunc func(byte), val byte) {
	c := val & 0x1
	newVal := (val >> 1) & 0xFF
	setFunc(newVal)

	cpu.setZ(newVal == 0)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setC(c == 1)
}

func (cpu *CPU) instrSWAP(setFunc func(byte), val byte) {
	swappedVal := (val >> 4) & 0xF | (val << 4) & 0xF0
	setFunc(swappedVal)

	cpu.setZ(swappedVal == 0)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setC(false)
}

func (cpu *CPU) instrBIT(val byte, i byte) {
	cpu.setZ(!common.TestBitAtIndex(val, i))
	cpu.setN(false)
	cpu.setH(true)
}


func (cpu *CPU) instrRES(setFunc func(byte), val, i byte) {
	newVal := common.ResetBitAtIndex(val, i)
	setFunc(newVal)
}

func (cpu *CPU) instrSET(setFunc func(byte), val, i byte) {
	newVal := common.SetBitAtIndex(val, i)
	setFunc(newVal)
}

func init() {
	// Octal format - B,C,D,E,H,L,[HL],A
	for i := 0; i < 8; i++ {
		i := i
		// RLC
		cbInstructions[0x00 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRLC(params.setter, params.val)
		}
		// RRC
		cbInstructions[0x08 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRRC(params.setter, params.val)
		}
		// RL
		cbInstructions[0x10 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRL(params.setter, params.val)
		}
		// RR
		cbInstructions[0x18 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRR(params.setter, params.val)
		}
		// SLA
		cbInstructions[0x20 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSLA(params.setter, params.val)
		}
		// SRA
		cbInstructions[0x28 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSRA(params.setter, params.val)
		}
		// SWAP
		cbInstructions[0x30 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSWAP(params.setter, params.val)
		}
		// SRL
		cbInstructions[0x38 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSRL(params.setter, params.val)
		}
		
		// BIT
		cbInstructions[0x40 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrBIT(params.val, 0)
		}
		cbInstructions[0x48 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrBIT(params.val, 1)
		}
		cbInstructions[0x50 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrBIT(params.val, 2)
		}
		cbInstructions[0x58 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrBIT(params.val, 3)
		}
		cbInstructions[0x60 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrBIT(params.val, 4)
		}
		cbInstructions[0x68 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrBIT(params.val, 5)
		}
		cbInstructions[0x70 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrBIT(params.val, 6)
		}
		cbInstructions[0x78 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrBIT(params.val, 7)
		}

		// RES
		cbInstructions[0x80 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRES(params.setter, params.val, 0)
		}
		cbInstructions[0x88 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRES(params.setter, params.val, 1)
		}
		cbInstructions[0x90 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRES(params.setter, params.val, 2)
		}
		cbInstructions[0x98 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRES(params.setter, params.val, 3)
		}
		cbInstructions[0xA0 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRES(params.setter, params.val, 4)
		}
		cbInstructions[0xA8 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRES(params.setter, params.val, 5)
		}
		cbInstructions[0xB0 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRES(params.setter, params.val, 6)
		}
		cbInstructions[0xB8 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrRES(params.setter, params.val, 7)
		}

		// SET
		cbInstructions[0xC0 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSET(params.setter, params.val, 0)
		}
		cbInstructions[0xC8 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSET(params.setter, params.val, 1)
		}
		cbInstructions[0xD0 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSET(params.setter, params.val, 2)
		}
		cbInstructions[0xD8 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSET(params.setter, params.val, 3)
		}
		cbInstructions[0xE0 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSET(params.setter, params.val, 4)
		}
		cbInstructions[0xE8 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSET(params.setter, params.val, 5)
		}
		cbInstructions[0xF0 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSET(params.setter, params.val, 6)
		}
		cbInstructions[0xF8 + i] = func(cpu *CPU) {
			params := buildCbInstrParams(cpu, i)
			cpu.instrSET(params.setter, params.val, 7)
		}
	}
}