// Handles all interactions between different components (CPU, PPU and MMU).
package motherboard

import (
	"gopherboy/cpu"
	"gopherboy/mmu"
)

type Motherboard struct {
	CPU cpu.CPU
	MMU mmu.MMU
}