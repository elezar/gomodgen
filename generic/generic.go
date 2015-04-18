package generic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/elezar/gomodgen/impl"
)

// Generator allows for the generation of the Name and Body of a specific implemenation
// of the generic interface
type Generator interface {
	Name(basename, typename string, dim int) string
	Body(basename, typename string, dim int) string
}

// Generic represents a Fortran generic interface.
type Generic struct {
	Name       string
	BodyFile   string
	Types      []string
	Dimensions []int
	def        impl.Impl
}

// Load the generic representation from a file
func (g *Generic) Load(filename string) error {
	g.Name = ""
	g.BodyFile = ""
	g.Types = []string{}
	g.Dimensions = []int{}

	// Load the json file
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return err
	}

	err = json.Unmarshal(jsonData, g)
	if err != nil {
		fmt.Printf("Error reading json file %s: %v\n", filename, err)
		return err
	}

	if len(g.Dimensions) == 0 {
		g.Dimensions = append(g.Dimensions, 0)
	}

	folder := filepath.Dir(filename)
	// Set the properties for the definition
	g.def.Basename = g.Name
	g.def.LoadBody(path.Join(folder, g.BodyFile))

	return nil
}

// Declaration returns the generic interface declaration defined by the structure
func (g Generic) Declaration() string {

	// Ensure that the name is set.
	g.def.Basename = g.Name

	s := "interface " + g.Name + "\n"
	for _, t := range g.Types {
		for _, d := range g.Dimensions {
			g.def.Typename = t
			g.def.Dimension = d
			s += "module procedure " + g.def.Name() + "\n"
		}
	}

	s += "end interface"
	return s
}

// Definition returns the specific implementations of the generic interface
func (g Generic) Definition() string {

	s := "\n"
	for _, t := range g.Types {
		for _, d := range g.Dimensions {
			g.def.Typename = t
			g.def.Dimension = d
			s += g.def.Definition() + "\n"
		}
	}

	return s
}
