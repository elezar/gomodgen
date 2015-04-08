package impl

import (
	"fmt"
	"strings"
)

const (
	typetag = "{{type}}"
	nametag = "{{name}}"
	ntag    = "{{n}}"
)

// Impl represents a specific function implementation.
type Impl struct {
	Basename string
	Body     string
	Type     string
	Dim      int
}

func (i Impl) String() string {
	r := i.newReplacer()

	s := r.Replace(i.Body)

	return s
}

// GetName returns the expanded name of the Impl instance.
func (i Impl) GetName() string {

	s := i.Basename + "_" + i.Type

	if i.Dim != 0 {
		s += "_" + i.nd()
	}

	return s
}

func (i Impl) n() string {
	return fmt.Sprintf("%v", i.Dim)
}

func (i Impl) nd() string {
	if i.Dim != 0 {
		return i.n() + "d"
	}
	return ""
}

func (i Impl) newReplacer() *strings.Replacer {
	r := strings.NewReplacer(typetag, i.Type, nametag, i.GetName(), ntag, i.n())
	return r
}
