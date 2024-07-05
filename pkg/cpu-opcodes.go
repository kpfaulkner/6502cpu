package pkg

import (
	"fmt"
	"time"
)

// opcodes
func (c *CPU) ADC() uint8 {

	c.fetch()

	var carryBit uint16
	if c.getFlag(C) {
		carryBit = 1
	}

	temp := uint16(c.a) + uint16(c.fetched) + carryBit
	c.setFlag(C, temp > 255)
	c.setFlag(Z, (temp&0x00FF) == 0)
	c.setFlag(N, temp&0x80 > 0)
	c.setFlag(V, (^(uint16(c.a)^uint16(c.fetched))&(uint16(c.a)^temp)&0x0080) > 0)
	c.a = uint8(temp & 0x00FF)
	return 1
}

func (c *CPU) AND() uint8 {
	c.fetch()
	c.a = c.a & c.fetched
	c.setFlag(Z, c.a == 0x00)
	c.setFlag(N, c.a&0x80 > 0)
	return 1
}

func (c *CPU) ASL() uint8 {

	c.fetch()
	temp := uint16(c.fetched) << 1
	c.setFlag(C, (temp&0xFF00) > 0)
	c.setFlag(Z, (temp&0x00FF) == 0)
	c.setFlag(N, temp&0x80 > 0)
	if c.addrMode == IMP {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.write(c.addrAbs, uint8(temp&0x00FF))
	}
	return 0
}

