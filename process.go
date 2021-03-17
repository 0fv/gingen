package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func processDir(dir string) (rs routeList, err error) {
	var tempFuncs functionList
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		fInfo, err := file.Info()
		if err != nil {
			return nil, err
		}
		if !fInfo.IsDir() && strings.HasSuffix(fInfo.Name(), ".go") {
			rt, fs, err := processFile(dir + "/" + fInfo.Name())
			if err != nil {
				return nil, err
			}
			//assign function
			for i, v := range rt {
				in, ex := fs.splitByRecv(v.name)
				in2, tempFuncs := tempFuncs.splitByRecv(v.name)
				tempFuncs = append(tempFuncs, ex...)
				rt[i].functionList = append(rt[i].functionList, in...)
				rt[i].functionList = append(rt[i].functionList, in2...)
			}
			rs = append(rs, rt...)
		}
	}
	return
}

func processFile(fileName string) (rs routeList, fs functionList, err error) {
	fset := token.NewFileSet()
	path, _ := filepath.Abs(fileName)
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}
	cmap := ast.NewCommentMap(fset, f, f.Comments)
	for k, v := range cmap {
		var fInfo *functionInfo
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
