package impl

import (
	"fmt"
	"io/ioutil"
	"strings"
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
func (i Impl) newReplacer(typename string, dim int) *strings.Replacer {
	r := strings.NewReplacer(
		typetag, typename,
		nametag, i.Name(typename, dim),
		ntag, n(dim),
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
}

// LoadBody loads the body of an Impl from the file specified
func (i *Impl) LoadBody(filename string) error {

	i.BodyTemplate = ""

	dataIn, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filename, err)
		return err
	}

	i.BodyTemplate = string(dataIn)

	return nil
}

// Definition produces the string representation of the implemention, making the required
// replacements.
func (i Impl) Definition(typename string, dim int) string {
	r := i.newReplacer(typename, dim)

	s := r.Replace(i.BodyTemplate)

	return s
}

// Name returns the expanded name of the Impl instance. This takes the type name and
// dimensionality into account.
func (i Impl) Name(typename string, dim int) string {

	s := i.Basename + "_" + typename

	if dim != 0 {
		s += "_" + nd(dim)
	}

	return s
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
