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

	var o *Outputer = new(Outputer)

	// Add the description as a line comment.
	m.Description(o)
	o.Add("module " + m.Name)

	o.Indent()
	o.Add("implicit none")
	o.Newline()
	o.Add("private")
	o.Newline()

	// Add the declartion for the module.
	m.Declaration(o)

	// Contains
	o.Newline()
	o.Deindent()
	o.Add("contains")

	// Add the module body.
	m.Definition(o)

	// End module
	o.AddLine()
	o.Add("end module " + m.Name)
	o.Newline()

	return o.String()
}

// Description returns the text description for the module.
func (m Module) Description(o *Outputer) {
	o.AddComment("Automatically generated on " + fmt.Sprint(time.Now()))
	o.Newline()
	o.AddComment(m.Desc)
}

// Declaration returns the declaration part of the module. This includes
// Any generic interfaces.
func (m Module) Declaration(o *Outputer) {

	for _, e := range m.entities {
		e.Description(o)
		e.Declaration(o)
	}

}

// Definition returns the defintion part (body) of the module. This includes
// the bodies of any specific implementations of generic interfaces.
func (m Module) Definition(o *Outputer) {

	for _, e := range m.entities {
		e.Definition(o)
	}

}

// Add an entity to the module.
func (m *Module) Add(e interfaces.Entity) {
	m.entities = append(m.entities, e)
}
