package main

import (
	"fmt"
	"io/ioutil"

	"github.com/elezar/gomodgen/impl"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := ioutil.ReadFile("test_subroutine_body.in.f90")
	check(err)

	fmt.Println(string(dat))

	i := impl.Impl{Body: string(dat), Basename: "foo", Type: "integer"}

	fmt.Println(i)

}
