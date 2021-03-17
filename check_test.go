package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"testing"
)

func loadFile(fileName string, t *testing.T) ast.CommentMap {
	fset := token.NewFileSet()
	path, _ := filepath.Abs(fileName)
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	cmap := ast.NewCommentMap(fset, f, f.Comments)
	return cmap
}

func Test_CheckFunc(t *testing.T) {
	cmap := loadFile("./example/foo.go", t)
	for k, v := range cmap {
		t.Log(checkFunc(k, v))
	}
	cmap = loadFile("./example/foo2.go", t)
	for k, v := range cmap {
		t.Log(checkStruct(k, v))
	}
}

func Test_CheckStruct(t *testing.T) {
	cmap := loadFile("./example/foo.go", t)
	for k, v := range cmap {
		t.Log(checkStruct(k, v))
	}
	cmap = loadFile("./example/bar.go", t)
	for k, v := range cmap {
		t.Log(checkStruct(k, v))
	}
}

func Test_processFile(t *testing.T) {
	t.Log(processFile("./example/foo.go"))
	t.Log(processFile("./example/foo2.go"))
}

func Test_processDir(t *testing.T) {
	t.Log(processDir("./example"))
}

func Test_buildTree(t *testing.T) {
	rs, err := processDir("./example")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rs.buildTree())
}

func Test(t *testing.T) {
	
}