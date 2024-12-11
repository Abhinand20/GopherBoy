package gameboy

/*
Handles a PPU "tick". Synchronization is achieved by emulating
as many clocks as the CPU had moved forward by in the current tick (per-instruction synchronization).

PPU is modeled as a state machine moving between Modes 0-3.
*/
import (
	"image/color"
)

type PPUState uint8

const (
	OAMSearch = iota
	Draw
	Hblank
	Vblank
)


type PPU struct {
	
	// TODO: Handle registers
	State PPUState
	FrameBuffer [160*144]color.RGBA
	// cycles in current scanline (move to next scanline after 456 cycles)
	nCycles uint16
	// currently rendering scanline, resets after 153 and enters VBlank
	nScanline uint8

	gb *GB
}

func (ppu *PPU) Init(gb *GB) {
	ppu.gb = gb
	ppu.nCycles = 0
	ppu.State = OAMSearch
	ppu.nScanline = 0
}

func (ppu *PPU) Tick(cpuCycles uint8) {
	// TODO: Update this
	ppu.nCycles += uint16(cpuCycles)

	switch ppu.State {
	case OAMSearch: {
		if ppu.nCycles >= 80 {
			// TODO: Implement OAM search logic here
			ppu.nCycles -= 80
			ppu.State = Draw
		}
	}
	case Draw: {
		if ppu.nCycles >= 172 {
			ppu.nCycles -= 172
			// TODO: Implement draw logic here
			ppu.State = Hblank
		}
	}
	case Hblank: {
		// 456 - (OAM + Draw) = 204
		if ppu.nCycles >= 204 {
			ppu.nCycles -= 204
			ppu.nScanline++
			// TODO: Implement Hblank logic here
			if ppu.nScanline >= 144 {
				ppu.State = Vblank
				// TODO: Implement Vblank interrupt
			} else {
				ppu.State = OAMSearch
			}
		}
	}
	case Vblank: {
		if ppu.nCycles >= 456 {
			ppu.nCycles -= 456
			ppu.nScanline++
			if ppu.nScanline == 154 {
				ppu.nScanline = 0
				ppu.State = OAMSearch
			}
		}
	}
	}

}