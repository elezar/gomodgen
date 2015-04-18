package impl

import (
	"testing"

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
		this := Impl{
			BodyTemplate: c.def,
			Typename:     "integer",
		}

		assert.Equal(t, c.want, this.Definition(), "")
	}
}

func TestBasenameReplacement(t *testing.T) {
	this := Impl{BodyTemplate: "{{basename}}", Basename: "foo"}

	assert.Equal(t, "foo", this.Definition(), "")

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

		assert.Equal(t, c.want, this.Definition(), "")
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

		assert.Equal(t, c.want, this.Definition(), "")
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
