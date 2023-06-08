package data

import (
	"strconv"

	"github.com/kshmatov/nand2tetris/vm2asm/internal/parser/data/counter"
)

type function struct {
	name string
	args int
}

func (f *function) Out() ([]string, error) {
	res := []string{
		"// function" + f.name + " " + strconv.Itoa(f.args),
		"(" + f.name + ")",
		"D=0",
	}

	for i := 0; i < int(f.args); i++ {
		res = append(res, pushAsm...)
	}
	return res, nil
}

type ret struct {
}

func (r *ret) Out() ([]string, error) {
	frame := "R13"
	retAddr := "R14"

	res := []string{
		"// return",
		"@LCL",
		"D=M",
		"@" + frame,
		"M=D",

		"@" + frame,
		"D=M",
		"@5",
		"A=D-A",
		"D=M",
		"@" + retAddr,
		"M=D",
	}

	res = append(res, popAsm...)

	res = append(res,
		"D=M",
		"@ARG",
		"A=M",
		"M=D",

		"@ARG",
		"D=M",
		"@SP",
		"M=D+1")

	offset := 1
	for i, addr := range []string{"@THAT", "@THIS", "@ARG", "@LCL"} {
		res = append(res,
			"@"+frame,
			"D=M",
			"@"+strconv.Itoa(i+offset),
			"A=D-A",
			"D=M",
			addr,
			"M=D",
		)
	}

	res = append(res,
		"@"+retAddr,
		"A=M",
		"0;JMP",
	)

	return res, nil
}

type call struct {
	name string
	args int
}

func (c *call) Out() ([]string, error) {
	retAddr := c.name + "RET" + counter.Get()

	res := []string{
		"// call " + c.name + " " + strconv.Itoa(c.args),
		"@" + retAddr,
		"D=A",
	}

	res = append(res, pushAsm...)
	for _, addr := range []string{"@LCL", "@ARG", "@THIS", "@THAT"} {
		res = append(res,
			addr,
			"D=M",
		)
		res = append(res, pushAsm...)
	}

	res = append(res,
		"@SP",
		"D=M",
		"@LCL",
		"M=D",

		"@"+strconv.Itoa(c.args+5),
		"D=D+A",
		"@ARG",
		"M=D",

		"@"+c.name,
		"0;JMP",
		"("+retAddr+")",
	)

	return res, nil
}
