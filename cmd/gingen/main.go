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
	flag.Parse()
	gingen.UnderlineSet(*underline)
	rs, err := gingen.ProcessDir()
	if err != nil {
		log.Fatal(err)
	}
	if len(rs) == 0 {
		log.Fatal("no relative route found")
	}
	if err = gingen.GenFile(rs.BuildTree()); err != nil {
		log.Fatal(err)
	}
}
