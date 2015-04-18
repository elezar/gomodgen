package module

import (
	"fmt"
	"testing"

	"github.com/elezar/gomodgen/interfaces"
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
func (g G) Description() string { return "Description" }

func TestAddEntity(t *testing.T) {

	m := Module{
		Desc: "A collection of string utilities",
		Name: "string_utils",
	}

	var g interfaces.EntityTester

	m.Add(g)

	assert.Equal(t, "string_utils", m.Name, "")
	fmt.Println(m.Generate())
}
