package pkg

import (
	"reflect"
)

type Flag uint8

const (
	lowestBusAddress  = 0x0000
	highestBusAddress = 0xFFFF

	// Flags
	C Flag = 1 << 0 // Carry Bit
	Z Flag = 1 << 1 // Zero
	I Flag = 1 << 2 // Disable Interrupts
	D Flag = 1 << 3 // Decimal Mode
	B Flag = 1 << 4 // Break
	U Flag = 1 << 5 // Unused
	V Flag = 1 << 6 // Overflow
	N Flag = 1 << 7 // Negative
)

type Instruction struct {
	name   string
	op     func() uint8
	addr   func() uint8
	cycles uint8
}

type CPU struct {
	bus *Bus

	a      uint8
	x      uint8
	y      uint8
	stkp   uint8
	pc     uint16
	status Flag

	fetched uint8

	addrAbs uint16
	addrRel uint16
	opCode  uint8
	cycles  uint8

	lookup []Instruction
}

func NewCPU(bus *Bus) *CPU {
	c := &CPU{bus: bus}
	c.lookup = c.generateLookup()
	return c
}

func (c *CPU) read(addr uint16) uint8 {
	return c.bus.Read(addr, false)
}

func (c *CPU) write(addr uint16, data uint8) {
	c.bus.Write(addr, data)
}

func (c *CPU) getFlag(f Flag) bool {
	return c.status&f > 0
}

func (c *CPU) setFlag(f Flag, v bool) {

	if v {
		c.status |= f
	} else {
		c.status &= ^f
	}
}

func (c *CPU) Clock() {
	if c.cycles == 0 {
		opCode := c.read(c.pc)
		c.pc++
		cycles := c.lookup[opCode].cycles
		additionalCycle1 := c.lookup[opCode].addr()
		additionalCycle2 := c.lookup[opCode].op()
		cycles += (additionalCycle1 & additionalCycle2)
	}
	c.cycles--
}

func (c *CPU) Reset() {
	c.a = 0
	c.x = 0
	c.y = 0
	c.stkp = 0xFD
	c.status = 0x00 | U
	c.addrAbs = 0xFFFC
	lo := uint16(c.read(c.addrAbs))
	hi := uint16(c.read(c.addrAbs + 1))
	c.pc = (hi << 8) | lo
	c.addrRel = 0
	c.addrAbs = 0
	c.fetched = 0
	c.cycles = 8

}

func (c *CPU) irq() {
	if !c.getFlag(I) {
		c.write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
		c.stkp--
		c.write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
		c.stkp--

		c.setFlag(B, false)
		c.setFlag(U, true)
		c.setFlag(I, true)
		c.write(0x0100+uint16(c.stkp), uint8(c.status))
		c.stkp--
		c.addrAbs = 0xFFFE
		lo := uint16(c.read(c.addrAbs))
		hi := uint16(c.read(c.addrAbs + 1))
		c.pc = (hi << 8) | lo
		c.cycles = 7
	}
}

func (c *CPU) nmi() {
	c.write(0x100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
	c.stkp--
	c.write(0x100+uint16(c.stkp), uint8(c.pc&0x00FF))
	c.stkp--

	c.setFlag(B, false)
	c.setFlag(U, true)
	c.setFlag(I, true)
	c.write(0x100+uint16(c.stkp), uint8(c.status))
	c.stkp--
	c.addrAbs = 0xFFFA
	lo := uint16(c.read(c.addrAbs))
	hi := uint16(c.read(c.addrAbs + 1))
	c.pc = (hi << 8) | lo
	c.cycles = 8

}

// FIXME(kpfaulkner) REALLY hate the use of reflection here.
func (c *CPU) fetch() uint8 {

	if reflect.ValueOf(c.lookup[c.opCode].addr).Pointer() != reflect.ValueOf(c.IMP).Pointer() {
		c.fetched = c.read(c.addrAbs)
	}
	return c.fetched
}
