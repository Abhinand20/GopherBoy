package mmu

import (
	"fmt"
	"os"
)
type MMU struct {
	memory [0xFFFF]byte
	bootRomPath string
	cartridgePath string
}

func (mmu *MMU) Init() error {
	boot, err := os.ReadFile(mmu.bootRomPath)
	if err != nil {
		return fmt.Errorf("could not read the boot rom, %v", err)
	}
	_, err = os.ReadFile(mmu.cartridgePath)
	if err != nil {
		return fmt.Errorf("could not read the cartridge, %v", err)
	}
	n := copy(mmu.memory[:], boot)
	fmt.Printf("Copied boot rom into memory: %d bytes\n", n)
	// TODO(abhinandj): Store the cartridge ROM into memory.
	return nil
}


func NewMMU(bootRomPath, cartridgePath string) *MMU {
	return &MMU{
		bootRomPath: bootRomPath,
		cartridgePath: cartridgePath,
	}
}