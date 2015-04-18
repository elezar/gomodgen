package generic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/elezar/gomodgen/impl"
	"github.com/elezar/gomodgen/interfaces"
)

type GenericLoader struct {
	Desc       string
	Name       string
	BodyFile   string
	Types      []string
	Dimensions []int
}

// Generic represents a Fortran generic interface.
type Generic struct {
	Desc     string
	Name     string
	entities []interfaces.Entity
}

func NewFromFile(filename string) *Generic {
	return Load(filename)
}

func New() *Generic {
	return &Generic{}
}

// Load the generic representation from a json file.
func Load(filename string) *Generic {
	var gfile GenericLoader
	// Load the json file
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return nil
	}

	// Marshal the JSON data.
	err = json.Unmarshal(jsonData, &gfile)
	if err != nil {
		fmt.Printf("Error reading json file %s: %v\n", filename, err)
		return nil
	}

	if len(gfile.Dimensions) == 0 {
		gfile.Dimensions = append(gfile.Dimensions, 0)
	}

	var g *Generic = new(Generic)

	bodyfile := path.Join(filepath.Dir(filename), gfile.BodyFile)

	// Set the properties of the Generic structure.
	g.Desc = gfile.Desc
	g.Name = gfile.Name
	g.entities = make([]interfaces.Entity, 0, len(gfile.Types)*len(gfile.Dimensions))

	for _, t := range gfile.Types {
		for _, d := range gfile.Dimensions {
			g.Add(impl.NewFromFile(g.Name, t, d, bodyfile))
		}
	}

	return g
}

func (g Generic) Description(o interfaces.Outputer) {
	o.AddComment(g.Desc)
}

// Declaration returns the generic interface declaration defined by the structure
func (g Generic) Declaration(o interfaces.Outputer) {

	o.Add("interface " + g.Name)
	o.Indent()
	for _, e := range g.entities {
		e.Declaration(o)
	}
	o.Deindent()
	o.Add("end interface")

}

// Definition returns the specific implementations of the generic interface
func (g Generic) Definition(o interfaces.Outputer) {

	o.Add("")
	for _, e := range g.entities {
		e.Definition(o)
	}
}

// Add an entity to the generic interface.
func (g *Generic) Add(e interfaces.Entity) {
	g.entities = append(g.entities, e)
}
