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