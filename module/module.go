package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/elezar/gomodgen/generic"
	"github.com/elezar/gomodgen/interfaces"
)

const (
	commentTag = "! "
	tab        = "  "
)

type lines []string

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
	var mLines lines

	indent := 0

	// Add the description as a line comment.
	mLines.Add(m.Description(), indent)
	mLines.Add("module "+m.Name, indent)

	// indent++
	mLines.Add("implicit none", indent)
	mLines.Add("private", indent)

	// Add the declartion for the module.
	mLines.Add(m.Declaration(), indent)

	mLines.Add("contains", indent)

	// Add the module body.
	mLines.Add(m.Definition(), indent)

	// indent--
	mLines.Add("end module "+m.Name, indent)
	mLines.Add("", 0)

	s := strings.Join(mLines, "\n")

	return s
}

// Description returns the text description for the module.
func (m Module) Description() string {
	var s lines
	s.Add(commentTag+"Automatically generated on "+fmt.Sprint(time.Now()), 0)
	s.Add("", 0)
	s.Add(commentTag+m.Desc, 0)

	return strings.Join(s, "\n")
}

// Declaration returns the declaration part of the module. This includes
// Any generic interfaces.
func (m Module) Declaration() string {

	var s []string

	for _, e := range m.entities {
		s = append(s, e.Declaration())
	}

	return strings.Join(s, "\n")
}

// Definition returns the defintion part (body) of the module. This includes
// the bodies of any specific implementations of generic interfaces.
func (m Module) Definition() string {

	var s []string

	for _, e := range m.entities {
		s = append(s, e.Definition())
	}

	return strings.Join(s, "\n")
}

// Add an entity to the module.
func (m *Module) Add(e interfaces.Entity) {
	m.entities = append(m.entities, e)
}

// Add adds a newLine to a set of lines using the specified indent.
func (l *lines) Add(newLine string, indent int) {
	var s string

	for i := 0; i < indent; i++ {
		s += tab
	}

	*l = append(*l, s+newLine)
}
