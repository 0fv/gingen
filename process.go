package gingen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

var pkgName string

//ProcessDir ...
func ProcessDir() (rs RouteList, err error) {
	return processDir("./")
}

func processDir(dir string) (rs RouteList, err error) {
	var tempFuncs FunctionList
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		fInfo, err := file.Info()
		if err != nil {
			return nil, err
		}
		if !fInfo.IsDir() && strings.HasSuffix(fInfo.Name(), ".go") && (!strings.HasSuffix(fInfo.Name(), suffix)) {
			cmap, err := cmapgen(dir + "/" + fInfo.Name())
			if err != nil {
				return nil, err
			}
			rt, fs, err := processFile(cmap, fInfo.Name())
			if err != nil {
				return nil, err
			}
			//assign function
			for i, v := range rt {
				in, ex := fs.splitByRecv(v.Name)
				in2, tempFuncs := tempFuncs.splitByRecv(v.Name)
				tempFuncs = append(tempFuncs, ex...)
				rt[i].FunctionList = append(rt[i].FunctionList, in...)
				rt[i].FunctionList = append(rt[i].FunctionList, in2...)
			}
			rs = append(rs, rt...)
		}
	}
	return
}

func cmapgen(fileName string) (ast.CommentMap, error) {
	fset := token.NewFileSet()
	path, _ := filepath.Abs(fileName)
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	if pkgName == "" {
		pkgName = f.Name.Name
	}
	return ast.NewCommentMap(fset, f, f.Comments), nil
}

func processFile(cmap ast.CommentMap, fileName string) (rs RouteList, fs FunctionList, err error) {
	for k, v := range cmap {
		var fInfo *FunctionInfo
		fInfo, err = checkFunc(k, v)
		if err != nil {
			err = fmt.Errorf("%v: %v", fileName, err.Error())
			return
		}
		if fInfo != nil {
			fs = append(fs, fInfo)
			continue
		}
		rInfo, ok := checkStruct(k, v)
		if ok {
			rs = append(rs, rInfo)
		}
	}
	return
}
