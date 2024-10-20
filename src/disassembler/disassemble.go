package main

import (
	"flag"
	"fmt"
	"gopherboy/common"
	"os"
)

var file string
var outDir string


var instrLen = [0x100]int{1, 3, 1, 1, 1, 1, 2, 1, 3, 1, 1, 1, 1, 1, 2, 1, 2, 3, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 3, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 3, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 3, 3, 1, 2, 1, 1, 1, 3, 1, 3, 3, 2, 1, 1, 1, 3, 1, 3, 1, 2, 1, 1, 1, 3, 1, 3, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 1, 1, 2, 1}

type InstrInfo struct {
	opcode byte
	addr uint16
	instr string
}

func (ii InstrInfo) String() string {
	return fmt.Sprintf("%#4x %#2x\t%s\n", ii.addr, ii.opcode, ii.instr)
}

func disassemble(bytecode []byte) []InstrInfo {
	i := 0
	var out []InstrInfo
	for i < len(bytecode) {
		opcode := bytecode[i]
		len := instrLen[opcode]
		instrLookup := common.InstrDebugLookup
		if opcode == 0xCB {
			opcode = bytecode[i + 1]
			len = 2
			instrLookup = common.PrefixInstrDebugLookup
		}
		instrInfo := InstrInfo{
			opcode: opcode,
			addr: uint16(i),
			instr: instrLookup[opcode],
		}
		out = append(out, instrInfo)
		i += len
	}
	return out
}


func main() {
	flag.StringVar(&file, "file", "../../roms/dmg_boot.bin", "The path for the program to disassemble.")
	flag.StringVar(&outDir, "out_dir", "./", "The path for the program to disassemble.")
	flag.Parse()
	
	bytecode, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read the code file '%s' got: %v", file, err)
		os.Exit(1)
	}
	fmt.Printf("Disassembling: %v\n", file)
	src := disassemble(bytecode)
	for _,v := range src {
		fmt.Print(v.String())
	}
}