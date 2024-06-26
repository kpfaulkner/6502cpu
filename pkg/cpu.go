package pkg

import (
	"fmt"
)

type Flag uint8

func (f Flag) String() string {
	s := ""
	if f&N > 0 {
		s += "N"
	} else {
		s += "-"
	}
	if f&V > 0 {
		s += "V"
	} else {
		s += "-"
	}
	if f&U > 0 {
		s += "U"
	} else {
		s += "-"
	}
	if f&B > 0 {
		s += "B"
	} else {
		s += "-"
	}
	if f&D > 0 {
		s += "D"
	} else {
		s += "-"
	}
	if f&I > 0 {
		s += "I"
	} else {
		s += "-"
	}
	if f&Z > 0 {
		s += "Z"
	} else {
		s += "-"
	}
	if f&C > 0 {
		s += "C"
	} else {
		s += "-"
	}
	return s
}

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
	addr   func() (uint8, addrMode)
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

	// current AddressMode
	addrMode addrMode
	addrAbs  uint16
	addrRel  uint16
	opCode   uint8
	cycles   uint8

	lookup []Instruction

	stackSnoop  []uint8
	memorySnoop []uint8
}

func NewCPU(bus *Bus) *CPU {
	c := &CPU{bus: bus}
	c.lookup = c.generateLookup()
	c.stackSnoop = c.bus.ram.ram[0x0100+240 : 0x0100+256]
	c.memorySnoop = c.bus.ram.ram[0xC000 : 0xC000+40]
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
		fmt.Printf("OP: %s\n", c.lookup[opCode].name)
		c.pc++
		cycles := c.lookup[opCode].cycles
		additionalCycle1, addrMode := c.lookup[opCode].addr()
		c.addrMode = addrMode
		additionalCycle2 := c.lookup[opCode].op()

		cycles += (additionalCycle1 & additionalCycle2)
		c.dump()
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

	if c.addrMode != IMP {
		c.fetched = c.read(c.addrAbs)
	}
	return c.fetched
}

// dump internals out to stdout
func (c *CPU) dump() {
	fmt.Printf("A %02x X %02x Y %02x STKP %02x PC %04x Flag %s\n", c.a, c.x, c.y, c.stkp, c.pc, c.status.String())
}
