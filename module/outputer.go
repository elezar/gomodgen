package module

import (
	"strings"
)

const (
	commentTag = "! "
	tab        = "  "
)

type Outputer struct {
	lines []string
	ind   int
}

func (o *Outputer) Indent() {
	o.ind += 1
}

func (o *Outputer) Deindent() {
	o.ind -= 1
}

// Add adds a newLine to a set of lines using the specified indent.
func (o *Outputer) Add(newLine string) {
	var prefix string

	if newLine != "" {
		prefix = o.indentPrefix()
	}
	o.lines = append(o.lines, prefix+newLine)

}

// AddLine adds a horizontal line to the outputer.
func (o *Outputer) AddLine() {
	s := strings.TrimSpace(commentTag)
	i := o.ind

	for len(s) < 90 {
		s += "="
	}

	o.ind = 0
	o.Add(s)
	o.ind = i
}

func (o *Outputer) AddComment(comment string) {
	if len(comment) > 0 {
		o.Add(commentTag + comment)
	}
}

func (o *Outputer) Newline() {
	o.Add("")
}

func (o Outputer) indentPrefix() string {
	var s string
	for i := 0; i < o.ind; i++ {
		s += tab
	}
	return s
}

func (o Outputer) String() string {
	return strings.Join(o.lines, "\n")
}
