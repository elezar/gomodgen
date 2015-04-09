package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	this := Generic{
		Name:       "foo",
		Types:      []string{"integer"},
		Dimensions: []int{0},
	}

	expected := "interface foo\nmodule procedure foo_integer\nend interface"
	assert.Equal(t, expected, this.Declaration(), "")
}
