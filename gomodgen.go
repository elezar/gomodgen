package main

import (
	"flag"
	"fmt"

	"github.com/elezar/gomodgen/generic"
)

func generate(outfile string, infile string) error {
	var g generic.Generic
	var err error

	err = g.Load(infile)
	if err != nil {
		return err
	}

	fmt.Print(g.Declaration())
	fmt.Print(g.Definition())
	return nil
}

func main() {
	var input, output string

	flag.StringVar(&input, "input", "", "Input filename")
	flag.StringVar(&output, "output", "", "Ouptut filename")

	flag.Parse()

	if input == "" || output == "" {
		flag.Usage()
	}

	generate(output, input)
}
