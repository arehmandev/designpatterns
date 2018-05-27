package main

import (
	"designpatterns/singleton/pkg/singleton"
	"fmt"
)

func main() {

	s := singleton.New()

	s["this"] = "that"

	s2 := singleton.New()

	fmt.Println(s2, s)
}
