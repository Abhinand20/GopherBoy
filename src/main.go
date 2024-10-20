package main

import (
	"flag"
	"fmt"
	"gopherboy/pkg/gameboy"
	"os"
)

var bootRom string
var cartridge string
var debug bool

func main() {
	flag.StringVar(&bootRom, "boot_rom", "../roms/dmg_boot.bin", "The path for the boot rom binary.")
	flag.StringVar(&cartridge, "cartridge", "../roms/dmg_boot.bin", "The path for dmg game cartridge.")
	flag.BoolVar(&debug, "debug", false, "Whether to print debug logs or not.")
	
	flag.Parse()
	
	fmt.Printf("Debug: %v\n", debug)
	gb := gameboy.NewGB(bootRom, cartridge, debug)
	if err := gb.Init(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	if err := gb.Emulate(); err != nil {
		fmt.Printf("Stopped emulation: %v\n", err)
		os.Exit(1)
	}
}