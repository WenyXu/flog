package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/mingrammer/cfmt"
)

func main() {
	opts := ParseOptions()
	gofakeit.GlobalFaker = gofakeit.New(opts.Seed)
	if err := Run(opts); err != nil {
		cfmt.Warningln(err.Error())
	}
}
