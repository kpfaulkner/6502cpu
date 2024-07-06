package main

import (
	"fmt"
	"os"

	"github.com/kpfaulkner/6502cpu/pkg"
)

func loadBinaryAtMemoryLocation(bus *pkg.Bus, startAddress uint16, data []byte) {
	for i := 0; i < len(data); i++ {
		bus.Write(startAddress+uint16(i), data[i])
	}
}

func main() {
	bus := pkg.NewBus()
	cpu := pkg.NewCPU(bus)
	bus.Connect(cpu)

	//data, err := os.ReadFile("data/6502-JSR-RTS-8000-offset.bin")
	data, err := os.ReadFile("data/nestest.nes")
	// skip over first 16 bytes (NES header info afaik)
	data = data[16:]
	//data, err := os.ReadFile("data/6502-primes.bin")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	startAddr := uint16(0xC000)

	loadBinaryAtMemoryLocation(bus, startAddr, data)
	//loadBinaryAtMemoryLocation(bus, 0x0000, data)

	// reset vector
	bus.Write(0xFFFC, uint8(startAddr&0x00FF))
	bus.Write(0xFFFD, uint8((startAddr>>8)&0x00FF))

	//bus.Write(0xFFFC, 0x00)
	//bus.Write(0xFFFD, 0x00)

	cpu.Reset()
	for {
		cpu.Clock()
	}
}
