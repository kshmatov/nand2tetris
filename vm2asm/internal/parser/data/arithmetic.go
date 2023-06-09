package data

import (
	"github.com/kshmatov/nand2tetris/vm2asm/internal/parser/data/counter"
)

var (
	opsAsm = map[string][]string{
		cAdd: {"M=M+D"},
		cSub: {"M=M-D"},
		cNeg: {"M=-M"},
		cAnd: {"M=M&D"},
		cOr:  {"M=M|D"},
		cNot: {"M=!M"},
	}
)

func getOps(op string) []string {

	switch op {
	case cAdd, cSub, cNeg, cAnd, cOr, cNot:
		return opsAsm[op]
	case cEq, cLt, cGt:
		jmp := "JEQ"
		switch op {
		case cLt:
			jmp = "JLT"
		case cGt:
			jmp = "JGT"
		}
		t := counter.Get()
		return []string{
			"D=M-D",
			"@" + fileName + "$TRUE" + t,
			"D;JLT",
			"@SP",
			"A=M",
			"M=0",
			"@" + fileName + "$END" + t,
			"0;" + jmp,
			"(" + fileName + "$TRUE" + t + ")",
			"@SP",
			"A=M",
			"M=-1",
			"(" + fileName + "$END" + t + ")",
		}
	}
	return nil
}

type oneOperand struct {
	op string
}

func (o *oneOperand) Out() ([]string, error) {
	res := []string{"// " + o.op}
	res = append(res, popAsm...)
	res = append(res, opsAsm[o.op]...)
	res = append(res,
		"@SP",
		"M=M+1",
	)
	return res, nil
}

type twoOperands struct {
	op string
}

func (o *twoOperands) Out() ([]string, error) {
	res := []string{"// " + o.op}
	res = append(res, popAsm...)
	res = append(res,
		"D=M",
	)
	res = append(res, popAsm...)
	res = append(res, getOps(o.op)...)
	res = append(res,
		"@SP",
		"M=M+1",
	)
	return res, nil

}
