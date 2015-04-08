package main

import (
	"io/ioutil"
	"testing"

	"github.com/elezar/gomodgen/impl"
	"github.com/stretchr/testify/assert"
)

func TestFooInteger(t *testing.T) {
	datIn, err := ioutil.ReadFile("test_subroutine_body.in.f90")
	if err != nil {
		t.Error(err)
	}

	datOut, err := ioutil.ReadFile("expected_subroutine_body.out.f90")
	if err != nil {
		t.Error(err)
	}

	this := impl.Impl{Body: string(datIn), Basename: "foo", Type: "integer"}
	assert.Equal(t, string(datOut), this.String(), "")

	// One can write out the file to disk so that the expected and actual results can be
	// properly compared.
	// ioutil.WriteFile("foo.f90", []byte(this.String()), 0644)

}
