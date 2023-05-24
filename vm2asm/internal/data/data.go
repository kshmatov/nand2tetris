package data

import (
	"errors"
)

type cType int
type sType int 
type addr uint16


const (
	cPush = CType(0) iota
	cPop
	cArithmetic
	cLabel
	cGoto
	cIf
	cFunction
	cReturn
	cCall

	sStack  = sType(0) iota
	sLocal
	sArgs
	sThis
	sThat
	sConst
	sStatic
	sTemp
	sPointer

	sp = addr(0) iota
	lcl 
	arg
	this
	that
)

var (
	ErrMemSegmentUndefined = errors

	memSegments = map[sType]addr {
		sStack: 256,
		sStatic: 16,
		sTemp: 5,
	}
)

type Command interface {
	Out() ([]string, error)
}

type push struct {
	comment string
	segment sStack
	addr uint16
}

func (p *push) Out()([]string, error) {
	segAddr, ok := memSegments[p.segment]
	if ! ok {
		return nil, ErrMemSegmentUndefined
	}
	return []string {
		"// " +comment,
		""
	}
}

type pop struct {
	segment sStack
	addr uint16
}



