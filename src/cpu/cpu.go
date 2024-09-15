package cpu

import (
	"gopherboy/common"
	"gopherboy/mmu"
	"log"
)

type Cycles common.Cycles

// Flag bit indexes within the F register.
const (
	Z_IDX uint8 = 7
	N_IDX uint8 = 6
	H_IDX uint8 = 5
	C_IDX uint8 = 4
)


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

func (r *Register) Value() uint16 {
	return r.value
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


func (cpu *CPU) setFlag(index uint8, on bool) {
	flags := cpu.AF.Lo()
	if on {
		cpu.AF.SetLo(common.SetBitAtIndex(flags, index))
		return
	}
	cpu.AF.SetLo(common.ResetBitAtIndex(flags, index))
}

func (cpu *CPU) setZ(on bool) {
	cpu.setFlag(Z_IDX, on)
}

func (cpu *CPU) setN(on bool) {
	cpu.setFlag(N_IDX, on)
}

func (cpu *CPU) setH(on bool) {
	cpu.setFlag(H_IDX, on)
}
func (cpu *CPU) setC(on bool) {
	cpu.setFlag(C_IDX, on)
}

func (cpu *CPU) Init() error {
	// TODO(abhinandj): Update the mapping from opcode to instructions
	cpu.PC = 0
	// Only the upper 4 bits of the F register are in-use.
	cpu.AF.mask = 0xFFF0
	return cpu.MMU.Init()
}

// TODO: Emulate a single CPU tick, return number of instruction cycles elapsed.
func (cpu *CPU) Tick() Cycles {
	addr := cpu.PC
	opcode := cpu.popPC8()
	opcodeStr := common.InstrDebugLookup[opcode]
	if opcode != 0xCB {
		log.Printf("%#4x %#2x\t%s\n", addr, opcode, opcodeStr)
		instructions[opcode](cpu)
		return OpcodeCycles[opcode] * 4
	}
	addr = cpu.PC
	opcode = cpu.popPC8()
	opcodeStr = common.PrefixInstrDebugLookup[opcode]
	log.Printf("%#4x %#2x\t%s\n", addr, opcode, opcodeStr)
	cbInstructions[opcode](cpu)
	return CBOpcodeCycles[opcode] * 4
}


func NewCPU(bootRomPath, cartridgePath string) *CPU {
	return &CPU{MMU: mmu.NewMMU(bootRomPath, cartridgePath)}
}