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
		this := Impl{Body: c.def, Type: "integer"}

		assert.Equal(t, c.want, this.String(), "")
	}
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
		this := Impl{Basename: "foo", Body: "{{name}}", Type: "integer", Dim: c.d}

		assert.Equal(t, c.want, this.String(), "")
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
		this := Impl{Basename: "foo", Type: "integer", Dim: c.d}

		assert.Equal(t, c.want, this.GetName(), "")
	}

}

func TestExpandedName1D(t *testing.T) {
	this := Impl{Basename: "foo", Type: "integer", Dim: 1}

	assert.Equal(t, "foo_integer_1d", this.GetName(), "")
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
		this := Impl{Dim: c.d}

		assert.Equal(t, c.want, this.nd(), "")
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
		this := Impl{Body: "{{n}}", Dim: c.d}

		assert.Equal(t, c.want, this.String(), "")
	}

}
