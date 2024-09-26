// Handles all interactions between different components (CPU, PPU and MMU).
package motherboard

import (
	"fmt"
	"gopherboy/common"
	"gopherboy/cpu"
	"time"
)

type Motherboard struct {
	CPU *cpu.CPU
	// TODO(abhinandj): Add display
	masterClk *time.Ticker
	debug bool
}

func NewMotherboard(bootRomPath, cartridgePath string, debug bool) (*Motherboard) {
	return &Motherboard{CPU: cpu.NewCPU(bootRomPath, cartridgePath, debug), debug: debug}
}

func (mb *Motherboard) Init() error {
	// Set up the clock
	// Initialize other components
	if err := mb.CPU.Init(); err != nil {
		return fmt.Errorf("failed to initialize CPU: %v", err)
	}
	mb.masterClk = time.NewTicker(time.Second / time.Duration(common.ClkFrequency))
	return nil
}

// Main emulation loop
func (mb *Motherboard) Emulate() error {
	fmt.Println("Started emulation...")
	totalCycles := 0
	i := 0
	stop := 10
	for {
		<- mb.masterClk.C
		elapsedCycles := mb.CPU.Tick()
		totalCycles += int(elapsedCycles)
		i += 1
		if i == stop {
			return nil
		}
	}
}