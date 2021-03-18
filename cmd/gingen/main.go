package main

import (
	"flag"
	"log"

	"github.com/0fv/gingen"
)

var (
	underline = flag.String("u", "true", "struct name to underline route")
)

func main() {
	gingen.UnderlineSet(*underline)
	rs, err := gingen.ProcessDir()
	if err != nil {
		log.Fatal(err)
	}
	rs.BuildTree()
	if err = gingen.GenFile(rs); err != nil {
		log.Fatal(err)
	}
}
