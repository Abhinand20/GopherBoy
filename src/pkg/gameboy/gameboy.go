// Handles all interactions between different components (CPU, PPU and MMU).
package gameboy

import (
	"fmt"
	"gopherboy/pkg/common"
	"time"
)

type GB struct {
	CPU *CPU
	MMU *MMU
	// TODO(abhinandj): Add display
	masterClk *time.Ticker
	debug bool
}

func NewGB(bootRomPath, cartridgePath string, debug bool) (*GB) {
	mmu := NewMMU(bootRomPath, cartridgePath)
	return &GB{
		CPU: NewCPU(mmu, debug), 
		MMU: mmu, 
		debug: debug,
	}
}

func (gb *GB) Init() error {
	// Set up the clock
	// Initialize other components
	if err := gb.CPU.Init(); err != nil {
		return fmt.Errorf("failed to initialize CPU: %v", err)
	}
	gb.MMU.Init(gb)
	gb.masterClk = time.NewTicker(time.Second / time.Duration(common.ClkFrequency))
	return nil
}

// Main emulation loop
func (gb *GB) Emulate() error {
	fmt.Println("Started emulation...")
	totalCycles := 0
	i := 0
	for {
		<- gb.masterClk.C
		elapsedCycles := gb.CPU.Tick()
		totalCycles += elapsedCycles
		i += 1
	}
}