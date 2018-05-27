package main

import (
	"fmt"

	"github.com/arehmandev/designpatterns/singleton/pkg/singleton"
)

func main() {

	s := singleton.New()

	s["this"] = "that"

	s2 := singleton.New()

	fmt.Println(s2, s)
}
