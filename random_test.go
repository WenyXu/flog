package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
)

func ExampleRandResourceURI() {
	gofakeit.GlobalFaker = gofakeit.New(11)
	fmt.Print(RandResourceURI())
	// Output: /infrastructures
}
