package pkg

import (
	"log"
)

type Bus struct {
	cpu *CPU
	ram *RAM
}

func NewBus() *Bus {
	b := &Bus{}
	b.ram = NewRAM()
	return b
}

func (b *Bus) Connect(cpu *CPU) {
	b.cpu = cpu
}

func (b *Bus) Write(addr uint16, data uint8) {
	if addr < lowestBusAddress || addr > highestBusAddress {
		// panic... since if reading bus is wrong just fail immediately
		log.Fatalf("invalid address: %d", addr)
	}
	b.ram.ram[addr] = data
}

func (b *Bus) Read(addr uint16, readOnly bool) uint8 {
	if addr < lowestBusAddress || addr > highestBusAddress {
		// panic... since if reading bus is wrong just fail immediately
		log.Fatalf("invalid address: %d", addr)
	}
	return b.ram.ram[addr]
}
