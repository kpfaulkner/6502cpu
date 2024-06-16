package main

import (
	"fmt"
	"github.com/kpfaulkner/6502cpu/pkg"
	"os"
)

func loadBinaryAtMemoryLocation(bus *pkg.Bus, startAddress uint16, data []byte) {
	for i := 0; i < len(data); i++ {
		bus.Write(startAddress+uint16(i), data[i])
	}
}

func main() {
	fmt.Printf("so it begins....\n")

	bus := pkg.NewBus()
	cpu := pkg.NewCPU(bus)
	bus.Connect(cpu)

	data, err := os.ReadFile("data/6502-stack.bin")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	loadBinaryAtMemoryLocation(bus, 0x8000, data)

	// reset vector
	bus.Write(0xFFFC, 0x00)
	bus.Write(0xFFFD, 0x80)

	cpu.Reset()
	for {
		cpu.Clock()
	}
}
