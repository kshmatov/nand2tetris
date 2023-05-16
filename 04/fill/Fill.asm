// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

(LOOP)
    @pos
    M=0
    @KBD
    D=M
    @filler
    M=D

(FILLSCREEN)
    @8192 // Size of the screenmap 256*32
    D=A
    @pos
    D=D-M
    @END
    D;JLE

    @filler
    D=M
    @WHITE
    D;JEQ

// BLACK
    @SCREEN
    D=A
    @pos
    A=D+M
    M=-1
    @DONE
    0;JMP

(WHITE)
    @SCREEN
    D=A
    @pos
    A=D+M
    M=0

(DONE)

    @pos
    M=M+1
    @FILLSCREEN
    0;JMP

(END)
    @LOOP
    0;JMP