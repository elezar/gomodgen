package impl

import (
	"fmt"
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
func (i Impl) newReplacer() *strings.Replacer {
	r := strings.NewReplacer(typetag, i.Type, nametag, i.GetName(), ntag, i.n(), basetag, i.Basename)
	return r
}

// Impl represents a specific function or subroutine implementation for a given type and
// dimension. This allows for the creation of a specific implementation of a generic
// interface. Certain tags (e.g. {{type}}, {{name}}) are replaced for the final string
// representation.
type Impl struct {
	Basename string
	Body     string
	Type     string
	Dim      int
}

// String produces the string representation of the implemention, making the required
// replacements.
func (i Impl) String() string {
	r := i.newReplacer()

	s := r.Replace(i.Body)

	return s
}

// GetName returns the expanded name of the Impl instance. This takes the type name and
// dimensionality into account.
func (i Impl) GetName() string {

	s := i.Basename + "_" + i.Type

	if i.Dim != 0 {
		s += "_" + i.nd()
	}

	return s
}

// n returns the string representation of the
func (i Impl) n() string {
	return fmt.Sprintf("%v", i.Dim)
}

func (i Impl) nd() string {
	if i.Dim != 0 {
		return i.n() + "d"
	}
	return ""
}
