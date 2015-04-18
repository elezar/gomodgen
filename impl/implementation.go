package impl

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/elezar/gomodgen/interfaces"
)

// Constants representing the tags in the implementation body.
const (
	basetag = "{{basename}}"
	typetag = "{{type}}"
	nametag = "{{name}}"
	ntag    = "{{n}}"
)

// The replacer is used to replace the occurrences of the tags in the implementation
// representation.
func (i Impl) newReplacer() *strings.Replacer {
	r := strings.NewReplacer(
		typetag, i.Typename,
		nametag, i.Name(),
		ntag, n(i.Dimension),
		basetag, i.Basename)
	return r
}

// Impl represents a specific function or subroutine implementation for a given type and
// dimension. This allows for the creation of a specific implementation of a generic
// interface. Certain tags (e.g. {{type}}, {{name}}) are replaced for the final string
// representation.
type Impl struct {
	Basename     string
	BodyTemplate string
	Typename     string
	Dimension    int
}

func New(basename string, bodyTemplate string, typename string, dim int) *Impl {
	return &Impl{Basename: basename,
		BodyTemplate: bodyTemplate,
		Typename:     typename,
		Dimension:    dim,
	}
}

// Create a new implementation, loading the body from a file.
func NewFromFile(Basename string, typename string, dim int, filename string) *Impl {
	i := New(Basename, "", typename, dim)

	i.LoadBody(filename)
	return i
}

// LoadBody loads the body of an Impl from the file specified
func (i *Impl) LoadBody(filename string) error {

	i.BodyTemplate = ""

	dataIn, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filename, err)
		return err
	}

	i.BodyTemplate = strings.TrimSpace(string(dataIn))

	return nil
}

func (i Impl) Description(o interfaces.Outputer) {
	r := i.newReplacer()
	var s string
	s = "{{name}} implements {{basename}} for a {{n}}-d {{type}} parameter"

	o.AddComment(r.Replace(s))
}

func (i Impl) Declaration(o interfaces.Outputer) {
	o.Add("module procedure " + i.Name())
}

// Definition produces the string representation of the implemention, making the required
// replacements.
func (i Impl) Definition(o interfaces.Outputer) {
	r := i.newReplacer()

	s := r.Replace(i.BodyTemplate)

	o.AddLine()
	i.Description(o)
	o.Add(s)
	o.Newline()
}

// Name returns the expanded name of the Impl instance. This takes the type name and
// dimensionality into account.
func (i Impl) Name() string {

	s := i.Basename + "_" + stripTypename(i.Typename)

	if i.Dimension != 0 {
		s += "_" + nd(i.Dimension)
	}

	return s
}

// Strip unwanted characters (e.g. '(', ')') from the typename.
func stripTypename(typename string) string {
	r := strings.NewReplacer(
		"(", "",
		")", "",
	)

	t := r.Replace(typename)
	return t
}

// n returns the string representation of the
func n(dim int) string {
	return fmt.Sprintf("%v", dim)
}

func nd(dim int) string {
	if dim != 0 {
		return n(dim) + "d"
	}
	return ""
}
