// Handles all interactions between different components (CPU, PPU and MMU).
package gameboy

import (
	"fmt"
	"gopherboy/pkg/common"
	"time"
)

const (
	IE_ADDR = 0xFFFF
	IF_ADDR = 0xFF0F
)


type GB struct {
	CPU *CPU
	MMU *MMU
	// TODO(abhinandj): Add display
	masterClk *time.Ticker
	interruptsEnabled bool
	debug bool
	// TODO: Memory access depends upon the current state (VBLANK, HBLANK etc)
	// We should keep track of it here
}

func NewGB(bootRomPath, cartridgePath string, debug bool) (*GB) {
	mmu := NewMMU(bootRomPath, cartridgePath)
	return &GB{
		CPU: NewCPU(mmu, debug), 
		MMU: mmu, 
		debug: debug,
	}
}

func (gb *GB) RequestInterrupt(idx uint8) {
	existingVal := gb.MMU.ReadAt(IF_ADDR)
	newVal := common.SetBitAtIndex(existingVal, idx)
	gb.MMU.WriteAt(IF_ADDR, newVal)
}

func (gb *GB) resetIFFlag(idx uint8) {
	existingVal := gb.MMU.ReadAt(IF_ADDR)
	newVal := common.ResetBitAtIndex(existingVal, idx)
	gb.MMU.WriteAt(IF_ADDR, newVal)
}

func (gb *GB) executeInterrupt(idx uint8) {
	// Disable further interrupts
	gb.ResetIME()
	gb.resetIFFlag(idx)
	// Start handling the interrupt
	gb.CPU.instrPushSPn16(gb.CPU.PC)
	switch idx {
		case 0: gb.CPU.PC = 0x40
		case 1: gb.CPU.PC = 0x48
		case 2: gb.CPU.PC = 0x50
		case 3: gb.CPU.PC = 0x58
		case 4: gb.CPU.PC = 0x60
	}
}


func (gb *GB) handleInterrupts() int {
	if !gb.interruptsEnabled {
		return 0
	}
	// Handle pending interrupts (if any based) on priority
	requestedInterrupts := gb.MMU.ReadAt(IF_ADDR)
	enabledInterrupts := gb.MMU.ReadAt(IE_ADDR)
	cycles := 0
	for i := uint8(0); i < 5; i++ {
		if common.TestBitAtIndex(enabledInterrupts, i) && common.TestBitAtIndex(requestedInterrupts, i) {
			gb.executeInterrupt(i)
			cycles += 5
		}
	}
	return cycles
}

// Enable all interrupts globally
func (gb *GB) SetIME() {
	gb.interruptsEnabled = true
}

func (gb *GB) ResetIME() {
	gb.interruptsEnabled = false
}

func (gb *GB) Init() error {
	// Set up the clock
	// Initialize other components
	if err := gb.CPU.Init(gb); err != nil {
		return fmt.Errorf("failed to initialize CPU: %v", err)
	}
	gb.MMU.Init(gb)
	gb.masterClk = time.NewTicker(time.Second / time.Duration(common.ClkFrequency))
	gb.interruptsEnabled = false
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
		interruptCycles := gb.handleInterrupts()
		totalCycles = elapsedCycles + interruptCycles
		i += 1
		// TODO: Handle synchronization
		if totalCycles == 0 {
			continue
		}
	}
}