package module

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringUtils(t *testing.T) {
	m := Module{
		Desc: "A collection of string utilities",
		Name: "string_utils",
	}

	assert.Equal(t, "string_utils", m.Name, "")
	fmt.Println(m.Generate())

}

type G int

func (g G) Declaration() string { return "Declaration" }
func (g G) Definition() string  { return "Definition" }

func TestAddEntity(t *testing.T) {

	m := Module{
		Desc: "A collection of string utilities",
		Name: "string_utils",
	}

	var g G

	m.Add(g)

	assert.Equal(t, "string_utils", m.Name, "")
	fmt.Println(m.Generate())
}
