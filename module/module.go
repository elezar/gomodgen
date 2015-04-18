package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"time"

	"github.com/elezar/gomodgen/generic"
	"github.com/elezar/gomodgen/interfaces"
)

var o Outputer

type Module struct {
	Desc     string
	Name     string
	entities []interfaces.Entity
}

type ModuleLoader struct {
	Desc     string
	Name     string
	Generics []string
}

// Load the module representation from a json file.
func Load(filename string) *Module {
	var mfile ModuleLoader

	// Load the json file
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return nil
	}

	// Marshal the JSON data.

	err = json.Unmarshal(jsonData, &mfile)
	if err != nil {
		fmt.Printf("Error reading json file %s: %v\n", filename, err)
		return nil
	}

	// Set the properties of the Module structure
	var m *Module = new(Module)
	m.Desc = mfile.Desc
	m.Name = mfile.Name

	m.entities = make([]interfaces.Entity, 0, len(mfile.Generics))

	usepath := filepath.Dir(filename)

	for _, g := range mfile.Generics {
		m.Add(generic.NewFromFile(path.Join(usepath, g)))
	}

	return m
}

// Generate generates the source for the module.
func (m Module) Generate() string {
	// Add the description as a line comment.
	o.Add(m.Description())
	o.Add("module " + m.Name)

	// indent++
	o.Add("implicit none")
	o.Add("private")

	// Add the declartion for the module.
	o.Add(m.Declaration())

	o.Add("contains")

	// Add the module body.
	o.Add(m.Definition())

	// indent--
	o.Add("end module " + m.Name)
	o.Add("", 0)

	return o.String()
}

// Description returns the text description for the module.
func (m Module) Description() string {
	var s Outputer
	s.AddComment("Automatically generated on " + fmt.Sprint(time.Now()))
	s.AddBlankLine()
	s.AddComment(m.Desc)

	return s.String()
}

// Declaration returns the declaration part of the module. This includes
// Any generic interfaces.
func (m Module) Declaration() string {

	var s Outputer

	for _, e := range m.entities {
		s.AddComment(e.Description())
		s.Add(e.Declaration())
	}

	return s.String()
}

// Definition returns the defintion part (body) of the module. This includes
// the bodies of any specific implementations of generic interfaces.
func (m Module) Definition() string {

	var s Outputer

	for _, e := range m.entities {
		s.Add(e.Definition())
	}

	return s.String()
}

// Add an entity to the module.
func (m *Module) Add(e interfaces.Entity) {
	m.entities = append(m.entities, e)
}
