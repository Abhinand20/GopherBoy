package cpu

import (
	"gopherboy/common"
	"gopherboy/mmu"
)

type Cycles common.Cycles

type Register struct {
	value uint16
	// Used to reset the 'F' register after operations.
	mask uint16
}

func (r *Register) Lo() byte {
	return byte(r.value & 0xFF)
}

func (r *Register) Hi() byte {
	return byte(r.value >> 8)
}

func (r *Register) SetLo(b byte) {
	r.value = uint16(b) | (r.value & 0xFF00)
	r.applyMask()
}
func (r *Register) SetHi(b byte) {
	r.value = (r.value & 0x00FF) | (uint16(b) << 8)
	r.applyMask()
}

func (r Register) Set(v uint16) {
	r.value = v
	r.applyMask()
}

func (r *Register) applyMask() {
	if r.mask != 0 {
		r.value &= r.mask
	}
}

type CPU struct {
	MMU *mmu.MMU
	/* registers */
	AF Register
	BC Register
	DE Register
	HL Register

	PC uint16
	// Initialized to 0xFFFE, but usually overriden by program
	SP Register
}

// Read the following byte from PC and advance the pointer.
func (cpu *CPU) popPC8() byte {
	val := cpu.MMU.ReadAt(cpu.PC)
	cpu.PC += 1
	return val
}

// Read the following 16bits from PC and advance the pointer.
func (cpu *CPU) popPC16() uint16 {
	val := uint16(cpu.popPC8())
	val <<= 8
	val |= uint16(cpu.popPC8())
	return val
}


func (cpu *CPU) Init() error {
	// TODO(abhinandj): Update the mapping from opcode to instructions
	cpu.PC = 0
	cpu.AF.mask = 0xFFF0
	return cpu.MMU.Init()
}

// TODO: Emulate a single CPU tick, return number of instruction cycles elapsed.
func (cpu *CPU) Tick() Cycles {
	opcode := cpu.popPC8()
	instructions[opcode](cpu)
	return OpcodeCycles[opcode]
}


func NewCPU(bootRomPath, cartridgePath string) *CPU {
	return &CPU{MMU: mmu.NewMMU(bootRomPath, cartridgePath)}
}