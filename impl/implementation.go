package impl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
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

	output := ""

	start := strings.Index(s, "{{")
	for start > -1 {
		end := strings.Index(s, "}}")

		if end < 0 {
			break
		}
		end += 2

		macro := s[start:end]
		replace := expand(macro)
		output += s[:start] + replace

		s = strings.Replace(s[end:], macro, replace, -1)
		start = strings.Index(s, "{{")
	}

	output += s
	o.AddLine()
	i.Description(o)
	o.Add(output)
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

func repeat(text string, delimiter string, times int) string {
	var s []string

	s = make([]string, 0, times)
	for i := 0; i < times; i++ {
		s = append(s, text)
	}

	return strings.Join(s, delimiter)
}

func forloop(text string, delimiter string, start, stop, step int) string {
	var s []string
	times := (stop - start) / step

	// Ensure that there is a %d in the text.
	if !strings.Contains(text, "%d") {
		return repeat(text, delimiter, times)
	}

	s = make([]string, 0, times)

	for i := start; i < stop; i += step {
		s = append(s, fmt.Sprintf(text, i))
	}

	return strings.Join(s, delimiter)
}

func expand(macro string) string {
	opened := strings.HasPrefix(macro, "{{")
	closed := strings.HasSuffix(macro, "}}")

	if !(opened && closed) {
		return macro
	}

	parts := strings.Split(macro[2:len(macro)-2], ";")

	op := parts[0]

	switch op {
	case "for":
		start, _ := strconv.Atoi(parts[1])
		stop, _ := strconv.Atoi(parts[2])
		step, _ := strconv.Atoi(parts[3])
		text := parts[4]
		delimiter := ""
		if len(parts) > 5 {
			delimiter = parts[5]
		}
		macro = forloop(text, delimiter, start, stop, step)
	default:
		panic(errors.New("Undefined macro: " + macro))
	}

	macro = strings.Replace(macro, "\\n", "\n", -1)
	return macro
}
