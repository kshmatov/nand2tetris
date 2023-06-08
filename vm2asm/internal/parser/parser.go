package parser

import (
	"fmt"
	"strings"

	"github.com/kshmatov/nand2tetris/vm2asm/internal/command"
	"github.com/kshmatov/nand2tetris/vm2asm/internal/parser/data"
)

func Parse(vmCode []string, name string) []command.Command {
	data.SetFileName(name)

	res := make([]command.Command, 0, len(vmCode)+1)
	for _, line := range vmCode {
		fmt.Println(line)
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		l := strings.Split(line, "//")
		if l[0] == "" {
			continue
		}
		l = strings.Split(line, " ")
		res = append(res, data.New(l[0], l[1:]...))
	}
	return res
}

func End() string {
	d, _ := (&data.End{}).Out()
	return strings.Join(d, "\n")
}

func Head() string {
	d, _ := (&data.Head{}).Out()
	return strings.Join(d, "\n")
}
