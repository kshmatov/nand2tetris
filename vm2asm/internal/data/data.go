package data

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Command interface {
	Out() ([]string, error)
}

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
	cFunction = "?"
	cReturn   = "return"
	cCall     = "?"

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

	labelCnt = 0

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

func New(op string, data ...string) Command {
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
	}
	return nil
}

type End struct {
}

func (e *End) Out() ([]string, error) {
	m := "END" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	return []string{
		"// this is the end...",
		"(" + m + ")",
		"@" + m,
		"0;JMP",
	}, nil
}
