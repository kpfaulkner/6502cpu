package pkg

// opcodes
func (c *CPU) ADC() uint8 {

	c.fetch()

	temp := uint16(c.a) + uint16(c.fetched) + uint16(c.getFlag(C))
	c.setFlag(C, temp > 255)
	c.setFlag(Z, (temp&0x00FF) == 0)
	c.setFlag(N, temp&0x80 == 1)
	c.setFlag(V, (^(uint16(c.a)^uint16(c.fetched))&(uint16(c.a)^temp)&0x0080) == 1)
	c.a = uint8(temp & 0x00FF)
	return 1
}

func (c *CPU) AND() uint8 {
	c.fetch()
	c.a = c.a & c.fetched
	c.setFlag(Z, c.a == 0x00)
	c.setFlag(N, c.a&0x80 == 1)
	return 1
}

func (c *CPU) ASL() uint8 {
	return 0
}

func (c *CPU) BCC() uint8 {
	if c.getFlag(C) == 0 {
		c.cycles++
		c.addrAbs = c.pc + c.addrRel

		if (c.addrAbs & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addrAbs
	}
	return 0
}

func (c *CPU) BCS() uint8 {
	if c.getFlag(C) == 1 {
		c.cycles++
		c.addrAbs = c.pc + c.addrRel

		if (c.addrAbs & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addrAbs

	}
	return 0
}

func (c *CPU) BEQ() uint8 {
	if c.getFlag(Z) == 1 {
		c.cycles++
		c.addrAbs = c.pc + c.addrRel

		if (c.addrAbs & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addrAbs
	}

	return 0
}

func (c *CPU) BIT() uint8 {
	return 0
}

func (c *CPU) BMI() uint8 {
	if c.getFlag(N) == 1 {
		c.cycles++
		c.addrAbs = c.pc + c.addrRel

		if (c.addrAbs & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addrAbs

	}
	return 0
}

func (c *CPU) BNE() uint8 {
	if c.getFlag(Z) == 0 {
		c.cycles++
		c.addrAbs = c.pc + c.addrRel

		if (c.addrAbs & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addrAbs

	}
	return 0
}

func (c *CPU) BPL() uint8 {
	if c.getFlag(N) == 0 {
		c.cycles++
		c.addrAbs = c.pc + c.addrRel

		if (c.addrAbs & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addrAbs

	}
	return 0
}

func (c *CPU) BRK() uint8 {
	return 0
}

func (c *CPU) BVC() uint8 {
	if c.getFlag(V) == 0 {
		c.cycles++
		c.addrAbs = c.pc + c.addrRel

		if (c.addrAbs & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addrAbs

	}
	return 0
}

func (c *CPU) BVS() uint8 {
	if c.getFlag(V) == 1 {
		c.cycles++
		c.addrAbs = c.pc + c.addrRel

		if (c.addrAbs & 0xFF00) != (c.pc & 0xFF00) {
			c.cycles++
		}

		c.pc = c.addrAbs
	}
	return 0
}

func (c *CPU) CLC() uint8 {
	c.setFlag(C, false)
	return 0
}

func (c *CPU) CLD() uint8 {
	c.setFlag(D, false)
	return 0
}

func (c *CPU) CLI() uint8 {

	c.setFlag(I, false)
	return 0

}

func (c *CPU) CLV() uint8 {
	c.setFlag(V, false)
	return 0
}

func (c *CPU) CMP() uint8 {
	return 0
}

func (c *CPU) CPX() uint8 {
	return 0
}

func (c *CPU) CPY() uint8 {
	return 0
}

func (c *CPU) DEC() uint8 {
	return 0
}
func (c *CPU) DEX() uint8 {
	c.x--
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 == 1)
	return 0
}

// decrement Y register
func (c *CPU) DEY() uint8 {

	c.y--
	c.setFlag(Z, c.y == 0x00)
	c.setFlag(N, c.y&0x80 == 1)
	return 0
}

func (c *CPU) EOR() uint8 {
	return 0
}

func (c *CPU) INC() uint8 {
	return 0
}
func (c *CPU) INX() uint8 {
	return 0
}
func (c *CPU) INY() uint8 {
	return 0
}
func (c *CPU) JMP() uint8 {
	return 0
}

func (c *CPU) JSR() uint8 {
	return 0
}
func (c *CPU) LDA() uint8 {
	c.fetch()
	c.a = c.fetched
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 == 1)
	return 0
}
func (c *CPU) LDX() uint8 {
	c.fetch()
	c.x = c.fetched
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 == 1)
	return 0
}
func (c *CPU) LDY() uint8 {

	c.fetch()
	c.y = c.fetched
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 == 1)
	return 0
}

func (c *CPU) LSR() uint8 {
	return 0
}
func (c *CPU) NOP() uint8 {
	return 0
}
func (c *CPU) ORA() uint8 {
	return 0
}
func (c *CPU) PHA() uint8 {
	c.write(0x0100+uint16(c.stkp), c.a)
	c.stkp--
	return 0
}

func (c *CPU) PHP() uint8 {
	return 0
}
func (c *CPU) PLA() uint8 {
	c.stkp++
	c.a = c.read(0x0100 + uint16(c.stkp))
	c.setFlag(Z, c.a == 0x00)
	c.setFlag(N, c.a&0x80 == 1)
	return 0
}
func (c *CPU) PLP() uint8 {
	return 0
}
func (c *CPU) ROL() uint8 {
	return 0
}

func (c *CPU) ROR() uint8 {
	return 0
}
func (c *CPU) RTI() uint8 {
	c.stkp++
	c.status = Flag(c.read(0x0100 + uint16(c.stkp)))
	c.status &= ^B
	c.status &= ^U
	c.stkp++
	c.pc = uint16(c.read(0x0100 + uint16(c.stkp)))
	c.stkp++
	c.pc |= uint16(c.read(0x0100+uint16(c.stkp))) << 8
	return 0
}

func (c *CPU) RTS() uint8 {
	return 0
}

func (c *CPU) SBC() uint8 {
	c.fetch()

	value := uint16(c.fetched) ^ 0x00FF

	temp := uint16(c.a) + value + uint16(c.getFlag(C))
	c.setFlag(C, temp > 255)
	c.setFlag(Z, (temp&0x00FF) == 0)
	c.setFlag(N, temp&0x80 == 1)
	c.setFlag(V, (^(uint16(c.a)^uint16(c.fetched))&(uint16(c.a)^temp)&0x0080) == 1)
	c.a = uint8(temp & 0x00FF)
	return 1
}

func (c *CPU) SEC() uint8 {
	return 0
}
func (c *CPU) SED() uint8 {
	return 0
}
func (c *CPU) SEI() uint8 {
	return 0
}
func (c *CPU) STA() uint8 {
	c.bus.Write(c.addrAbs, c.a)
	return 0
}

// stores X register to addrAbs ?
func (c *CPU) STX() uint8 {
	c.bus.Write(c.addrAbs, c.x)
	return 0
}
func (c *CPU) STY() uint8 {
	return 0
}
func (c *CPU) TAX() uint8 {
	return 0
}
func (c *CPU) TAY() uint8 {
	return 0
}

func (c *CPU) TSX() uint8 {
	return 0
}
func (c *CPU) TXA() uint8 {
	return 0
}
func (c *CPU) TXS() uint8 {
	return 0
}
func (c *CPU) TYA() uint8 {
	return 0
}

// illegal
func (c *CPU) XXX() uint8 {
	return 0
}

func addInstruction(name string, op func() uint8, addr func() uint8, cycles uint8) Instruction {
	i := Instruction{
		name:   name,
		op:     op,
		addr:   addr,
		cycles: cycles,
	}
	return i
}

func (c *CPU) generateLookup() []Instruction {
	lookup := []Instruction{}
	lookup = append(lookup, addInstruction("BRK", c.BRK, c.IMM, 7))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 3))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ZP0, 3))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.ZP0, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("PHP", c.PHP, c.IMP, 3))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.IMM, 2))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ABS, 4))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.ABS, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("BPL", c.BPL, c.REL, 2))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ZPX, 4))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.ZPX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("CLC", c.CLC, c.IMP, 2))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ABY, 4))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ABX, 4))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.ABX, 7))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("JSR", c.JSR, c.ABS, 6))
	lookup = append(lookup, addInstruction("AND", c.AND, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("BIT", c.BIT, c.ZP0, 3))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ZP0, 3))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.ZP0, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("PLP", c.PLP, c.IMP, 4))
	lookup = append(lookup, addInstruction("AND", c.AND, c.IMM, 2))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("BIT", c.BIT, c.ABS, 4))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ABS, 4))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.ABS, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("BMI", c.BMI, c.REL, 2))
	lookup = append(lookup, addInstruction("AND", c.AND, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ZPX, 4))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.ZPX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("SEC", c.SEC, c.IMP, 2))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ABY, 4))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ABX, 4))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.ABX, 7))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("RTI", c.RTI, c.IMP, 6))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 3))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.ZP0, 3))
	lookup = append(lookup, addInstruction("LSR", c.LSR, c.ZP0, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("PHA", c.PHA, c.IMP, 3))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.IMM, 2))
	lookup = append(lookup, addInstruction("LSR", c.LSR, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("JMP", c.JMP, c.ABS, 3))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.ABS, 4))
	lookup = append(lookup, addInstruction("LSR", c.LSR, c.ABS, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("BVC", c.BVC, c.REL, 2))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.ZPX, 4))
	lookup = append(lookup, addInstruction("LSR", c.LSR, c.ZPX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("CLI", c.CLI, c.IMP, 2))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.ABY, 4))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.ABX, 4))
	lookup = append(lookup, addInstruction("LSR", c.LSR, c.ABX, 7))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("RTS", c.RTS, c.IMP, 6))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 3))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.ZP0, 3))
	lookup = append(lookup, addInstruction("ROR", c.ROR, c.ZP0, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("PLA", c.PLA, c.IMP, 4))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.IMM, 2))
	lookup = append(lookup, addInstruction("ROR", c.ROR, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("JMP", c.JMP, c.IND, 5))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.ABS, 4))
	lookup = append(lookup, addInstruction("ROR", c.ROR, c.ABS, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("BVS", c.BVS, c.REL, 2))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.ZPX, 4))
	lookup = append(lookup, addInstruction("ROR", c.ROR, c.ZPX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("SEI", c.SEI, c.IMP, 2))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.ABY, 4))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.ABX, 4))
	lookup = append(lookup, addInstruction("ROR", c.ROR, c.ABX, 7))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("STA", c.STA, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("STY", c.STY, c.ZP0, 3))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ZP0, 3))
	lookup = append(lookup, addInstruction("STX", c.STX, c.ZP0, 3))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 3))
	lookup = append(lookup, addInstruction("DEY", c.DEY, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("TXA", c.TXA, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("STY", c.STY, c.ABS, 4))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ABS, 4))
	lookup = append(lookup, addInstruction("STX", c.STX, c.ABS, 4))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 4))
	lookup = append(lookup, addInstruction("BCC", c.BCC, c.REL, 2))
	lookup = append(lookup, addInstruction("STA", c.STA, c.IZY, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("STY", c.STY, c.ZPX, 4))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ZPX, 4))
	lookup = append(lookup, addInstruction("STX", c.STX, c.ZPY, 4))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 4))
	lookup = append(lookup, addInstruction("TYA", c.TYA, c.IMP, 2))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ABY, 5))
	lookup = append(lookup, addInstruction("TXS", c.TXS, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 5))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ABX, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.IMM, 2))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.IZX, 6))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.IMM, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.ZP0, 3))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ZP0, 3))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.ZP0, 3))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 3))
	lookup = append(lookup, addInstruction("TAY", c.TAY, c.IMP, 2))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.IMM, 2))
	lookup = append(lookup, addInstruction("TAX", c.TAX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.ABS, 4))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ABS, 4))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.ABS, 4))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 4))
	lookup = append(lookup, addInstruction("BCS", c.BCS, c.REL, 2))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.ZPX, 4))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ZPX, 4))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.ZPY, 4))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 4))
	lookup = append(lookup, addInstruction("CLV", c.CLV, c.IMP, 2))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ABY, 4))
	lookup = append(lookup, addInstruction("TSX", c.TSX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 4))
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.ABX, 4))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ABX, 4))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.ABY, 4))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 4))
	lookup = append(lookup, addInstruction("CPY", c.CPY, c.IMM, 2))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("CPY", c.CPY, c.ZP0, 3))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ZP0, 3))
	lookup = append(lookup, addInstruction("DEC", c.DEC, c.ZP0, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("INY", c.INY, c.IMP, 2))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.IMM, 2))
	lookup = append(lookup, addInstruction("DEX", c.DEX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("CPY", c.CPY, c.ABS, 4))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ABS, 4))
	lookup = append(lookup, addInstruction("DEC", c.DEC, c.ABS, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("BNE", c.BNE, c.REL, 2))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ZPX, 4))
	lookup = append(lookup, addInstruction("DEC", c.DEC, c.ZPX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("CLD", c.CLD, c.IMP, 2))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ABY, 4))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ABX, 4))
	lookup = append(lookup, addInstruction("DEC", c.DEC, c.ABX, 7))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("CPX", c.CPX, c.IMM, 2))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("CPX", c.CPX, c.ZP0, 3))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ZP0, 3))
	lookup = append(lookup, addInstruction("INC", c.INC, c.ZP0, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("INX", c.INX, c.IMP, 2))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.IMM, 2))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.SBC, c.IMP, 2))
	lookup = append(lookup, addInstruction("CPX", c.CPX, c.ABS, 4))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ABS, 4))
	lookup = append(lookup, addInstruction("INC", c.INC, c.ABS, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("BEQ", c.BEQ, c.REL, 2))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ZPX, 4))
	lookup = append(lookup, addInstruction("INC", c.INC, c.ZPX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("SED", c.SED, c.IMP, 2))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ABY, 4))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 4))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ABX, 4))
	lookup = append(lookup, addInstruction("INC", c.INC, c.ABX, 7))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))

	return lookup
}
