package data

type label struct {
	name string
}

func (l *label) Out() ([]string, error) {
	return []string{
		"// label " + l.name,
		"(" + fileName + "$" + l.name + ")",
	}, nil
}

type gotoCmd struct {
	label string
}

func (g *gotoCmd) Out() ([]string, error) {
	return []string{
		"// goto " + g.label,
		"@" + fileName + "$" + g.label,
		"0;JMP",
	}, nil
}

type ifGoto struct {
	label string
}

func (g *ifGoto) Out() ([]string, error) {
	res := []string{
		"// if-goto " + g.label,
	}
	res = append(res, popAsm...)
	res = append(res,
		"D=M",
		"@"+fileName+"$"+g.label,
		"D; JNE",
	)
	return res, nil
}
