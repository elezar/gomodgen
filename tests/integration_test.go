package main

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/elezar/gomodgen/generic"
	"github.com/elezar/gomodgen/impl"
	"github.com/elezar/gomodgen/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestFooInteger(t *testing.T) {

	this := impl.Impl{
		Basename: "foo",
		Typename: "integer",
	}
	err := this.LoadBody("test_subroutine_body.in.f90")
	if err != nil {
		t.Error(err)
	}

	datOut, err := ioutil.ReadFile("expected_subroutine_body.out.f90")
	if err != nil {
		t.Error(err)
	}

	var o *interfaces.OutputTester = new(interfaces.OutputTester)
	this.Definition(o)
	assert.Equal(t, strings.TrimSpace(string(datOut)), o.String(), "")

	// One can write out the file to disk so that the expected and actual results can be
	// properly compared.
	// ioutil.WriteFile("foo.f90", []byte(this.String()), 0644)

}

func TestFooGenericInterface(t *testing.T) {

	var this *generic.Generic

	this = generic.Load("test_generic_interface.in.json")
	if this == nil {
		panic(errors.New("Could not load generic"))
	}

	expected := generic.Generic{
		Name: "foo",
	}

	assert.Equal(t, expected.Name, this.Name, "")

}
