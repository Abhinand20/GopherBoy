package cpu

import (
	"gopherboy/common"
	"gopherboy/mmu"
)

type Cycles common.Cycles

type CPU struct {
	MMU *mmu.MMU
	/* registers */
	a byte
	b byte
	c byte
	d byte
	e byte
	f byte
	h byte
	l byte
	pc uint16
	// Initialized to 0xFFFE, but usually overriden by program
	sp uint16 
}

func (cpu *CPU) Init() error {
	return cpu.MMU.Init()
}

// TODO: Emulate a single CPU tick, return number of instruction cycles elapsed.
func (cpu *CPU) Tick() Cycles { return 0 }

func NewCPU(bootRomPath, cartridgePath string) *CPU {
	return &CPU{MMU: mmu.NewMMU(bootRomPath, cartridgePath)}
}