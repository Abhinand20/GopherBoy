package cpu

import (
	"fmt"
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

func (r *Register) Set(v uint16) {
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
	debug bool
}

// Read the following byte from PC and advance the pointer.
func (cpu *CPU) popPC8() byte {
	val := cpu.MMU.ReadAt(cpu.PC)
	cpu.PC += 1
	return val
}

// Read the following 16bits from PC and advance the pointer.
func (cpu *CPU) popPC16() uint16 {
	b1 := uint16(cpu.popPC8())
	b2 := uint16(cpu.popPC8())
	return b2 << 8 | b1
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

func (cpu *CPU) testZ() bool {
	return common.TestBitAtIndex(cpu.AF.Lo(), Z_IDX)
}
func (cpu *CPU) testC() bool {
	return common.TestBitAtIndex(cpu.AF.Lo(), C_IDX)
}
func (cpu *CPU) testN() bool {
	return common.TestBitAtIndex(cpu.AF.Lo(), N_IDX)
}
func (cpu *CPU) testH() bool {
	return common.TestBitAtIndex(cpu.AF.Lo(), H_IDX)
}

func (cpu *CPU) printRegisterDump() {
	out := fmt.Sprintf(
		"B: %#4x\tC: %#4x\nD: %#4x\tE: %#4x\nH: %#4x\tL: %#4x\nA: %#4x\n",
		cpu.BC.Hi(),
		cpu.BC.Lo(),
		cpu.DE.Hi(),
		cpu.DE.Lo(),
		cpu.HL.Hi(),
		cpu.HL.Lo(),
		cpu.AF.Hi(),
	)
	reserved := fmt.Sprintf(
		"PC: %#4x\tSP: %x\n",
		cpu.PC,
		cpu.SP.value,
	)
	flags := fmt.Sprintf(
		"Z: %v\tC: %v\tN: %v\tH :%v\n",
		cpu.testZ(),
		cpu.testC(),
		cpu.testN(),
		cpu.testH(),
	)
	fmt.Printf("[REGISTERS]\n%s%s[FLAGS]\n%s\n",out, reserved, flags)
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
	instructionMapping := instructions
	opcodeCyclesMapping := OpcodeCycles
	if opcode == 0xCB {
		addr = cpu.PC
		opcode = cpu.popPC8()
		opcodeStr = common.PrefixInstrDebugLookup[opcode]
		instructionMapping = cbInstructions
		opcodeCyclesMapping = CBOpcodeCycles
	}
	log.Printf("%#4x %#2x\t%s\n", addr, opcode, opcodeStr)
	instructionMapping[opcode](cpu)
	cycles := opcodeCyclesMapping[opcode] * 4
	if cpu.debug {
		cpu.printRegisterDump()
	}
	return cycles
}


func NewCPU(bootRomPath, cartridgePath string, debug bool) *CPU {
	return &CPU{MMU: mmu.NewMMU(bootRomPath, cartridgePath), debug: debug}
}