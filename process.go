package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strings"
)

func processFile(fileName string) {
	fset := token.NewFileSet()
	path, _ := filepath.Abs(fileName)
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}
	cmap := ast.NewCommentMap(fset, f, f.Comments)
	rMap := make(map[string]routeInfo)
	for k, v := range cmap {
		if spec, ok := k.(*ast.GenDecl); ok {
			if spec.Tok == token.TYPE {
				name := (spec.Specs[0]).(*ast.TypeSpec).Name.Name
				for _, comment := range v {
					cm := comment.Text()
					if strings.HasPrefix(cm, "! ") {
						log.Println("struct get:", name)
						rMap[name] = routeInfo{
							comment:    cm,
							structInfo: spec,
						}
						break
					}
				}
			}

		}
	}
}
