// push constant 17
@SP
A=M
M=D
@SP
M=M+1
A=@17
D=A
// push constant 17
@SP
A=M
M=D
@SP
M=M+1
A=@17
D=A
// eq
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
@TRUE563039700
M-D;JLT
M=0
@END563039700
0;JEQ
(TRUE563039700)
M=-1
@END563039700
@SP
M=M+1
// push constant 17
@SP
A=M
M=D
@SP
M=M+1
A=@17
D=A
// push constant 16
@SP
A=M
M=D
@SP
M=M+1
A=@16
D=A
// eq
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
@TRUE563566000
M-D;JLT
M=0
@END563566000
0;JEQ
(TRUE563566000)
M=-1
@END563566000
@SP
M=M+1
// push constant 16
@SP
A=M
M=D
@SP
M=M+1
A=@16
D=A
// push constant 17
@SP
A=M
M=D
@SP
M=M+1
A=@17
D=A
// eq
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
@TRUE564092500
M-D;JLT
M=0
@END564092500
0;JEQ
(TRUE564092500)
M=-1
@END564092500
@SP
M=M+1
// push constant 892
@SP
A=M
M=D
@SP
M=M+1
A=@892
D=A
// push constant 891
@SP
A=M
M=D
@SP
M=M+1
A=@891
D=A
// lt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
@TRUE564092500
M-D;JLT
M=0
@END564092500
0;JLT
(TRUE564092500)
M=-1
@END564092500
@SP
M=M+1
// push constant 891
@SP
A=M
M=D
@SP
M=M+1
A=@891
D=A
// push constant 892
@SP
A=M
M=D
@SP
M=M+1
A=@892
D=A
// lt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
@TRUE564092500
M-D;JLT
M=0
@END564092500
0;JLT
(TRUE564092500)
M=-1
@END564092500
@SP
M=M+1
// push constant 891
@SP
A=M
M=D
@SP
M=M+1
A=@891
D=A
// push constant 891
@SP
A=M
M=D
@SP
M=M+1
A=@891
D=A
// lt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
@TRUE564615400
M-D;JLT
M=0
@END564615400
0;JLT
(TRUE564615400)
M=-1
@END564615400
@SP
M=M+1
// push constant 32767
@SP
A=M
M=D
@SP
M=M+1
A=@32767
D=A
// push constant 32766
@SP
A=M
M=D
@SP
M=M+1
A=@32766
D=A
// gt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
@TRUE564615400
M-D;JLT
M=0
@END564615400
0;JGT
(TRUE564615400)
M=-1
@END564615400
@SP
M=M+1
// push constant 32766
@SP
A=M
M=D
@SP
M=M+1
A=@32766
D=A
// push constant 32767
@SP
A=M
M=D
@SP
M=M+1
A=@32767
D=A
// gt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
@TRUE564615400
M-D;JLT
M=0
@END564615400
0;JGT
(TRUE564615400)
M=-1
@END564615400
@SP
M=M+1
// push constant 32766
@SP
A=M
M=D
@SP
M=M+1
A=@32766
D=A
// push constant 32766
@SP
A=M
M=D
@SP
M=M+1
A=@32766
D=A
// gt
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
@TRUE564615400
M-D;JLT
M=0
@END564615400
0;JGT
(TRUE564615400)
M=-1
@END564615400
@SP
M=M+1
// push constant 57
@SP
A=M
M=D
@SP
M=M+1
A=@57
D=A
// push constant 31
@SP
A=M
M=D
@SP
M=M+1
A=@31
D=A
// push constant 53
@SP
A=M
M=D
@SP
M=M+1
A=@53
D=A
// add
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=M+D
@SP
M=M+1
// push constant 112
@SP
A=M
M=D
@SP
M=M+1
A=@112
D=A
// sub
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=M-D
@SP
M=M+1
// neg
@SP
M=M-1
A=M
M=-M
@SP
M=M+1
// and
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=M&D
@SP
M=M+1
// push constant 82
@SP
A=M
M=D
@SP
M=M+1
A=@82
D=A
// or
@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=M|D
@SP
M=M+1
// not
@SP
M=M-1
A=M
M=!M
@SP
M=M+1
