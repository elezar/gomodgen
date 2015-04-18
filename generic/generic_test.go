package generic

import (
	"testing"

	"github.com/elezar/gomodgen/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	this := Generic{
		Name: "foo",
	}

	i := interfaces.EntityTester{Decl: "module procedure foo_integer"}

	this.Add(i)
	expected := "interface foo\nmodule procedure foo_integer\nend interface"
	var o *interfaces.OutputTester = new(interfaces.OutputTester)
	this.Declaration(o)
	assert.Equal(t, expected, o.String(), "")
}
