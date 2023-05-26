package data

import (
	"strconv"
	"time"
)

var (
	opsAsm = map[string][]string{
		cAdd: {"M=M+D"},
		cSub: {"M=M-D"},
		cNeg: {"M=-M"},
		cEq:  {"0;JEQ"},
		cGt:  {"0;JGT"},
		cLt:  {"0;JLT"},
		cAnd: {"M=M&D"},
		cOr:  {"M=M|D"},
		cNot: {"M=!M"},
	}
)

func getOps(op string) []string {
	t := strconv.Itoa(time.Now().Nanosecond())
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
		return []string{
			"@TRUE" + t,
			"M-D;JLT",
			"M=0",
			"@END" + t,
			"0;" + jmp,
			"(TRUE" + t + ")",
			"M=-1",
			"@END" + t,
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
