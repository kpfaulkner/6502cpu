# 6502 emulator

Just for fun.
This is 95% based off https://www.youtube.com/watch?v=8XmxKPJDGU0
No original thought from me :L)

## Notes

The nestest completes (https://github.com/christopherpow/nes-test-roms/blob/master/other/nestest.nes) as in all operations
in the correct order are executed. Checks for the registers/flags not performed yet (but the fact the ops complete is a good 
indication)

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


For comparison instructions, the results are usually found in flags:
Compare Result	    N	Z	C
A, X, or Y < Memory	*	0	0
A, X, or Y = Memory	0	1	1
A, X, or Y > Memory	*	0	1



## See

https://www.youtube.com/watch?v=8XmxKPJDGU0
https://www.masswerk.at/6502/6502_instruction_set.html

http://www.6502.org/tutorials/compare_instructions.html
