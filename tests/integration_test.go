package main

import (
	"io/ioutil"
	"testing"

	"github.com/elezar/gomodgen/generic"
	"github.com/elezar/gomodgen/impl"
	"github.com/stretchr/testify/assert"
)

func TestFooInteger(t *testing.T) {

	this := impl.Impl{Basename: "foo"}
	err := this.LoadBody("test_subroutine_body.in.f90")
	if err != nil {
		t.Error(err)
	}

	datOut, err := ioutil.ReadFile("expected_subroutine_body.out.f90")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, string(datOut), this.Definition("integer", 0), "")

	// One can write out the file to disk so that the expected and actual results can be
	// properly compared.
	// ioutil.WriteFile("foo.f90", []byte(this.String()), 0644)

}

func TestFooGenericInterface(t *testing.T) {

	var this generic.Generic

	err := this.Load("test_generic_interface.in.json")
	if err != nil {
		panic(err)
	}

	expected := generic.Generic{
		Name:       "foo",
		BodyFile:   "test_subroutine_body.in.f90",
		Types:      []string{"integer"},
		Dimensions: []int{0}}

	assert.Equal(t, expected.Name, this.Name, "")
	assert.Equal(t, expected.BodyFile, this.BodyFile, "")
	assert.Equal(t, expected.Types, this.Types, "")
	assert.Equal(t, expected.Dimensions, this.Dimensions, "")

}
