package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/elezar/gomodgen/module"
)

func generate(outfile string, infile string) error {
	var m *module.Module

	m = module.Load(infile)
	if m == nil {
		return errors.New("Could not create module from " + infile)
	}

	fmt.Print(m.Generate())
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
