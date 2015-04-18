package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type I int

func (i I) Declaration() string { return "module procedure foo_integer" }
func (i I) Definition() string  { return "Definition" }

func TestBasic(t *testing.T) {
	this := Generic{
		Name: "foo",
	}

	var i I
	this.Add(i)
	expected := "interface foo\nmodule procedure foo_integer\nend interface"
	assert.Equal(t, expected, this.Declaration(), "")
}
