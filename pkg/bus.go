package pkg

import (
	"log"
)

type Bus struct {
	cpu *CPU
	ram *RAM
}

func NewBus(cpu *CPU) *Bus {
	b := &Bus{cpu: cpu}

	b.ram = NewRAM()
	return b
}

func (b *Bus) write(addr uint16, data uint8) {
	if addr < lowestBusAddress || addr > highestBusAddress {
		// panic... since if reading bus is wrong just fail immediately
		log.Fatalf("invalid address: %d", addr)
	}
	b.ram.ram[addr] = data
}

func (b *Bus) read(addr uint16, readOnly bool) uint8 {
	if addr < lowestBusAddress || addr > highestBusAddress {
		// panic... since if reading bus is wrong just fail immediately
		log.Fatalf("invalid address: %d", addr)
	}
	return b.ram.ram[addr]
}
