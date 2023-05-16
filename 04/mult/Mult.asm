// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// Put your code here.

    @R0
    D=M
    @mul
    M=D
    @R1
    D=M
    @cnt
    M=D
    @i
    M=0
    @res
    M=0

    @mul
    D=M
    @DONE
    D;JEQ

    @cnt
    D=M
    @DONE
    D;JEQ

(LOOP)
    @cnt
    D=M
    @i
    D=D-M
    @DONE
    D;JLE

    @res
    D=M
    @mul
    D=D+M
    @res
    M=D

    @i
    D=M
    D=D+1
    M=D

    @LOOP
    0;JMP

(DONE)
    @res
    D=M
    @R2
    M=D

(END)
    @END
    0;JMP
