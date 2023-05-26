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

	cnt, err := os.ReadFile(*in)
	if err != nil {
		log.Fatal(err)
	}
	src := strings.Split(string(cnt), "\n")
	res := parser.Parse(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	var df io.Writer
	if *out == "" {
		df = os.Stdout
	} else {
		f, err := os.OpenFile(*out, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		df = f
	}

	for _, command := range res {
		asm, err := command.Out()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, s := range asm {
			_, err := df.Write([]byte(s + "\n"))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Println("Done")
}
