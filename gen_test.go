package gingen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
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

// func Test_CheckPackage(t *testing.T) {
// 	cmap = loadFile("./example/bar.go", t)
// 	processPkgName
// }

func Test_processDir(t *testing.T) {
	t.Log(processDir("./example"))
}

func br(t *testing.T) RouteList {
	rs, err := processDir("./example")
	if err != nil {
		t.Fatal(err)
	}
	return rs.BuildTree()
}

func Test_buildTree(t *testing.T) {
	t.Log(br(t))
}

func Test_genRoot(t *testing.T) {
	rs := br(t)
	err := genRoot(rs, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_genSub(t *testing.T) {
	r := br(t)
	err := genSub(r[0], os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_genAll(t *testing.T) {
	r := br(t)
	err := genAll(r, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}
