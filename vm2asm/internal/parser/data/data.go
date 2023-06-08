package data

import (
	"fmt"
	"strconv"

	"github.com/kshmatov/nand2tetris/vm2asm/internal/command"
	"github.com/kshmatov/nand2tetris/vm2asm/internal/parser/data/counter"
	"github.com/pkg/errors"
)

const (
	cPush     = "push"
	cPop      = "pop"
	cAdd      = "add"
	cSub      = "sub"
	cNeg      = "neg"
	cEq       = "eq"
	cGt       = "gt"
	cLt       = "lt"
	cAnd      = "and"
	cOr       = "or"
	cNot      = "not"
	cLabel    = "label"
	cGoto     = "goto"
	cIf       = "if"
	cFunction = "function"
	cReturn   = "return"
	cCall     = "call"

	sStack   = "stack"
	sLocal   = "local"
	sArgs    = "argument"
	sThis    = "this"
	sThat    = "that"
	sConst   = "constant"
	sStatic  = "static"
	sTemp    = "temp"
	sPointer = "pointer"

	sp   = uint16(0)
	lcl  = uint16(1)
	arg  = uint16(2)
	this = uint16(3)
	that = uint16(4)
)

var (
	ErrMemSegmentUndefined = errors.New("segment undefined")
	ErrPopToConst          = errors.New("can't update constant")

	fileName = ""

	memSegments = map[string]uint16{
		sStack:  256,
		sStatic: 16,
		sTemp:   5,
	}

	pToSegments = map[string]string{
		sLocal: "LCL",
		sArgs:  "ARG",
		sThis:  "THIS",
		sThat:  "THAT",
	}

	// set M to top value on stack
	popAsm = []string{
		"@SP",
		"M=M-1",
		"A=M",
	}

	// push value from D to stack
	pushAsm = []string{
		// get pointer to stack head and add value from D to the top
		"@SP",
		"A=M",
		"M=D",
		// inc stack pointer
		"@SP",
		"M=M+1",
	}
)

func SetFileName(s string) {
	fileName = s
}

func New(op string, data ...string) command.Command {
	fmt.Println(op, data)
	switch op {
	case cPop:
		return &pop{
			comment: op + " " + data[0] + " " + data[1],
			segment: data[0],
			addr:    data[1],
		}
	case cPush:
		return &push{
			comment: op + " " + data[0] + " " + data[1],
			segment: data[0],
			addr:    data[1],
		}
	case cAdd, cSub, cEq, cGt, cLt, cAnd, cOr:
		return &twoOperands{
			op: op,
		}
	case cNeg, cNot:
		return &oneOperand{
			op: op,
		}
	case cLabel:
		return &label{
			name: data[0],
		}
	case cGoto:
		return &gotoCmd{
			label: data[0],
		}
	case cIf:
		return &ifGoto{
			label: data[0],
		}
	case cFunction:
		i, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}
		return &function{
			name: data[0],
			args: i,
		}
	case cCall:
		i, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}
		return &call{
			name: data[0],
			args: i,
		}
	case cReturn:
		return &ret{}
	}
	return nil
}

type End struct {
}

func (e *End) Out() ([]string, error) {
	m := fileName + "$END" + counter.Get()
	return []string{
		"// this is the end...",
		"(" + m + ")",
		"@" + m,
		"0;JMP",
	}, nil
}

type Head struct {
}

func (h *Head) Out() ([]string, error) {
	res := []string{
		"@256",
		"D=A",
		"@SP",
		"M=D",
	}
	c := call{
		name: "Sys.init",
		args: 0,
	}
	cout, _ := c.Out()
	res = append(res, cout...)
	return res, nil
}
