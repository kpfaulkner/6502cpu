package pkg

// addressing modes
func (c *CPU) IMP() uint8 {
	c.fetched = c.a
	return 0
}
func (c *CPU) IMM() uint8 {
	c.addrAbs = c.pc
	c.pc++
	return 0
}

func (c *CPU) ZP0() uint8 {
	c.addrAbs = uint16(c.read(c.pc))
	c.pc++
	return 0
}

func (c *CPU) ZPX() uint8 {
	c.addrAbs = uint16(c.read(c.pc) + c.x)
	c.pc++
	c.addrAbs &= 0x00FF
	return 0
}

func (c *CPU) ZPY() uint8 {
	c.addrAbs = uint16(c.read(c.pc) + c.y)
	c.pc++
	c.addrAbs &= 0x00FF
	return 0
}

func (c *CPU) REL() uint8 {
	c.addrRel = uint16(c.read(c.pc))
	c.pc++
	if c.addrRel&0x80 != 0 {
		c.addrRel |= 0xFF00
	}
	return 0
}

func (c *CPU) ABS() uint8 {
	lo := uint16(c.read(c.pc))
	c.pc++
	hi := uint16(c.read(c.pc))
	c.pc++
	c.addrAbs = (hi << 8) | lo
	return 0
}

func (c *CPU) ABX() uint8 {
	lo := uint16(c.read(c.pc))
	c.pc++
	hi := uint16(c.read(c.pc))
	c.pc++
	c.addrAbs = (hi << 8) | lo
	c.addrAbs += uint16(c.x)

	if (c.addrAbs & 0xFF00) != (hi << 8) {
		return 1
	} else {
		return 0
	}
}

func (c *CPU) ABY() uint8 {
	lo := uint16(c.read(c.pc))
	c.pc++
	hi := uint16(c.read(c.pc))
	c.pc++
	c.addrAbs = (hi << 8) | lo
	c.addrAbs += uint16(c.y)

	if (c.addrAbs & 0xFF00) != (hi << 8) {
		return 1
	} else {
		return 0
	}
}

// 6502 has a hardware bug, need to emulate thhe error.
func (c *CPU) IND() uint8 {
	lo := uint16(c.read(c.pc))
	c.pc++
	hi := uint16(c.read(c.pc))
	c.pc++
	ptr := (hi << 8) | lo
	if lo == 0x00FF {
		c.addrAbs = (uint16(c.read(ptr&0xFF00)) << 8) | uint16(c.read(ptr+0))
	} else {
		c.addrAbs = (uint16(c.read(ptr+1)) << 8) | uint16(c.read(ptr+0))
	}
	return 0
}

func (c *CPU) IZX() uint8 {

	t := uint16(c.read(c.pc))
	c.pc++
	lo := uint16(c.read(uint16(t+uint16(c.x)) & 0x00FF))
	hi := uint16(c.read(uint16(t+uint16(c.x)+1) & 0x00FF))
	c.addrAbs = (hi << 8) | lo
	return 0
}

func (c *CPU) IZY() uint8 {
	t := uint16(c.read(c.pc))
	c.pc++
	lo := uint16(c.read(t & 0x00FF))
	hi := uint16(c.read(uint16(t+1) & 0x00FF))
	c.addrAbs = (hi << 8) | lo
	c.addrAbs += uint16(c.y)

	if c.addrAbs&0xFF00 != (hi << 8) {
		return 1
	} else {
		return 0
	}
}
