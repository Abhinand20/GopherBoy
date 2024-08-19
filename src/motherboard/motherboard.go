// Handles all interactions between different components (CPU, PPU and MMU).
package motherboard

import (
	"fmt"
	"gopherboy/cpu"
	"os"
)

type Motherboard struct {
	CPU *cpu.CPU
	bootRom []byte
	cartridge []byte
}

func NewMotherboard(bootRomPath, cartridgePath string) (*Motherboard, error) {
	boot, err := os.ReadFile(bootRomPath)
	if err != nil {
		return nil, fmt.Errorf("could not read the boot rom, %v", err)
	}
	cartridge , err := os.ReadFile(cartridgePath)
	if err != nil {
		return nil, fmt.Errorf("could not read the cartridge, %v", err)
	}
	mb := Motherboard{
		CPU: cpu.NewCPU(),
		bootRom: boot,
		cartridge: cartridge,
	}
	return &mb, nil
}

func (mb *Motherboard) Emulate() error {
	fmt.Println("Started emulation...")
	mb.CPU.Init()
	return nil
}