func (c *CPU) BCC() uint8 {
	if !c.getFlag(C) {
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
	if c.getFlag(C) {
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
	if c.getFlag(Z) {
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

	c.fetch()
	temp := c.a & c.fetched
	c.setFlag(Z, (temp&0x00FF) == 0)
	c.setFlag(N, (c.fetched&(1<<7)) > 0)
	c.setFlag(V, (c.fetched&(1<<6)) > 0)
	return 0
}

func (c *CPU) BMI() uint8 {
	if c.getFlag(N) {
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
	if !c.getFlag(Z) {
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
	if !c.getFlag(N) {
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

	c.pc++
	c.setFlag(I, true)
	c.write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
	c.stkp--
	c.write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
	c.stkp--
	c.setFlag(B, true)
	c.write(0x0100+uint16(c.stkp), uint8(c.status))
	c.stkp--
	c.setFlag(B, false)
	c.pc = uint16(c.read(0xFFFE)) | (uint16(c.read(0xFFFF)) << 8)

	// just pause here... so we can look at stuff?
	time.Sleep(10000 * time.Second)
	return 0
}

func (c *CPU) BVC() uint8 {
	if !c.getFlag(V) {
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
	if c.getFlag(V) {
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

	c.fetch()
	temp := uint16(c.a) - uint16(c.fetched)
	c.setFlag(C, c.a >= c.fetched)
	c.setFlag(Z, temp&0x00FF == 0) // check.
	c.setFlag(N, temp&0x80 > 0)
	return 1
}

// CPX: Compare memory and index X
// Remember for comparison we use the flags to indicate > = or <
// See http://www.6502.org/tutorials/compare_instructions.html
func (c *CPU) CPX() uint8 {

	c.fetch()
	temp := uint16(c.x) - uint16(c.fetched)
	c.setFlag(C, c.x >= c.fetched)
	c.setFlag(Z, temp&0x00FF == 0) // check.
	c.setFlag(N, temp&0x80 > 0)

	return 0
}

func (c *CPU) CPY() uint8 {
	c.fetch()
	temp := uint16(c.y) - uint16(c.fetched)
	c.setFlag(C, c.y >= c.fetched)
	c.setFlag(Z, temp&0x00FF == 0) // check.
	c.setFlag(N, temp&0x80 > 0)

	return 0
}

func (c *CPU) DCP() uint8 {

	c.fetch()
	temp := uint16(c.fetched) - 1
	c.write(c.addrAbs, uint8(temp&0x00FF))
	temp2 := uint16(c.a) - temp
	//c.setFlag(C, uint16(c.a) >= temp)

	c.setFlag(C, c.a >= c.fetched)
	//c.setFlag(C, temp2&0xFF00 > 0)
	c.setFlag(Z, (temp2&0x00FF) == 0x00)
	c.setFlag(N, temp2&0x80 > 0)

	return 0
}

func (c *CPU) DEC() uint8 {
	c.fetch()
	temp := c.fetched - 1
	c.write(c.addrAbs, temp&0x00FF)
	c.setFlag(Z, (temp&0x00FF) == 0x00)
	c.setFlag(N, temp&0x80 > 0)

	return 0
}

func (c *CPU) DEX() uint8 {
	c.x--
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 > 0)
	return 0
}

// decrement Y register
func (c *CPU) DEY() uint8 {

	c.y--
	c.setFlag(Z, c.y == 0x00)
	c.setFlag(N, c.y&0x80 > 0)
	return 0
}

func (c *CPU) EOR() uint8 {
	c.fetch()
	c.a = c.a ^ c.fetched
	c.setFlag(Z, c.a == 0x00)
	c.setFlag(N, c.a&0x80 > 0)
	return 0
}

func (c *CPU) INC() uint8 {
	c.fetch()
	temp := c.fetched + 1
	c.write(c.addrAbs, temp&0x00FF)
	c.setFlag(Z, temp == 0x00)
	c.setFlag(N, temp&0x80 > 0)
	return 0
}

// INX: increment X
func (c *CPU) INX() uint8 {

	c.x++
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 > 0)
	return 0
}

// INX: increment Y
func (c *CPU) INY() uint8 {
	c.y++
	c.setFlag(Z, c.y == 0x00)
	c.setFlag(N, c.y&0x80 > 0)
	return 0
}

func (c *CPU) ISBWIP() uint8 {
	//panic("TODO")
	c.fetch()
	temp := uint16(c.fetched) + 1
	c.write(c.addrAbs, uint8(temp&0x00FF))

	value := temp ^ 0x00FF

	var carryBit uint16
	if c.getFlag(C) {
		carryBit = 1
	}

	temp2 := uint16(c.a) - temp - uint16(carryBit)
	c.setFlag(C, temp2&0xFF00 > 0)
	c.setFlag(Z, (temp2&0x00FF) == 0)
	c.setFlag(N, temp2&0x80 > 0)
	//c.setFlag(V, (^(uint16(c.a)^uint16(c.fetched))&(uint16(c.a)^temp)&0x0080) > 0)
	c.setFlag(V, ((temp2^uint16(c.a))&(temp2^value)&0x0080) > 0)

	return 0

	return 0
}

func (c *CPU) ISB() uint8 {
	c.INC()
	c.SBC()
	return 0
}

func (c *CPU) JMP() uint8 {
	c.pc = c.addrAbs
	return 0
}

// JSR: take current program counter, push (MSB then LSB) to stack
// then set the pc to the absoluteAddress that was read in previously.
func (c *CPU) JSR() uint8 {
	c.pc--

	// MSB
	c.write(0x0100+uint16(c.stkp), uint8((c.pc>>8)&0x00FF))
	c.stkp--

	// LSB
	c.write(0x0100+uint16(c.stkp), uint8(c.pc&0x00FF))
	c.stkp--

	// now PC is where the absolute address is set.
	c.pc = c.addrAbs
	return 0
}

func (c *CPU) LAX() uint8 {
	c.fetch()
	c.a = c.fetched
	c.x = c.fetched
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 > 0)
	return 0
}

func (c *CPU) LDA() uint8 {
	c.fetch()
	c.a = c.fetched
	c.setFlag(Z, c.a == 0x00)
	c.setFlag(N, c.a&0x80 > 0)
	return 1
}

func (c *CPU) LDX() uint8 {
	c.fetch()
	c.x = c.fetched
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 > 0)
	return 1
}
func (c *CPU) LDY() uint8 {

	c.fetch()
	c.y = c.fetched
	c.setFlag(Z, c.y == 0x00)
	c.setFlag(N, c.y&0x80 > 0)
	return 1
}

// LSR: Logical Shift Right
// if implied addressing mode, then we're just shifting the accumulator right one bit.
func (c *CPU) LSR() uint8 {
	c.fetch()

	// if LSB is 1, set carry flag.
	if c.fetched&0x0001 > 0 {
		c.setFlag(C, true)
	} else {
		c.setFlag(C, false)
	}

	temp := c.fetched >> 1
	c.setFlag(Z, (temp&0x00FF) == 0)

	// if MSB is 1 of the byte, set N flag.
	c.setFlag(N, temp&0x80 > 0)

	// if implied mode, then just set A
	if c.addrMode == IMP {
		c.a = temp & 0x00FF
	} else {
		c.write(c.addrAbs, temp&0x00FF)
	}

	return 0
}

func (c *CPU) NOP() uint8 {
	c.fetch()
	switch c.opCode {
	case 0x1C:
	case 0x3C:
	case 0x5C:
	case 0x7C:
	case 0xDC:
	case 0xFC:
		return 1
		break
	}
	return 0
}

func (c *CPU) ORA() uint8 {

	c.fetch()
	c.a = c.a | c.fetched
	c.setFlag(Z, c.a == 0x00)
	c.setFlag(N, c.a&0x80 > 0)
	return 0
}

func (c *CPU) PHA() uint8 {
	c.write(0x0100+uint16(c.stkp), c.a)
	c.stkp--
	return 0
}

func (c *CPU) PHP() uint8 {
	c.write(0x0100+uint16(c.stkp), uint8(c.status|B|U))
	c.setFlag(B, false)
	c.setFlag(U, false)
	c.stkp--
	return 0
}
func (c *CPU) PLA() uint8 {
	c.stkp++
	c.a = c.read(0x0100 + uint16(c.stkp))
	c.setFlag(Z, c.a == 0x00)
	c.setFlag(N, c.a&0x80 > 0)
	return 0
}
func (c *CPU) PLP() uint8 {
	c.stkp++
	c.status = Flag(c.read(0x0100 + uint16(c.stkp)))
	c.setFlag(U, true)
	return 0
}

func (c *CPU) RLA() uint8 {
	c.ROL()
	c.AND()

	return 0
}

func (c *CPU) ROL() uint8 {
	c.fetch()
	var v uint16
	if c.getFlag(C) {
		v = 1
	}
	temp := uint16(c.fetched) << 1
	temp = temp | v
	c.setFlag(C, temp&0xFF00 > 0)
	c.setFlag(Z, (temp&0x00FF) == 0)
	c.setFlag(N, temp&0x80 > 0)
	if c.addrMode == IMP {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.write(c.addrAbs, uint8(temp&0x00FF))

	}
	return 0
}

func (c *CPU) ROR() uint8 {
	c.fetch()
	var v uint16
	if c.getFlag(C) {
		v = 0x80
	}
	temp := v | uint16(c.fetched>>1)
	c.setFlag(C, c.fetched&0x01 > 0)
	c.setFlag(Z, (temp&0x00FF) == 0)
	c.setFlag(N, temp&0x80 > 0)
	if c.addrMode == IMP {
		c.a = uint8(temp & 0x00FF)
	} else {
		c.write(c.addrAbs, uint8(temp&0x00FF))
	}
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

// RTS: restore pc from the stack.
func (c *CPU) RTS() uint8 {

	c.stkp++
	// LSB
	temp := uint16(c.read(0x0100 + uint16(c.stkp)))

	c.stkp++
	// MSB
	temp = temp + (uint16(c.read(0x0100+uint16(c.stkp))) << 8)

	// next instruction.
	c.pc = temp + 1
	return 0
}

func (c *CPU) SAX() uint8 {

	temp := c.a & c.x
	c.bus.Write(c.addrAbs, temp)
	return 0
}

func (c *CPU) SBC() uint8 {
	c.fetch()

	value := uint16(c.fetched) ^ 0x00FF

	var carryBit uint16
	if c.getFlag(C) {
		carryBit = 1
	}
	temp := uint16(c.a) + value + carryBit
	c.setFlag(C, temp&0xFF00 > 0)
	c.setFlag(Z, (temp&0x00FF) == 0)
	c.setFlag(N, temp&0x80 > 0)
	//c.setFlag(V, (^(uint16(c.a)^uint16(c.fetched))&(uint16(c.a)^temp)&0x0080) > 0)
	c.setFlag(V, ((temp^uint16(c.a))&(temp^value)&0x0080) > 0)
	c.a = uint8(temp & 0x00FF)
	return 1
}

func (c *CPU) SEC() uint8 {
	c.setFlag(C, true)
	return 0
}
func (c *CPU) SED() uint8 {
	c.setFlag(D, true)
	return 0
}
func (c *CPU) SEI() uint8 {
	c.setFlag(I, true)
	return 0
}

func (c *CPU) SLO() uint8 {
	c.ASL()
	c.ORA()
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
	c.write(c.addrAbs, c.y)
	return 0
}

// TAX: Transfer accumulator to X
func (c *CPU) TAX() uint8 {
	c.x = c.a
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 > 0)

	return 0
}
func (c *CPU) TAY() uint8 {
	c.y = c.a
	c.setFlag(Z, c.y == 0x00)
	c.setFlag(N, c.y&0x80 > 0)

	return 0
}

func (c *CPU) TSX() uint8 {
	c.x = c.stkp
	c.setFlag(Z, c.x == 0x00)
	c.setFlag(N, c.x&0x80 > 0)
	return 0
}
func (c *CPU) TXA() uint8 {
	c.a = c.x
	c.setFlag(Z, c.a == 0x00)
	c.setFlag(N, c.a&0x80 > 0)
	return 0
}
func (c *CPU) TXS() uint8 {

	c.stkp = c.x
	return 0
}
func (c *CPU) TYA() uint8 {
	c.a = c.y
	c.setFlag(Z, c.a == 0x00)
	c.setFlag(N, c.a&0x80 > 0)
	return 0
}

// illegal
func (c *CPU) XXX() uint8 {
	//panic("XXX")
	fmt.Printf("XXXXXXXXXXXXX\n")
	return 0
}

func addInstruction(name string, op func() uint8, addr func() (uint8, addrMode), cycles uint8) Instruction {
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

	// 0
	lookup = append(lookup, addInstruction("BRK", c.BRK, c.IMM, 7))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("SLO", c.SLO, c.IZX, 8))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ZP0, 3))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ZP0, 3))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.ZP0, 5))
	lookup = append(lookup, addInstruction("SLO", c.SLO, c.ZP0, 5))
	lookup = append(lookup, addInstruction("PHP", c.PHP, c.IMP, 3))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.IMM, 2))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ABS, 4))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ABS, 4))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.ABS, 6))
	lookup = append(lookup, addInstruction("SLO", c.SLO, c.ABS, 6))

	// 1
	lookup = append(lookup, addInstruction("BPL", c.BPL, c.REL, 2))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("SLO", c.SLO, c.IZY, 8))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ZPX, 4))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ZPX, 4))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.ZPX, 6))
	lookup = append(lookup, addInstruction("SLO", c.SLO, c.ZPX, 6))
	lookup = append(lookup, addInstruction("CLC", c.CLC, c.IMP, 2))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ABY, 4))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("SLO", c.SLO, c.ABY, 7))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ABX, 4))
	lookup = append(lookup, addInstruction("ORA", c.ORA, c.ABX, 4))
	lookup = append(lookup, addInstruction("ASL", c.ASL, c.ABX, 7))
	lookup = append(lookup, addInstruction("SLO", c.SLO, c.ABX, 7))

	// 2
	lookup = append(lookup, addInstruction("JSR", c.JSR, c.ABS, 6))
	lookup = append(lookup, addInstruction("AND", c.AND, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("RLA", c.RLA, c.IZX, 8))
	lookup = append(lookup, addInstruction("BIT", c.BIT, c.ZP0, 3))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ZP0, 3))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.ZP0, 5))
	lookup = append(lookup, addInstruction("RLA", c.RLA, c.ZP0, 5))
	lookup = append(lookup, addInstruction("PLP", c.PLP, c.IMP, 4))
	lookup = append(lookup, addInstruction("AND", c.AND, c.IMM, 2))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("BIT", c.BIT, c.ABS, 4))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ABS, 4))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.ABS, 6))
	lookup = append(lookup, addInstruction("RLA", c.RLA, c.ABS, 6))

	// 3
	lookup = append(lookup, addInstruction("BMI", c.BMI, c.REL, 2))
	lookup = append(lookup, addInstruction("AND", c.AND, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("RLA", c.RLA, c.IZY, 8))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ZPX, 4))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ZPX, 4))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.ZPX, 6))
	lookup = append(lookup, addInstruction("RLA", c.RLA, c.ZPX, 6))
	lookup = append(lookup, addInstruction("SEC", c.SEC, c.IMP, 2))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ABY, 4))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("RLA", c.RLA, c.ABY, 7))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ABX, 4))
	lookup = append(lookup, addInstruction("AND", c.AND, c.ABX, 4))
	lookup = append(lookup, addInstruction("ROL", c.ROL, c.ABX, 7))
	lookup = append(lookup, addInstruction("RLA", c.RLA, c.ABX, 7))

	// 4
	lookup = append(lookup, addInstruction("RTI", c.RTI, c.IMP, 6))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ZP0, 3))
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

	// 5
	lookup = append(lookup, addInstruction("BVC", c.BVC, c.REL, 2))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ZPX, 4))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.ZPX, 4))
	lookup = append(lookup, addInstruction("LSR", c.LSR, c.ZPX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("CLI", c.CLI, c.IMP, 2))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.ABY, 4))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ABX, 4))
	lookup = append(lookup, addInstruction("EOR", c.EOR, c.ABX, 4))
	lookup = append(lookup, addInstruction("LSR", c.LSR, c.ABX, 7))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))

	// 6
	lookup = append(lookup, addInstruction("RTS", c.RTS, c.IMP, 6))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.IZX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ZP0, 3))
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

	// 7
	lookup = append(lookup, addInstruction("BVS", c.BVS, c.REL, 2))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 8))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ZPX, 4))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.ZPX, 4))
	lookup = append(lookup, addInstruction("ROR", c.ROR, c.ZPX, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("SEI", c.SEI, c.IMP, 2))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.ABY, 4))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ABX, 4))
	lookup = append(lookup, addInstruction("ADC", c.ADC, c.ABX, 4))
	lookup = append(lookup, addInstruction("ROR", c.ROR, c.ABX, 7))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 7))

	// 8
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMM, 2))
	lookup = append(lookup, addInstruction("STA", c.STA, c.IZX, 6))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMM, 2))
	lookup = append(lookup, addInstruction("SAX", c.SAX, c.IZX, 6))
	lookup = append(lookup, addInstruction("STY", c.STY, c.ZP0, 3))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ZP0, 3))
	lookup = append(lookup, addInstruction("STX", c.STX, c.ZP0, 3))
	lookup = append(lookup, addInstruction("SAX", c.SAX, c.ZP0, 3))
	lookup = append(lookup, addInstruction("DEY", c.DEY, c.IMP, 2))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMM, 2))
	lookup = append(lookup, addInstruction("TXA", c.TXA, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("STY", c.STY, c.ABS, 4))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ABS, 4))
	lookup = append(lookup, addInstruction("STX", c.STX, c.ABS, 4))
	lookup = append(lookup, addInstruction("SAX", c.SAX, c.ABS, 4))

	// 9
	lookup = append(lookup, addInstruction("BCC", c.BCC, c.REL, 2))
	lookup = append(lookup, addInstruction("STA", c.STA, c.IZY, 6))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 6))
	lookup = append(lookup, addInstruction("STY", c.STY, c.ZPX, 4))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ZPX, 4))
	lookup = append(lookup, addInstruction("STX", c.STX, c.ZPY, 4))
	lookup = append(lookup, addInstruction("SAX", c.SAX, c.ZPY, 4))
	lookup = append(lookup, addInstruction("TYA", c.TYA, c.IMP, 2))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ABY, 5))
	lookup = append(lookup, addInstruction("TXS", c.TXS, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("???", c.NOP, c.IMP, 5))
	lookup = append(lookup, addInstruction("STA", c.STA, c.ABX, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 5))

	// A
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.IMM, 2))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.IZX, 6))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.IMM, 2))
	lookup = append(lookup, addInstruction("LAX", c.LAX, c.IZX, 6))
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.ZP0, 3))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ZP0, 3))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.ZP0, 3))
	lookup = append(lookup, addInstruction("LAX", c.LAX, c.ZP0, 3))
	lookup = append(lookup, addInstruction("TAY", c.TAY, c.IMP, 2))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.IMM, 2))
	lookup = append(lookup, addInstruction("TAX", c.TAX, c.IMP, 2))
	lookup = append(lookup, addInstruction("LAX", c.LAX, c.ABS, 4))
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.ABS, 4))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ABS, 4))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.ABS, 4))
	lookup = append(lookup, addInstruction("LAX", c.LAX, c.ABS, 4))

	// B
	lookup = append(lookup, addInstruction("BCS", c.BCS, c.REL, 2))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("LAX", c.LAX, c.IZY, 5))
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.ZPX, 4))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ZPX, 4))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.ZPY, 4))
	lookup = append(lookup, addInstruction("LAX", c.LAX, c.ZPY, 4)) // added NOP instead of invalid? (EA?)
	lookup = append(lookup, addInstruction("CLV", c.CLV, c.IMP, 2))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ABY, 4))
	lookup = append(lookup, addInstruction("TSX", c.TSX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 4))
	lookup = append(lookup, addInstruction("LDY", c.LDY, c.ABX, 4))
	lookup = append(lookup, addInstruction("LDA", c.LDA, c.ABX, 4))
	lookup = append(lookup, addInstruction("LDX", c.LDX, c.ABY, 4))
	lookup = append(lookup, addInstruction("LAX", c.LAX, c.ABY, 4))

	// C
	lookup = append(lookup, addInstruction("CPY", c.CPY, c.IMM, 2))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.IZX, 6))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMM, 2))
	lookup = append(lookup, addInstruction("DCP", c.DCP, c.IZX, 8))
	lookup = append(lookup, addInstruction("CPY", c.CPY, c.ZP0, 3))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ZP0, 3))
	lookup = append(lookup, addInstruction("DEC", c.DEC, c.ZP0, 5))
	lookup = append(lookup, addInstruction("DCP", c.DCP, c.ZP0, 5))
	lookup = append(lookup, addInstruction("INY", c.INY, c.IMP, 2))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.IMM, 2))
	lookup = append(lookup, addInstruction("DEX", c.DEX, c.IMP, 2))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("CPY", c.CPY, c.ABS, 4))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ABS, 4))
	lookup = append(lookup, addInstruction("DEC", c.DEC, c.ABS, 6))
	lookup = append(lookup, addInstruction("DCP", c.DCP, c.ABS, 6))

	// D
	lookup = append(lookup, addInstruction("BNE", c.BNE, c.REL, 2))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("DCP", c.DCP, c.IZY, 8))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ZPX, 4))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ZPX, 4))
	lookup = append(lookup, addInstruction("DEC", c.DEC, c.ZPX, 6))
	lookup = append(lookup, addInstruction("DCP", c.DCP, c.ZPX, 6))
	lookup = append(lookup, addInstruction("CLD", c.CLD, c.IMP, 2))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ABY, 4))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("DCP", c.DCP, c.ABY, 7))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ABX, 4))
	lookup = append(lookup, addInstruction("CMP", c.CMP, c.ABX, 4))
	lookup = append(lookup, addInstruction("DEC", c.DEC, c.ABX, 7))
	lookup = append(lookup, addInstruction("DCP", c.DCP, c.ABX, 7))

	// E
	lookup = append(lookup, addInstruction("CPX", c.CPX, c.IMM, 2))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.IZX, 6))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMM, 2))
	lookup = append(lookup, addInstruction("ISB", c.ISB, c.IZX, 8))
	lookup = append(lookup, addInstruction("CPX", c.CPX, c.ZP0, 3))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ZP0, 3))
	lookup = append(lookup, addInstruction("INC", c.INC, c.ZP0, 5))
	lookup = append(lookup, addInstruction("ISB", c.ISB, c.ZP0, 5))
	lookup = append(lookup, addInstruction("INX", c.INX, c.IMP, 2))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.IMM, 2))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.IMM, 2))
	lookup = append(lookup, addInstruction("CPX", c.CPX, c.ABS, 4))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ABS, 4))
	lookup = append(lookup, addInstruction("INC", c.INC, c.ABS, 6))
	lookup = append(lookup, addInstruction("ISB", c.ISB, c.ABS, 6))

	// F
	lookup = append(lookup, addInstruction("BEQ", c.BEQ, c.REL, 2))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.IZY, 5))
	lookup = append(lookup, addInstruction("???", c.XXX, c.IMP, 2))
	lookup = append(lookup, addInstruction("ISB", c.ISB, c.IZY, 8))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ZPX, 4))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ZPX, 4))
	lookup = append(lookup, addInstruction("INC", c.INC, c.ZPX, 6))
	lookup = append(lookup, addInstruction("ISB", c.ISB, c.ZPX, 6))
	lookup = append(lookup, addInstruction("SED", c.SED, c.IMP, 2))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ABY, 4))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.IMP, 2))
	lookup = append(lookup, addInstruction("ISB", c.ISB, c.ABY, 7))
	lookup = append(lookup, addInstruction("NOP", c.NOP, c.ABX, 4))
	lookup = append(lookup, addInstruction("SBC", c.SBC, c.ABX, 4))
	lookup = append(lookup, addInstruction("INC", c.INC, c.ABX, 7))
	lookup = append(lookup, addInstruction("ISB", c.ISB, c.ABX, 7))

	return lookup
}
