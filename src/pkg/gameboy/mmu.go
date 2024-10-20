package gameboy

import (
	"fmt"
	"os"
)
type MMU struct {
	memory [0xFFFF]byte
	// 16KiB ROM0
	bank0 [0x4000]byte
	// 16KiB ROM-N
	bankN [0x4000]byte
	// 8KiB VRAM
	vram [0x2000]byte
	// 8 KiB ERAM
	eram [0x2000]byte
	// 8 KiB WRAM
	wram [0x2000]byte
	// 160 bytes OAM 
	oam [0xA0]byte
	// 128 Byte High RAM
	hram [0x80]byte

	biosEnabled bool
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
	return nil
}

func (mmu *MMU) ReadAt(addr uint16) byte {
	return mmu.memory[addr]
}

func (mmu *MMU) WriteAt(addr uint16, val byte) {
	mmu.memory[addr] = val
}


func NewMMU(bootRomPath, cartridgePath string) *MMU {
	return &MMU{
		bootRomPath: bootRomPath,
		cartridgePath: cartridgePath,
	}
}