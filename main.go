package main

import (
	"fmt"

	"github.com/ablades/go-pack/pack"
)

func main() {
	p := pack.PK{
		Name: "george",
		Test: 10,
	}

	fmt.Print(p)
}
