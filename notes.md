## Gameboy architecture

### Overall specs

```
CPU: 8-bit (similar to the z80 processor)
Clock-Speed: 4.194304MHz
Screen Resolution: 160x144
Vertical Sync: 59.73Hz
Internal Memory size: 64Kb
```
### Memory

### Timing

### CPU

### PPU (display)

Basics:

- 160x144 LCD screen, does not operate on a per-pixel level but on a "tile" level
- Tile = group of 8x8 pixels (20x18 tiles total), each pixel is 2-bit ID (0-3) mapping to the pallette
- Layers: Background, Window, Object
    - Background: 32x32 grid of tilemappings in which a 20x18 viewport is visibile (also used for scrolling)
    - Window: Fixed window
    - Object: Sprite (can be 8x8 or 8x16); mapped to OAM (total 40 supported) which contains info like index into actual pixels, priority, (x,y) etc

Memory:
- VRAM ($8000-$9FFF) + OAM ($FE00-FE9F) are relevant here, which are shared by both CPU and PPU
- Tile data ($8000-$97FF, **6Kb**)
    - 1 tile = 16 bytes (2-bit per pixel), total 384 tiles
    ![alt text](imgs/tile.png)
    - Addressing modes for tile data
    ![alt text](imgs/tile_addressing.png)
    - For each line, the first byte specifies the least significant bit of the color ID of each pixel, and the second byte specifies the most significant bit.
    - Refer to https://www.huderlem.com/demos/gameboy2bpp.html

- Tile map ($98FF-$9FFF, **2Kb**)
    - This area is what defines how the Tile Data should be put together to form a BG/Window
    - Further broken down into two 32x32 tile maps (1Kb each)
    - Which tilemap should be used is controlled by LCDC.3 — BG tile map area (0;BG = map0)
    - Since one tile has 8×8 pixels, each map holds a 256×256 pixels picture. Only 160×144 of those pixels are displayed on the LCD at any given time.
    - The SCY and SCX registers can be used to scroll the Background, specifying the origin of the visible 160×144 pixel area within the total 256×256 pixel Background map.

How does the PPU use all this info to render each scanline (*ignoring sprites for now*)?
- Use SCX and SCY to determine starting offsets inside the 32x32 grid (we can only draw 20x18)
- Now, use the offsets to index into the tile map (eg. [3,3] is the starting index)
- Use tile map to fetch the tile data, repeat for all pixels
- Map to palette and render the line


Object Attribute Memory ($FE00 - $FE9F)

- Supports 40 in total, and 10 at max per line (based on prio)
- Tile data for sprites are stored in BLOCK 0 and 1 of VRAM ($8000-8FFF), functionally same as BG tile data as above
- Each OAM entry is 4 byte
![alt text](imgs/oam_bytes.png)
- **Note:** Y position here is always offset by 16 so original position is (y-16), similarly X is offset by 8
    - This means visibility conditions for sprites are:
    OAM_X > 0 AND OAM_X < 168 && OAM_Y ≤ LY + 16 < OAM_Y + Sprite_Height
- Usually data writes are done through DMA


- LCD registers
    - LCDC (FF40): Controls behavior of a lot of things described above
    - LCD Status registers
        - FF44: LY
        - FF45: LY Compare
        - FF41: STAT interrupt
- More registers for scrolling, palettes
- Sync between CPU and PPU between various PPU modes
![alt text](imgs/cpu_ppu_memory_matrix.png)

### Milestone 1

Goal: Get the Nintendo logo up and running.

- Get a basic implementation of CPU and memory in place.
- Experiment with some test ROMs.
- Have access to register values and memory for easy debugging.

###### [CPU] Instruction execution model 

```
// Mapping between opcode and function implementation
op [256]func() int
prefixOp [256]func() int

// Declaration
type inc8 func(*byte) int
type inc16 func(*uint16) int

// Definition
func (cpu *cpu.CPU) inc8(register *byte) int {
    register++
    cpu.pc += 2
    return 2
}

func (cpu *cpu.CPU) inc16(register *uint16) int {
    register++
    cpu.pc += 2
    return 2
}


// Usage
// Increment 8-bit register C
op[0x0C] = func int { return inc8(&cpu.c) }
// Increment 16-bit register BC
op[0x03] = func int { return inc16(&cpu.combine(b,c)) }

cycles := op[pc]()
```


### Debugger

1. Replace placeholders (a16, n8 etc.) with values
2. Inspect memory + debug options
3. Take input commands (step in/out, continue)
4. Breakpoint