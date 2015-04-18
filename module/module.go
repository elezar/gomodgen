package module

import (
	"fmt"
	"strings"
	"time"
)

const (
	commentTag = "! "
	tab        = "  "
)

type lines []string

type Entity interface {
	Declaration() string
	Definition() string
}

type Module struct {
	Desc     string
	Name     string
	Entities []Entity
}

// Generate generates the source for the module.
func (m Module) Generate() string {
	var mLines lines

	indent := 0

	// Add the description as a line comment.
	mLines.Add(m.Description(), indent)
	mLines.Add("module "+m.Name, indent)

	indent++
	mLines.Add("implicit none", indent)
	mLines.Add("private", indent)

	// Add the declartion for the module.
	mLines.Add(m.Declaration(), indent)

	mLines.Add("contains", indent-1)

	// Add the module body.
	mLines.Add(m.Definition(), indent)

	indent--
	mLines.Add("end module "+m.Name, indent)

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

	for _, e := range m.Entities {
		s = append(s, e.Declaration())
	}

	return strings.Join(s, "\n")
}

// Definition returns the defintion part (body) of the module. This includes
// the bodies of any specific implementations of generic interfaces.
func (m Module) Definition() string {

	var s []string

	for _, e := range m.Entities {
		s = append(s, e.Definition())
	}

	return strings.Join(s, "\n")
}

// Add an entity to the module.
func (m *Module) Add(e Entity) {
	m.Entities = append(m.Entities, e)
}

// Add adds a newLine to a set of lines using the specified indent.
func (l *lines) Add(newLine string, indent int) {
	var s string

	for i := 0; i < indent; i++ {
		s += tab
	}

	*l = append(*l, s+newLine)
}
