package gameboy

import (
	"fmt"
	"os"
)

type MMU struct {
	// 256 Bytes BIOS
	bootRom [0x100]byte
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
	// 128 Byte High RAM + IO mapped region
	hram [0x100]byte

	gb *GB
	biosEnabled bool
	bootRomPath string
	cartridgePath string
}

func (mmu *MMU) Init(gb *GB) error {
	boot, err := os.ReadFile(mmu.bootRomPath)
	if err != nil {
		return fmt.Errorf("could not read the boot rom, %v", err)
	}
	_, err = os.ReadFile(mmu.cartridgePath)
	if err != nil {
		return fmt.Errorf("could not read the cartridge, %v", err)
	}
	n := copy(mmu.bootRom[:], boot)
	fmt.Printf("Copied boot rom into memory: %d bytes\n", n)
	mmu.biosEnabled = true
	mmu.gb = gb
	return nil
}

func (mmu *MMU) ReadAt(addr uint16) byte {
	index := addr & 0xF000
	switch {
	case index == 0x0: {
		if mmu.gb.CPU.PC == 0x100 {
			mmu.biosEnabled = false
		}
		if mmu.biosEnabled && addr < 0x100 {
			return mmu.bootRom[addr]
		}
		return mmu.bank0[addr]
	}
	case index < 0x4000: {
		return mmu.bank0[addr]
	}
	case index < 0x8000: {
		// TODO: Change this when bank switching is implemented
		return mmu.bankN[addr & 0x3FFF]
	}
	case index < 0xA000: {
		return mmu.vram[addr & 0x1FFF]
	}
	case index < 0xC000: {
		return mmu.eram[addr & 0x1FFF]
	}
	case index < 0xF000: {
		return mmu.wram[addr & 0x1FFF]
	}
	case index == 0xF000: {
		subIndex := addr & 0x0F00
		switch {
			// Echo RAM
			case subIndex < 0xE00: {
				return mmu.wram[addr & 0x1FFF]
			}
			// OAM RAM
			case subIndex < 0xF00: {
				if addr < 0xFEA0 {
					return mmu.oam[addr & 0x00FF]
				}
				// Should not be used
				return 0
			}
			// IO and HRAM
			case subIndex == 0xF00: {
				// HRAM
				if addr >= 0xFF80 {
					return mmu.hram[addr & 0x007F]
				}
				// TODO: Handle IO operations
				return mmu.readIO(addr)
			}
		}
	}
	}
	return 0
}

// TODO: Implement display IO handling
func (mmu *MMU) readIO(addr uint16) byte {
	switch {
		// TODO: Handle various inputs (Joypad, MBC, LCD etc)
		default:
			return mmu.hram[addr - 0xFF00]
	}
}


func (mmu *MMU) WriteAt(addr uint16, val byte) {
	index := addr & 0xF000
	switch {
	case index < 0x8000: {
		return
	}
	case index < 0xA000: {
		mmu.vram[addr & 0x1FFF] = val
	}
	case index < 0xC000: {
		mmu.eram[addr & 0x1FFF] = val
	}
	case index < 0xF000: {
		mmu.wram[addr & 0x1FFF] = val
	}
	case index == 0xF000: {
		subIndex := addr & 0x0F00
		switch {
			// Echo RAM
			case subIndex < 0xE00: {
				mmu.wram[addr & 0x1FFF] = val
			}
			// OAM RAM
			case subIndex < 0xF00: {
				if addr < 0xFEA0 {
					mmu.oam[addr & 0x00FF] = val
				}
			}
			// IO and HRAM
			case subIndex == 0xF00: {
				// HRAM
				if addr >= 0xFF80 {
					// TODO: Maybe handle this differently?
					mmu.hram[addr & 0x007F] = val
				} else {
					// TODO: Handle IO operations
					mmu.writeIO(addr, val)
				}
			}
		}
	}
	}
}

// TODO: Implement display IO handling
func (mmu *MMU) writeIO(addr uint16, val byte) {
	switch {
		// TODO: Handle various outputs (Joypad, MBC, LCD etc)
		default:
			mmu.hram[addr - 0xFF00] = val
	}
}

func NewMMU(bootRomPath, cartridgePath string) *MMU {
	return &MMU{
		bootRomPath: bootRomPath,
		cartridgePath: cartridgePath,
	}
}