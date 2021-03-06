package impl

import (
	"fmt"
	"testing"

	"github.com/elezar/gomodgen/interfaces"

	"github.com/stretchr/testify/assert"
)

func TestTypeNameReplacement(t *testing.T) {

	cases := []struct {
		def  string
		want string
	}{
		{"other string", "other string"},
		{"{{type", "{{type"},
		{"{type}", "{type}"},
		{"{{type}}", "integer"},
		{"{{type}} {{type}}", "integer integer"},
		{"preamble{{type}}", "preambleinteger"},
	}
	for _, c := range cases {

		var o *interfaces.OutputTester = new(interfaces.OutputTester)

		this := Impl{
			BodyTemplate: c.def,
			Typename:     "integer",
		}

		this.Definition(o)
		assert.Equal(t, c.want, o.String(), "")
	}
}

func TestBasenameReplacement(t *testing.T) {
	this := Impl{BodyTemplate: "{{basename}}", Basename: "foo"}

	var o *interfaces.OutputTester = new(interfaces.OutputTester)
	this.Definition(o)
	assert.Equal(t, "foo", o.String(), "")

}

func TestNameReplacement(t *testing.T) {
	cases := []struct {
		d    int
		want string
	}{
		{0, "foo_integer"},
		{1, "foo_integer_1d"},
		{2, "foo_integer_2d"},
		{99, "foo_integer_99d"},
	}
	for _, c := range cases {

		this := Impl{Basename: "foo",
			BodyTemplate: "{{name}}",
			Typename:     "integer",
			Dimension:    c.d,
		}

		var o *interfaces.OutputTester = new(interfaces.OutputTester)
		this.Definition(o)

		assert.Equal(t, c.want, o.String(), "")
	}
}

func TestExpandedName(t *testing.T) {

	cases := []struct {
		d    int
		want string
	}{
		{0, "foo_integer"},
		{1, "foo_integer_1d"},
		{2, "foo_integer_2d"},
		{99, "foo_integer_99d"},
	}
	for _, c := range cases {
		this := Impl{Basename: "foo",
			Typename:  "integer",
			Dimension: c.d,
		}

		assert.Equal(t, c.want, this.Name(), "")
	}

}

func TestDimensionString(t *testing.T) {

	cases := []struct {
		d    int
		want string
	}{
		{0, ""},
		{1, "1d"},
		{2, "2d"},
		{99, "99d"},
	}
	for _, c := range cases {
		assert.Equal(t, c.want, nd(c.d), "")
	}

}

func TestNString(t *testing.T) {

	cases := []struct {
		d    int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{2, "2"},
		{99, "99"},
	}
	for _, c := range cases {
		this := Impl{BodyTemplate: "{{n}}",
			Dimension: c.d}

		var o *interfaces.OutputTester = new(interfaces.OutputTester)
		this.Definition(o)
		assert.Equal(t, c.want, o.String(), "")
	}

}

func TestStripTypenames(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"integer", "integer"},
		{"real(4)", "real4"},
	}

	for _, c := range cases {
		assert.Equal(t, c.want, stripTypename(c.in), "For: "+c.in)
	}
}

func TestRepeat(t *testing.T) {
	text := ":"
	output := repeat(text, ",", 1)
	assert.Equal(t, ":", output, "")
}

func TestLoop(t *testing.T) {
	text := "d%d"
	output := forloop(text, ", ", 1, 3, 1)
	assert.Equal(t, "d1, d2", output, "")
}

func TestExpand(t *testing.T) {
	in := "{{for;0;3;1;:;,}}"
	output := expand(in)

	assert.Equal(t, ":,:,:", output, "")

}

func TestExpandNewLine(t *testing.T) {
	in := "{{for;0;3;1;:;\n}}"

	fmt.Println(in)
	output := expand(in)

	assert.Equal(t, ":\n:\n:", output, "")

}
