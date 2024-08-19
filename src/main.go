package main

import (
	"flag"
	"fmt"
	"gopherboy/motherboard"
	"os"
)

var bootRom string
var cartridge string

func main() {
	flag.StringVar(&bootRom, "boot_rom", "../roms/dmg_boot.bin", "The path for the boot rom binary.")
	flag.StringVar(&cartridge, "cartridge", "../roms/dmg_boot.bin", "The path for dmg game cartridge.")

	mb, err := motherboard.NewMotherboard(bootRom, cartridge)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	if err := mb.Emulate(); err != nil {
		fmt.Printf("Stopped emulation: %v\n", err)
		os.Exit(1)
	}
}