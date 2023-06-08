package data

import (
	"strconv"

	"github.com/pkg/errors"
)

type push struct {
	comment string
	segment string
	addr    string
}

func addr(s string) uint16 {
	v, _ := strconv.Atoi(s)
	return uint16(v)
}

func (p *push) Out() ([]string, error) {
	selector := p.segment
	if p.segment == sPointer {
		if p.addr == "0" {
			selector = sThis
		} else {
			selector = sThat
		}
	}

	addrStr := baseAddr(selector, p.addr)
	switch selector {
	case sLocal, sArgs, sThis, sThat:
		addrStr = append(addrStr, "D=M")
	case sConst:
		addrStr = append(addrStr, "D=A")
	case sTemp, sStatic:
		addrStr = append(addrStr, "D=M")
	}

	res := []string{
		"// " + p.comment,
	}
	// get value from memory segment
	res = append(res,
		addrStr...)

	res = append(res, pushAsm...)

	return res, nil
}

type pop struct {
	comment string
	segment string
	addr    string
}

func (p *pop) Out() ([]string, error) {
	if p.segment == sConst {
		return nil, errors.Wrap(ErrPopToConst, p.comment)
	}
	selector := p.segment
	if p.segment == sPointer {
		if p.addr == "0" {
			selector = sThis
		} else {
			selector = sThat
		}
	}

	addrStr := baseAddr(selector, p.addr)

	res := []string{
		"// " + p.comment,
	}

	res = append(res, addrStr...)
	res = append(res, "D=A", "@5", "M=D")
	res = append(res, popAsm...)
	res = append(res,
		"D=M", "@5", "A=M", "M=D",
	)
	return res, nil
}

func baseAddr(selector, shift string) []string {
	switch selector {
	case sLocal, sArgs, sThis, sThat:
		base := pToSegments[selector]
		return []string{
			"@" + base,
			"D=M",
			"@" + shift,
			"A=A+D",
		}
	case sConst:
		return []string{
			"@" + shift,
		}
	case sTemp, sStatic:
		base := memSegments[selector] + addr(shift)
		return []string{
			"@" + strconv.Itoa(int(base)),
		}
	}
	return nil
}
