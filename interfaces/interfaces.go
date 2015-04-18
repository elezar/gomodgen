package interfaces

import (
	"strings"
)

type Outputer interface {
	Add(s string)
	AddComment(s string)
	Newline()
	AddLine()
	Indent()
	Deindent()
}

type Entity interface {
	Description(o Outputer)
	Declaration(o Outputer)
	Definition(o Outputer)
}

type EntityLoader interface {
	Load(a EntityAdder, root string, filelist []string) error
}

type EntityAdder interface {
	Add(e Entity)
}

type OutputTester struct {
	text string
}

func (o *OutputTester) Add(s string) {
	o.text += "\n" + s
}
func (o OutputTester) String() string {
	return strings.TrimSpace(o.text)
}

func (o *OutputTester) AddComment(s string) {}
func (o *OutputTester) Newline()            {}
func (o *OutputTester) AddLine()            {}
func (o *OutputTester) Indent()             {}
func (o *OutputTester) Deindent()           {}

type EntityTester struct {
	Desc string
	Decl string
	Defn string
}

func (et EntityTester) Declaration(o Outputer) { o.Add(et.Decl) }
func (et EntityTester) Definition(o Outputer)  { o.Add(et.Defn) }
func (et EntityTester) Description(o Outputer) { o.Add(et.Desc) }
