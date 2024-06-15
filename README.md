# 6502 emulator

Just for fun.
This is ALL based off https://www.youtube.com/watch?v=8XmxKPJDGU0
No original thought from me :L)

## Notes

16 bit address bus
8 bit data bus

Registers

A: Accumulator
X: Register
Y: Register
(above 3 are 8 bit)

stkp: Stack pointer : 8 bit points to memory location
pc: Program counter : 16 bit. Address of next byte to read for program
status: status register  :  (last operation resulted in zero, disable interrupts etc. Each of these flags is a single bit)




## See

https://www.youtube.com/watch?v=8XmxKPJDGU0
https://www.masswerk.at/6502/6502_instruction_set.html