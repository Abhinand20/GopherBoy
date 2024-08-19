// Handles all interactions between different components (CPU, PPU and MMU).
package motherboard

import (
	"fmt"
	"gopherboy/cpu"
)

type Motherboard struct {
	CPU *cpu.CPU
	// TODO(abhinandj): Add display
}

func NewMotherboard(bootRomPath, cartridgePath string) (*Motherboard) {
	return &Motherboard{CPU: cpu.NewCPU(bootRomPath, cartridgePath)}
}

func (mb *Motherboard) Emulate() error {
	fmt.Println("Started emulation...")
	mb.CPU.Init()
	return nil
}