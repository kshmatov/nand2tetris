package parser

import (
	"strings"

	"github.com/kshmatov/nand2tetris/vm2asm/internal/data"
)

func Parse(vmCode []string) []data.Command {
	res := make([]data.Command, 0, len(vmCode))
	for _, line := range vmCode {
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
