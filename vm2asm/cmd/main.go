package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kshmatov/nand2tetris/vm2asm/internal/parser"
)

func main() {
	out := flag.String("o", "", "output file, defualt stdout")
	in := flag.String("i", "", "source file")
	flag.Parse()
	if *in == "" {
		fmt.Printf("no source file is given\n")
		flag.Usage()
		return
	}

	fs, err := os.Stat(*in)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tempFs := *out

	fns := []string{}

	if fs.IsDir() {
		tempFs = fs.Name() + ".asm"
		dirData, err := os.ReadDir(*in)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, f := range dirData {
			if f.Type().IsRegular() && strings.HasSuffix(f.Name(), ".vm") {
				fns = append(fns, *in+string(os.PathSeparator)+f.Name())
			}
		}
	} else {
		parts := strings.Split(*in, ".")
		parts[len(parts)-1] = "asm"
		tempFs = strings.Join(parts, ".")
	}

	if *out == "" {
		*out = tempFs
	}
	f, err := os.OpenFile(*out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.Write([]byte(parser.Head()))
	for _, fn := range fns {
		parts := strings.Split(fn, string(os.PathSeparator))
		name := strings.Split(parts[len(parts)-1], ".")[0]

		cnt, err := os.ReadFile(fn)
		if err != nil {
			log.Fatal(err)
		}
		src := strings.Split(string(cnt), "\n")

		build(src, f, name)
	}
	f.Write([]byte(parser.End()))

	fmt.Println("Done")
}

func build(src []string, out io.Writer, name string) error {
	res := parser.Parse(src, name)
	for _, command := range res {
		asm, err := command.Out()
		if err != nil {
			return err
		}
		for _, s := range asm {
			_, err := out.Write([]byte(s + "\n"))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
